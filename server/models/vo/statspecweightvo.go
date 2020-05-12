package vo

import "go.mongodb.org/mongo-driver/bson/primitive"

// StatSpecWeightVo Statistical species weight vo
type StatSpecWeightVo struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string             `json:"name"`
	Tag    string             `json:"tag"`
	Weight float32            `json:"weight"`
}
