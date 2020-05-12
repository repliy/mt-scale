package entitys

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Species Species
type Species struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name         string             `validate:"required" json:"name"`
	Tag          string             `json:"tag"`
	CreateTime   time.Time          `bson:"create_time" json:"create_time"`
	UpdateTime   time.Time          `bson:"update_time" json:"update_time"`
	Creator      string             `json:"creator"`
	LastOperator string             `bson:"last_operator" json:"last_operator"`
}
