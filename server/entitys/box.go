package entitys

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Box Box entity
type Box struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Type         string             `validate:"required"`
	Num          string             `validate:"required"`
	Status       string             `json:"status"`
	TaskID       primitive.ObjectID `validate:"required" bson:"task_id" json:"task_id"`
	CreateTime   time.Time          `bson:"create_time" json:"create_time"`
	UpdateTime   time.Time          `bson:"update_time" json:"update_time"`
	Creator      string             `json:"creator"`
	LastOperator string             `bson:"last_operator" json:"last_operator"`
}
