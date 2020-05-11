package dto

// QueryBoxDto Query boxes request params
type QueryBoxDto struct {
	BoxType string `json:"box_type" validate:"required"`
}
