package dto

// DelWeightDto entity
type DelWeightDto struct {
	ID string `validate:"required" bson:"id" json:"id"`
}
