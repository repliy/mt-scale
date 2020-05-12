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
