package vo

import (
	"mt-scale/entitys"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddBoxVo Add box return object
type AddBoxVo struct {
	BoxID   primitive.ObjectID `json:"box_id" bson:"box_id"`
	BoxType string             `json:"box_type" bson:"box_type"`
	BoxNum  string             `json:"box_num" bson:"box_num"`
}

// BoxValidateErrorVo Return when validate errors happen
type BoxValidateErrorVo struct {
	BoxType     string `json:"box_type"`
	BoxNum      string `json:"box_num"`
	ValidateMsg string `json:"validate_msg"`
}

// BoxTallyVo Vessel Plant Tally vo
type BoxTallyVo struct {
	Type        string
	Num         string
	TotalWeight float32 `json:"total_weight"`
	Weights     []entitys.WeightRecord
}
