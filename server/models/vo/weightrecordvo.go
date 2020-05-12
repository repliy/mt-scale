package vo

import (
	"mt-scale/entitys"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// WeightRecordVo weight page display
type WeightRecordVo struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Weight  float32            `json:"weight"`
	Species []entitys.Species  `bson:",omitempty" json:"species"`
	Index   int                `json:"index"`
}
