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

// AddTag Add fish tag
func AddTag(tag entitys.FishTag) primitive.ObjectID {
	col, ctx := Collection("fishtag")
	filter := bson.D{
		primitive.E{Key: "name", Value: tag.Name},
	}
	cur, err := col.Find(ctx, filter)
	if err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.DatabaseErrorCode)
	}
	if cur.Next(ctx) {
		exception.ThrowBusinessErrorMsg("该Tag信息已存在，请更改...")
	}
	timeNow := time.Now()
	tag.CreateTime = timeNow
	tag.UpdateTime = timeNow
	tag.Creator = ""
	tag.Used = false

	result, err := col.InsertOne(ctx, tag)
	if err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.DatabaseErrorCode)
	}
	return result.InsertedID.(primitive.ObjectID)
}

// SelectTagByID Select tag by id
func SelectTagByID(id primitive.ObjectID) entitys.FishTag {
	col, ctx := Collection("fishtag")
	filter := bson.D{
		primitive.E{Key: "_id", Value: id},
	}
	cur, err := col.Find(ctx, filter)
	if err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.DatabaseErrorCode)
	}
	var tag entitys.FishTag
	if cur.Next(ctx) {
		if err := cur.Decode(&tag); err != nil {
			syslog.Error(err)
			exception.ThrowBusinessError(common.DatabaseErrorCode)
		}
	}
	return tag
}

// SelectTagByName Select tag by tag name
func SelectTagByName(name string) (primitive.ObjectID, bool) {
	col, ctx := Collection("fishtag")
	filter := bson.D{
		primitive.E{
			Key:   "name",
			Value: name,
		},
	}
	cur, err := col.Find(ctx, filter)
	if err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.DatabaseErrorCode)
	}
	var tag entitys.FishTag
	if cur.Next(ctx) {
		if err := cur.Decode(&tag); err != nil {
			syslog.Error(err)
			exception.ThrowBusinessError(common.DatabaseErrorCode)
		}
	}
	return tag.ID, tag.Used
}

// UpdateFishTagStatus update fish tag status
func UpdateFishTagStatus(id primitive.ObjectID, used bool) {
	col, ctx := Collection("fishtag")
	filter := bson.D{
		primitive.E{Key: "_id", Value: id},
	}
	update := bson.D{
		primitive.E{
			Key: "$set", Value: bson.M{
				"used":          used,
				"update_time":   time.Now(),
				"last_operator": "",
			},
		},
	}
	if _, err := col.UpdateOne(ctx, filter, update); err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.DatabaseErrorCode)
	}
}
