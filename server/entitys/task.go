package entitys

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Task entity
type Task struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Creator    string             `json:"creator"`
	CreateTime time.Time          `bson:"create_time" json:"create_time"`
	Status     string             `json:"status"`
}
