package dto

// QueryBoxDto Query boxes request params
type QueryBoxDto struct {
	BoxType string `form:"box_type" validate:"required"`
}

// AddBoxDto Add singl one request
type AddBoxDto struct {
	Type string `validate:"required"`
	Num  string `validate:"required"`
}

// AddBoxListDto Add box with a request list
type AddBoxListDto struct {
	BoxList []AddBoxDto `validate:"required" json:"box_list"`
}
