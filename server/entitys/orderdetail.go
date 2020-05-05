package entitys

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// OrderDetail OrderDetail entity
type OrderDetail struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Orderid      primitive.ObjectID
	Product      string
	Item         string
	Quantity     int
	Boxtype      string
	EstLbsPerBox int `bson:"est_lbs_per_box,omitempty"`
	TotalLbs     int `bson:"total_lbs,omitempty"`
	User         string
	Time         time.Time
}
