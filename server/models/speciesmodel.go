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

// AddSpecies Add species tag
func AddSpecies(spec entitys.Species) primitive.ObjectID {
	if spec.Tag == "" {
		spec.Tag = "default"
	}
	col, ctx := Collection("species")
	filter := bson.D{
		primitive.E{Key: "name", Value: spec.Name},
		primitive.E{Key: "tag", Value: spec.Tag},
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
	insertedID := result.InsertedID.(primitive.ObjectID)
	syslog.Debug(">>>>>> insertedID:", insertedID)
	return insertedID
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
	if cur.Next(ctx) {
		var row entitys.Species
		if err := cur.Decode(&row); err != nil {
			syslog.Error(err)
			exception.ThrowBusinessError(common.DatabaseErrorCode)
		}
		result = append(result, row)
	}
	return result
}
