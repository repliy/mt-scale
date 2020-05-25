package dto

// LoginDto login dto
type LoginDto struct {
	Username string `validate:"required" bson:"username" json:"username"`
	Password string `validate:"required" bson:"password" json:"password"`
}
