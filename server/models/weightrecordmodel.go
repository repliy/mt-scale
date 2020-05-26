package models

import (
	"encoding/json"
	"mt-scale/common"
	"mt-scale/entitys"
	"mt-scale/exception"
	"mt-scale/models/dto"
	"mt-scale/models/vo"
	"mt-scale/syslog"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddWeightRecord record weight one time
func AddWeightRecord(param dto.AddWeightRecordDto) primitive.ObjectID {
	col, ctx := Collection("weightrecord")
	filter := bson.D{
		primitive.E{
			Key:   "index",
			Value: param.Index,
		},
		primitive.E{
			Key:   "task_id",
			Value: param.TaskID,
		},
	}
	cur, err := col.Find(ctx, filter)
	if err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.DatabaseErrorCode)
	}
	if cur.Next(ctx) {
		// override record first
		var record entitys.WeightRecord
		if err := cur.Decode(&record); err != nil {
			syslog.Error(err)
			exception.ThrowBusinessError(common.DatabaseErrorCode)
		}
		tag := SelectTagByID(record.ID)
		updRecord := dto.UpdWeightRecordDto{
			ID:        record.ID,
			Weight:    record.Weight,
			BoxID:     record.BoxID,
			SpeciesID: record.SpeciesID,
			Index:     record.Index,
			TagName:   tag.Name,
		}
		return UpdWeightRecord(updRecord)
	}
	// input tag is avliable
	var tag entitys.FishTag
	if param.TagName != "" {
		tag = SelectTagByName(param.TagName)
		if tag == (entitys.FishTag{}) {
			tag.ID = AddTag(entitys.FishTag{
				SpeciesID: param.SpeciesID,
				Name:      param.TagName,
			})
		}
		if tag.Used {
			exception.ThrowBusinessErrorMsg("绑定的Tag已经占用")
		}
	}
	jsonParam, _ := json.Marshal(param)
	entity := new(entitys.WeightRecord)
	json.Unmarshal(jsonParam, entity)

	entity.TagID = tag.ID
	nowtime := time.Now()
	entity.CreateTime = nowtime
	entity.UpdateTime = nowtime
	entity.Creator = ""

	result, err := col.InsertOne(ctx, entity)
	if err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.DatabaseErrorCode)
	}
	if tag.ID != primitive.NilObjectID {
		// flag tag is used
		UpdateFishTagStatus(tag.ID, true)
	}

	return result.InsertedID.(primitive.ObjectID)
}

// UpdWeightRecord update weight record by index
func UpdWeightRecord(dto dto.UpdWeightRecordDto) primitive.ObjectID {
	col, ctx := Collection("weightrecord")
	filter := bson.D{
		primitive.E{
			Key:   "_id",
			Value: dto.ID,
		},
	}
	var record entitys.WeightRecord
	if err := col.FindOne(ctx, filter).Decode(&record); err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.DatabaseErrorCode)
	}
	var tag entitys.FishTag
	if dto.TagName != "" {
		tag = SelectTagByName(dto.TagName)
		if tag == (entitys.FishTag{}) {
			tag.ID = AddTag(entitys.FishTag{
				SpeciesID: dto.SpeciesID,
				Name:      dto.TagName,
			})
		}
		if tag.Used {
			exception.ThrowBusinessErrorMsg("绑定的Tag已经占用")
		}
	}
	update := bson.D{
		primitive.E{
			Key: "$set",
			Value: bson.M{
				"weight":        dto.Weight,
				"box_id":        dto.BoxID,
				"species_id":    dto.SpeciesID,
				"tag_id":        tag.ID,
				"update_time":   time.Now(),
				"last_operator": "",
			},
		},
	}
	if _, err := col.UpdateOne(ctx, filter, update); err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.DatabaseErrorCode)
	}
	// change tag
	if record.TagID != tag.ID {
		// restore origin tag
		UpdateFishTagStatus(record.TagID, false)
		// flag new tag is used
		UpdateFishTagStatus(tag.ID, true)
	}
	return record.ID
}

func validateBoxAndSpec(record entitys.WeightRecord) {
	if box := SelectBoxByID(record.BoxID); box == (entitys.Box{}) {
		exception.ThrowBusinessErrorMsg("绑定的箱子不存在")
	}
	spec := SelectSpeciesByID(record.SpeciesID)
	if spec == (entitys.Species{}) {
		exception.ThrowBusinessErrorMsg("绑定的物种不存在")
	}
	if record.TagID != primitive.NilObjectID {
		if !spec.HasTag {
			exception.ThrowBusinessErrorMsg("该物种未设置可用Tag")
		}
	}
}

// FetchWeightRecord fetch weight record by conditions
func FetchWeightRecord(dto dto.QueryRecordDto) []vo.WeightRecordVo {
	col, ctx := Collection("weightrecord")
	pageNum := dto.PageNum
	if pageNum == 0 {
		pageNum = 1
	}
	pageSize := dto.PageSize
	if pageSize == 0 {
		pageSize = 10
	}
	limit := int64(pageSize)
	skip := int64((pageNum - 1) * pageSize)
	taskBsonID, _ := primitive.ObjectIDFromHex(dto.TaskID)
	filter := []bson.M{
		{
			"$match": bson.M{
				"task_id": taskBsonID,
			},
		},
		{
			"$lookup": bson.M{
				"from":         "species",
				"localField":   "species_id",
				"foreignField": "_id",
				"as":           "species",
			},
		},
		{
			"$unwind": bson.M{
				"path":                       "$species",
				"preserveNullAndEmptyArrays": true,
			},
		},
		{
			"$lookup": bson.M{
				"from":         "fishtag",
				"localField":   "tag_id",
				"foreignField": "_id",
				"as":           "tags",
			},
		},
		{
			"$unwind": bson.M{
				"path":                       "$tags",
				"preserveNullAndEmptyArrays": true,
			},
		},
		{
			"$lookup": bson.M{
				"from":         "box",
				"localField":   "box_id",
				"foreignField": "_id",
				"as":           "boxes",
			},
		},
		{
			"$unwind": bson.M{
				"path":                       "$boxes",
				"preserveNullAndEmptyArrays": true,
			},
		},
		{
			"$project": bson.M{
				"_id":           1,
				"weight":        1,
				"index":         1,
				"tag_id":        1,
				"tag":           "$tags.name",
				"species_id":    1,
				"species":       "$species.name",
				"species_color": "$species.color",
				"box_id":        1,
				"box_type":      "$boxes.type",
				"box_num":       "$boxes.num",
			},
		},
		{
			"$sort": bson.M{
				"index": -1,
			},
		},
		{
			"$skip": skip,
		},
		{
			"$limit": limit,
		},
	}
	if dto.BoxID != "" {
		boxBsonID, _ := primitive.ObjectIDFromHex(dto.BoxID)
		filter = append(filter, primitive.M{
			"$match": primitive.M{
				"box_id": boxBsonID,
			},
		})
	}
	if dto.SpeciesID != "" {
		specBsonID, _ := primitive.ObjectIDFromHex(dto.SpeciesID)
		filter = append(filter, primitive.M{
			"$match": primitive.M{
				"species_id": specBsonID,
			},
		})
	}
	if dto.Index != 0 {
		filter = append(filter, primitive.M{
			"$match": primitive.M{
				"index": dto.Index,
			},
		})
	}

	cur, err := col.Aggregate(ctx, filter)
	if err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.DatabaseErrorCode)
	}
	var result []vo.WeightRecordVo
	for cur.Next(ctx) {
		var row vo.WeightRecordVo
		if err := cur.Decode(&row); err != nil {
			syslog.Error(err)
			exception.ThrowBusinessError(common.DatabaseErrorCode)
		}
		result = append(result, row)
	}
	return result
}

// DeleteWeightRecord Delete weight record
func DeleteWeightRecord(dto dto.DelRecordDto) {
	col, ctx := Collection("weightrecord")
	recordBsonID, _ := primitive.ObjectIDFromHex(dto.ID)
	filter := bson.M{
		"_id": recordBsonID,
	}
	if _, err := col.DeleteOne(ctx, filter); err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.DatabaseErrorCode)
	}
}

// StatTotalWeight Statis total weight
func StatTotalWeight(taskID primitive.ObjectID) vo.StatTotalWeightVo {
	col, ctx := Collection("weightrecord")
	filter := []bson.M{
		{
			"$match": bson.M{
				"task_id": taskID,
			},
		},
		{
			"$group": bson.M{
				"_id": "null",
				"weight": bson.M{
					"$sum": "$weight",
				},
			},
		},
	}
	cur, err := col.Aggregate(ctx, filter)
	if err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.DatabaseErrorCode)
	}
	var total vo.StatTotalWeightVo
	if cur.Next(ctx) {
		if err := cur.Decode(&total); err != nil {
			syslog.Error(err)
			exception.ThrowBusinessError(common.DatabaseErrorCode)
		}
	}
	return total
}
