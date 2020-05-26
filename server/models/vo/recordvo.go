package vo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// WeightRecordVo weight page display
type WeightRecordVo struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Weight       float32            `json:"weight"`
	SpeciesID    primitive.ObjectID `bson:"species_id" json:"species_id"`
	Species      string             `bson:"species" json:"species"`
	SpeciesColor string             `bson:"species_color" json:"species_color"`
	TagID        primitive.ObjectID `bson:"tag_id" json:"tag_id"`
	Tag          string             `bson:"tag" json:"tag"`
	BoxID        primitive.ObjectID `bson:"box_id" json:"box_id"`
	BoxType      string             `bson:"box_type" json:"box_type"`
	BoxNum       string             `bson:"box_num" json:"box_num"`
	Index        int                `bson:"index" json:"index"`
}
