package models

import (
	"mt-scale/common"
	"mt-scale/entitys"
	"mt-scale/exception"
	"mt-scale/models/vo"
	"mt-scale/syslog"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddBox Add new box
func AddBox(box entitys.Box) primitive.ObjectID {
	col, ctx := Collection("box")
	filter := bson.D{
		primitive.E{Key: "type", Value: box.Type},
		primitive.E{Key: "num", Value: box.Num},
	}
	cur, err := col.Find(ctx, filter)
	if err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.DatabaseErrorCode)
	}
	if cur.Next(ctx) {
		exception.ThrowBusinessErrorMsg("该号码的箱子已存在，请更改...")
	}
	timeNow := time.Now()
	box.CreateTime = timeNow
	box.UpdateTime = timeNow

	result, err := col.InsertOne(ctx, box)
	if err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.DatabaseErrorCode)
	}
	insertedID := result.InsertedID.(primitive.ObjectID)
	syslog.Debug(">>>>>> insertedID:", insertedID)
	return insertedID
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
func FetchBoxes(boxType string) []entitys.Box {
	col, ctx := Collection("box")
	filter := []bson.M{
		{
			"$match": bson.M{
				"status": "enable",
				"type":   boxType,
			},
		},
		{"$sort": bson.M{
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
