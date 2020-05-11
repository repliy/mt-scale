package dto

// QueryBoxDto Query boxes request params
type QueryBoxDto struct {
	BoxType string `form:"box_type" validate:"required"`
}
