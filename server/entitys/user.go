package entitys

import "go.mongodb.org/mongo-driver/bson/primitive"

// User entity
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username string             `validate:"required" bson:"username" json:"username"`
	Password string             `validate:"required" bson:"password" json:"password"`
}
