package vo

import "go.mongodb.org/mongo-driver/bson/primitive"

// AuthTokenVo auth token object
type AuthTokenVo struct {
	AccessToken string `bson:"access_token" json:"access_token"`
}

// LoginVo login response vo
type LoginVo struct {
	ID       primitive.ObjectID `bson:"id" json:"id"`
	Username string             `bson:"username" json:"username"`
	Token    AuthTokenVo        `bson:"token" json:"token"`
}
