package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

// TaskUpdDto Update task status dto
type TaskUpdDto struct {
	ID     primitive.ObjectID `json:"id" validate:"required"`
	Status string             `validate:"required"`
}
