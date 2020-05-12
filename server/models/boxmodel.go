package models

import (
	"mt-scale/common"
	"mt-scale/entitys"
	"mt-scale/exception"
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
		err := cur.Decode(&row)
		if err != nil {
			syslog.Error(err)
			exception.ThrowBusinessError(common.DatabaseErrorCode)
		}
		result = append(result, row)
	}
	return result
}
