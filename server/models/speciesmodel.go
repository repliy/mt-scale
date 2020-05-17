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

// AddSpecies Add species tag
func AddSpecies(spec entitys.Species) primitive.ObjectID {
	col, ctx := Collection("species")
	filter := bson.D{
		primitive.E{Key: "name", Value: spec.Name},
	}
	cur, err := col.Find(ctx, filter)
	if err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.DatabaseErrorCode)
	}
	if cur.Next(ctx) {
		exception.ThrowBusinessErrorMsg("该物种信息已存在，请更改...")
	}
	timeNow := time.Now()
	spec.CreateTime = timeNow
	spec.UpdateTime = timeNow

	result, err := col.InsertOne(ctx, spec)
	if err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.DatabaseErrorCode)
	}
	return result.InsertedID.(primitive.ObjectID)
}

// SelectSpeciesByID Select species by id
func SelectSpeciesByID(id primitive.ObjectID) entitys.Species {
	col, ctx := Collection("species")
	filter := bson.D{
		primitive.E{Key: "_id", Value: id},
	}
	cur, err := col.Find(ctx, filter)
	if err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.DatabaseErrorCode)
	}
	var spec entitys.Species
	if cur.Next(ctx) {
		if err := cur.Decode(&spec); err != nil {
			syslog.Error(err)
			exception.ThrowBusinessError(common.DatabaseErrorCode)
		}
	}
	return spec
}

// FetchAllSpeciesInfo Fetch all species infos
func FetchAllSpeciesInfo() []entitys.Species {
	col, ctx := Collection("species")
	filter := bson.D{}
	cur, err := col.Find(ctx, filter)
	if err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.DatabaseErrorCode)
	}
	var result []entitys.Species
	for cur.Next(ctx) {
		var row entitys.Species
		if err := cur.Decode(&row); err != nil {
			syslog.Error(err)
			exception.ThrowBusinessError(common.DatabaseErrorCode)
		}
		result = append(result, row)
	}
	return result
}

// StatSpecieszWeight Statistical weighing record
func StatSpecieszWeight(taskID primitive.ObjectID) []vo.StatSpecWeightVo {
	col, ctx := Collection("species")
	filter := []bson.M{
		{
			"$lookup": bson.M{
				"from":         "weightrecord",
				"localField":   "_id",
				"foreignField": "species_id",
				"as":           "weights",
			},
		},
		{
			"$match": bson.M{
				"task_id": taskID,
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
				"_id": "$name",
				"weight": bson.M{
					"$sum": "$weights.weight",
				},
			},
		},
		{
			"$project": bson.M{
				"name":   "$_id",
				"weight": 1,
			},
		},
	}
	cur, err := col.Aggregate(ctx, filter)
	if err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.DatabaseErrorCode)
	}
	var result []vo.StatSpecWeightVo
	for cur.Next(ctx) {
		var row vo.StatSpecWeightVo
		if err := cur.Decode(&row); err != nil {
			syslog.Error(err)
			exception.ThrowBusinessError(common.DatabaseErrorCode)
		}
		result = append(result, row)
	}
	return result
}
