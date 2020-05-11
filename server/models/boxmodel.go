package models

import (
	"mt-scale/entitys"
	"mt-scale/syslog"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddBox Add new box
func AddBox(box entitys.Box) primitive.ObjectID {
	col, ctx := Collection("box")
	result, err := col.InsertOne(ctx, box)
	if err != nil {
		syslog.Error(">>>>>> inserted box error:", err)
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
	}
	cur, err := col.Aggregate(ctx, filter)
	if err != nil {
		syslog.Error(">>>>>> fetch boxes error:", err)
	}
	var result []entitys.Box
	for cur.Next(ctx) {
		var row entitys.Box
		err := cur.Decode(&row)
		if err != nil {
			syslog.Error(">>>>>> fetch boxes error:", err)
		}
		result = append(result, row)
	}
	return result
}
