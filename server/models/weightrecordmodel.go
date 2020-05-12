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

// AddWeightRecord record weight one time
func AddWeightRecord(record entitys.WeightRecord) primitive.ObjectID {
	col, ctx := Collection("weightrecord")
	filter := bson.D{
		primitive.E{Key: "index", Value: record.Index},
	}
	cur, err := col.Find(ctx, filter)
	if err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.DatabaseErrorCode)
	}
	if cur.Next(ctx) {
		// override record first
		var row entitys.WeightRecord
		if err := cur.Decode(&row); err != nil {
			syslog.Error(err)
			exception.ThrowBusinessError(common.DatabaseErrorCode)
		}
		return UpdWeightRecord(row)
	}
	record.CreateTime = time.Now()
	record.Creator = ""
	result, err := col.InsertOne(ctx, record)
	if err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.DatabaseErrorCode)
	}
	return result.InsertedID.(primitive.ObjectID)
}

// UpdWeightRecord update weight record by index
func UpdWeightRecord(record entitys.WeightRecord) primitive.ObjectID {
	col, ctx := Collection("weightrecord")
	filter := bson.D{primitive.E{Key: "index", Value: record.Index}}
	update := bson.D{
		primitive.E{
			Key: "$set", Value: bson.M{
				"weight":        record.Weight,
				"box_id":        record.BoxID,
				"species_id":    record.SpeciesID,
				"update_time":   time.Now(),
				"last_operator": "",
			},
		},
	}
	if _, err := col.UpdateOne(ctx, filter, update); err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.DatabaseErrorCode)
	}
	return record.ID
}
