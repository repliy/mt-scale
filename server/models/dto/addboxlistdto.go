package dto

// AddBoxListDto Add box with a request list
type AddBoxListDto struct {
	BoxList []AddBoxDto `validate:"required"`
}