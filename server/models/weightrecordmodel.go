package models

import (
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
func AddWeightRecord(param entitys.WeightRecord) primitive.ObjectID {
	col, ctx := Collection("weightrecord")
	filter := bson.D{
		primitive.E{Key: "index", Value: param.Index},
	}
	cur, err := col.Find(ctx, filter)
	if err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.DatabaseErrorCode)
	}
	if cur.Next(ctx) {
		// override record first
		return UpdWeightRecord(param)
	}
	validateBoxAndSpec(param)
	if param.TagID != primitive.NilObjectID {
		tag := SelectTagByID(param.TagID)
		if tag == (entitys.FishTag{}) {
			exception.ThrowBusinessErrorMsg("绑定的Tag不存在")
		}
		if tag.Used {
			exception.ThrowBusinessErrorMsg("绑定的Tag已经占用")
		}
	}
	param.CreateTime = time.Now()
	param.Creator = ""
	result, err := col.InsertOne(ctx, param)
	if err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.DatabaseErrorCode)
	}
	// flag tag is used
	UpdateFishTagStatus(param.TagID, true)
	return result.InsertedID.(primitive.ObjectID)
}

// UpdWeightRecord update weight record by index
func UpdWeightRecord(param entitys.WeightRecord) primitive.ObjectID {
	col, ctx := Collection("weightrecord")

	filter := bson.D{primitive.E{Key: "index", Value: param.Index}}
	var record entitys.WeightRecord
	if err := col.FindOne(ctx, filter).Decode(&record); err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.DatabaseErrorCode)
	}
	validateBoxAndSpec(record)
	if param.TagID != primitive.NilObjectID && param.TagID != record.TagID {
		tag := SelectTagByID(param.TagID)
		if tag == (entitys.FishTag{}) {
			exception.ThrowBusinessErrorMsg("绑定的Tag不存在")
		}
		if tag.Used {
			exception.ThrowBusinessErrorMsg("绑定的Tag已经占用")
		}
	}
	update := bson.D{
		primitive.E{
			Key: "$set", Value: bson.M{
				"weight":        param.Weight,
				"box_id":        param.BoxID,
				"species_id":    param.SpeciesID,
				"tag_id":        param.TagID,
				"update_time":   time.Now(),
				"last_operator": "",
			},
		},
	}
	if _, err := col.UpdateOne(ctx, filter, update); err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.DatabaseErrorCode)
	}
	if param.TagID != primitive.NilObjectID && param.TagID != record.TagID {
		// restore origin tag
		UpdateFishTagStatus(record.TagID, false)
		// flag new tag is used
		UpdateFishTagStatus(param.TagID, true)
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
	filter := []bson.M{
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
			"$project": bson.M{
				"_id":     1,
				"weight":  1,
				"index":   1,
				"tag":     "$tags.name",
				"species": "$species.name",
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
func DeleteWeightRecord(dto dto.DelWeightDto) {
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
