package entitys

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FishTag fish tag
type FishTag struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name         string             `validate:"required"`
	SpeciesID    primitive.ObjectID `validate:"required" bson:"species_id" json:"species_id"`
	Used         bool
	CreateTime   time.Time `bson:"create_time" json:"create_time"`
	UpdateTime   time.Time `bson:"update_time" json:"update_time"`
	Creator      string    `json:"creator"`
	LastOperator string    `bson:"last_operator" json:"last_operator"`
}
