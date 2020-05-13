package vo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// WeightRecordVo weight page display
type WeightRecordVo struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Weight  float32            `json:"weight"`
	Species string             ` json:"species"`
	Tag     string             ` json:"tags"`
	Index   int                `json:"index"`
}
