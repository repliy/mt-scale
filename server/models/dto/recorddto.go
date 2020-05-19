package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

// QueryRecordDto query weight record
type QueryRecordDto struct {
	BoxID     string `form:"box_id"`
	SpeciesID string `form:"species_id"`
	TaskID    string `form:"task_id" validate:"required"`
	Index     int    `form:"index" validate:"integer"`
	PageNum   int    `form:"page_num" validate:"integer"`
	PageSize  int    `form:"page_size" validate:"integer"`
}

// DelRecordDto entity
type DelRecordDto struct {
	ID string `validate:"required" bson:"id" json:"id"`
}

// AddWeightRecordDto entity
type AddWeightRecordDto struct {
	Weight    float32            `validate:"numeric" json:"weight"`
	BoxID     primitive.ObjectID `validate:"required" bson:"box_id" json:"box_id"`
	SpeciesID primitive.ObjectID `validate:"required" bson:"species_id" json:"species_id"`
	TaskID    primitive.ObjectID `validate:"required" bson:"task_id" json:"task_id"`
	Index     int                `validate:"gte=0" json:"index"`
	TagName   string             `bson:"tag_name" json:"tag_name"`
}
