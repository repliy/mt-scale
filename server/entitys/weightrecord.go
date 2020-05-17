package entitys

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// WeightRecord entity
type WeightRecord struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Weight       float32            `validate:"numeric" json:"weight"`
	BoxID        primitive.ObjectID `validate:"required" bson:"box_id" json:"box_id"`
	SpeciesID    primitive.ObjectID `validate:"required" bson:"species_id" json:"species_id"`
	TagID        primitive.ObjectID `bson:"tag_id,omitempty" json:"tag_id,omitempty"`
	TaskID       primitive.ObjectID `validate:"required" bson:"task_id" json:"task_id"`
	Index        int                `validate:"gte=0" json:"index"`
	CreateTime   time.Time          `bson:"create_time" json:"create_time"`
	UpdateTime   time.Time          `bson:"update_time" json:"update_time"`
	Creator      string             `json:"creator"`
	LastOperator string             `bson:"last_operator" json:"last_operator"`
}
