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
	"go.mongodb.org/mongo-driver/mongo/options"
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
func FetchWeightRecord(dto dto.QueryRecordDto) []entitys.WeightRecord {
	col, ctx := Collection("weightrecord")
	pageNum := dto.PageNum
	if pageNum == 0 {
		pageNum = 1
	}
	pageSize := dto.PageSize
	if pageSize == 0 {
		pageSize = 10
	}
	filter := bson.D{}
	if dto.BoxID != "" {
		objID, _ := primitive.ObjectIDFromHex(dto.BoxID)
		filter = append(filter, primitive.E{Key: "box_id", Value: objID})
	}
	if dto.SpeciesID != "" {
		specID, _ := primitive.ObjectIDFromHex(dto.SpeciesID)
		filter = append(filter, primitive.E{Key: "species_id", Value: specID})
	}
	if dto.Index != 0 {
		filter = append(filter, primitive.E{Key: "index", Value: dto.Index})
	}
	// find opitons
	opts := new(options.FindOptions)
	limit := int64(pageSize)
	skip := int64((pageNum - 1) * pageSize)
	sortMap := make(map[string]interface{})
	sortMap["create_time"] = -1
	opts.Sort = sortMap
	opts.Limit = &limit
	opts.Skip = &skip

	cur, err := col.Find(ctx, filter, opts)
	if err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.DatabaseErrorCode)
	}
	var result []entitys.WeightRecord
	for cur.Next(ctx) {
		var row entitys.WeightRecord
		if err := cur.Decode(&row); err != nil {
			syslog.Error(err)
			exception.ThrowBusinessError(common.DatabaseErrorCode)
		}
		result = append(result, row)
	}
	return result
}
