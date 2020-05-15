package dto

// AddBoxDto Add singl one request
type AddBoxDto struct {
	Type string `validate:"required"`
	Num string `validate:"required"`
}