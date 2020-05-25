package vo

import "go.mongodb.org/mongo-driver/bson/primitive"

// LoginVo login response vo
type LoginVo struct {
	ID          primitive.ObjectID `bson:"id" json:"id"`
	Username    string             `bson:"username" json:"username"`
	AccessToken string             `bson:"access_token" json:"access_token"`
}
