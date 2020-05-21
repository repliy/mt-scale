package dto

// QueryBoxDto Query boxes request params
type QueryBoxDto struct {
	BoxType string `form:"box_type" validate:"required"`
	TaskID  string `form:"task_id" validate:"required" bson:"task_id" json:"task_id"`
}

// QueryLatestBoxDto Query boxes
type QueryLatestBoxDto struct {
	TaskID string `form:"task_id" validate:"required" json:"task_id"`
}

// AddBoxDto Add singl one request
type AddBoxDto struct {
	Type   string `validate:"required"`
	Num    string `validate:"required"`
	TaskID string `validate:"required" bson:"task_id" json:"task_id"`
}

// AddBoxListDto Add box with a request list
type AddBoxListDto struct {
	BoxList []AddBoxDto `validate:"required" json:"box_list"`
}

// QueryTallyBoxDto Query boxes
type QueryTallyBoxDto struct {
	TaskID string `form:"task_id" validate:"required" json:"task_id"`
}
