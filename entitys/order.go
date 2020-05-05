package entitys

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Order Order entity
type Order struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Number            string
	Orderby           string
	Orderdate         string
	Shipto            string
	Departure         string
	AWB               string
	Destination       string
	Airline           string
	Currency          string
	HealthCertificate bool          `bson:"health_certificate,omitempty"`
	PackingNotes      string        `bson:"packing_notes,omitempty"`
	Detail            []OrderDetail `bson:"detail,omitempty"`
	TotalLbs          int           `bson:"total_lbs,-"`
	TotalBoxes        int           `bson:"total_boxes,-"`
	Status            string
	Workorderid       primitive.ObjectID `bson:",omitempty" json:",omitempty"`
	WorkOrder         []WorkOrder        `bson:",omitempty" json:",omitempty"`
	LogisticLoggerID  string             `bson:"logisticloggerid,omitempty"`
	Truck             string
	PickUpTime        string
	User              string
	Time              time.Time
}
