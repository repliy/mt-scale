package models

import (
	"mt-scale/common"
	"mt-scale/entitys"
	"mt-scale/exception"
	"mt-scale/models/dto"
	"mt-scale/syslog"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddTask Add a new task
func AddTask() primitive.ObjectID {
	col, ctx, cancel := Collection("task")
	defer cancel()
	task := entitys.Task{
		Creator:    "admin",
		CreateTime: time.Now(),
		Status:     "start",
	}
	result, err := col.InsertOne(ctx, task)
	if err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.DatabaseErrorCode)
	}
	return result.InsertedID.(primitive.ObjectID)
}

// GetCurrentTask Get current task ID
func GetCurrentTask() primitive.ObjectID {
	col, ctx, cancel := Collection("task")
	defer cancel()
	filter := []bson.M{
		{
			"$match": primitive.M{
				"status": "start",
			},
		},
		{
			"$sort": bson.M{
				"create_time": -1,
			},
		},
		{
			"$limit": 1,
		},
	}
	cur, err := col.Aggregate(ctx, filter)
	if err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.DatabaseErrorCode)
	}
	if cur.Next(ctx) {
		var row entitys.Task
		if err := cur.Decode(&row); err != nil {
			syslog.Error(err)
			exception.ThrowBusinessError(common.DatabaseErrorCode)
		}
		return row.ID
	}
	return AddTask()
}

// UpdTaskStatus Update task status(start/complete)
func UpdTaskStatus(dto dto.TaskUpdDto) {
	col, ctx, cancel := Collection("task")
	defer cancel()
	filter := bson.D{
		primitive.E{
			Key:   "_id",
			Value: dto.ID,
		},
	}
	update := bson.D{
		primitive.E{
			Key: "$set",
			Value: bson.M{
				"status":        dto.Status,
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
