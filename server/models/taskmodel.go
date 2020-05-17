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

// AddTask Add a new task
func AddTask() primitive.ObjectID {
	col, ctx := Collection("task")
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
	col, ctx := Collection("task")
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
