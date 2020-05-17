package models

import (
	"fmt"
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

// AddBoxList Request with list of box parameters
func AddBoxList(boxes []dto.AddBoxDto) (vo.BoxValidateErrorVo, []vo.AddBoxVo) {
	var valResult vo.BoxValidateErrorVo
	var boxList []vo.AddBoxVo
	// validate box info
	for _, box := range boxes {
		_, valMsg, err := validateBox(box)
		if err != nil {
			syslog.Error(err)
			exception.ThrowBusinessError(common.DatabaseErrorCode)
		}
		if valMsg != "" {
			valResult.BoxType = box.Type
			valResult.BoxNum = box.Num
			valResult.ValidateMsg = valMsg
			return valResult, boxList
		}
	}
	// add box with num
	for _, box := range boxes {
		var boxRes vo.AddBoxVo
		boxID := AddBox(box)

		boxRes.BoxID = boxID
		boxRes.BoxType = box.Type
		boxRes.BoxNum = box.Num
		boxList = append(boxList, boxRes)
	}

	return valResult, boxList
}

// AddBox Add new box
func AddBox(box dto.AddBoxDto) primitive.ObjectID {
	// validate request info
	boxID, valMsg, err := validateBox(box)
	if err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.DatabaseErrorCode)
	}
	if valMsg != "" {
		exception.ThrowBusinessErrorMsg(valMsg)
	}
	if boxID != primitive.NilObjectID {
		// already exist
		return boxID
	}
	// new one, insert to database
	col, ctx := Collection("box")
	timeNow := time.Now()
	taskBsonID, _ := primitive.ObjectIDFromHex(box.TaskID)
	insertObj := entitys.Box{
		Type:       box.Type,
		Num:        box.Num,
		Status:     "enable",
		TaskID:     taskBsonID,
		CreateTime: timeNow,
		UpdateTime: timeNow,
	}
	result, err := col.InsertOne(ctx, insertObj)
	if err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.DatabaseErrorCode)
	}
	insertedID := result.InsertedID.(primitive.ObjectID)
	return insertedID
}

func validateBox(box dto.AddBoxDto) (primitive.ObjectID, string, error) {
	col, ctx := Collection("box")
	taskBsonID, _ := primitive.ObjectIDFromHex(box.TaskID)
	filter := bson.D{
		primitive.E{
			Key:   "num",
			Value: box.Num,
		},
		primitive.E{
			Key:   "task_id",
			Value: taskBsonID,
		},
	}
	cur, err := col.Find(ctx, filter)
	if err != nil {
		return primitive.NilObjectID, "", err
	}
	if cur.Next(ctx) {
		var record entitys.Box
		if err := cur.Decode(&record); err != nil {
			return primitive.NilObjectID, "", err
		}
		if record.Type == box.Type {
			return record.ID, "", nil
		}
		return primitive.NilObjectID, fmt.Sprintf("箱号%s被%s箱子占用,请更改...", box.Num, record.Type), nil
	}
	return primitive.NilObjectID, "", nil
}

// SelectBoxByID Select box by id
func SelectBoxByID(id primitive.ObjectID) entitys.Box {
	col, ctx := Collection("box")
	filter := bson.D{
		primitive.E{Key: "_id", Value: id},
	}
	cur, err := col.Find(ctx, filter)
	if err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.DatabaseErrorCode)
	}
	var box entitys.Box
	if cur.Next(ctx) {
		if err := cur.Decode(&box); err != nil {
			syslog.Error(err)
			exception.ThrowBusinessError(common.DatabaseErrorCode)
		}
	}
	return box
}

// FetchBoxes Get boxes by type
func FetchBoxes(dto dto.QueryBoxDto) []entitys.Box {
	col, ctx := Collection("box")
	taskBsonID, _ := primitive.ObjectIDFromHex(dto.TaskID)
	filter := []bson.M{
		{
			"$match": bson.M{
				"status":  "enable",
				"type":    dto.BoxType,
				"task_id": taskBsonID,
			},
		},
		{
			"$sort": bson.M{
				"num": -1,
			},
		},
	}
	cur, err := col.Aggregate(ctx, filter)
	if err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.DatabaseErrorCode)
	}
	var result []entitys.Box
	for cur.Next(ctx) {
		var row entitys.Box
		if err := cur.Decode(&row); err != nil {
			syslog.Error(err)
			exception.ThrowBusinessError(common.DatabaseErrorCode)
		}
		result = append(result, row)
	}
	return result
}

// StatBoxWeight Statis the weight according to the type of box
func StatBoxWeight() []vo.StatBoxWeightVo {
	col, ctx := Collection("box")
	filter := []bson.M{
		{
			"$lookup": bson.M{
				"from":         "weightrecord",
				"localField":   "_id",
				"foreignField": "box_id",
				"as":           "weights",
			},
		},
		{
			"$unwind": bson.M{
				"path":                       "$weights",
				"preserveNullAndEmptyArrays": true,
			},
		},
		{
			"$group": bson.M{
				"_id": "$type",
				"weight": bson.M{
					"$sum": "$weights.weight",
				},
			},
		},
		{
			"$project": bson.M{
				"type":   "$_id",
				"weight": 1,
			},
		},
	}
	cur, err := col.Aggregate(ctx, filter)
	if err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.DatabaseErrorCode)
	}
	var result []vo.StatBoxWeightVo
	for cur.Next(ctx) {
		var row vo.StatBoxWeightVo
		if err := cur.Decode(&row); err != nil {
			syslog.Error(err)
			exception.ThrowBusinessError(common.DatabaseErrorCode)
		}
		result = append(result, row)
	}
	return result
}
