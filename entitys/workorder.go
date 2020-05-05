package entitys

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// WorkOrder WorkOrder entity
type WorkOrder struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Number      string
	PackingDate string  `bson:"packing_date,omitempty"`
	CommonNotes string  `bson:"common_notes,omitempty"`
	Order       []Order `bson:"order,omitempty"`
	Status      string
	User        string
	Time        time.Time
}
