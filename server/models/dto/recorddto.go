package dto

// QueryRecordDto query weight record
type QueryRecordDto struct {
	BoxID     string `form:"box_id"`
	SpeciesID string `form:"species_id"`
	TaskID    string `form:"task_id" validate:"required"`
	Index     int    `form:"index" validate:"integer"`
	PageNum   int    `form:"page_num" validate:"integer"`
	PageSize  int    `form:"page_size" validate:"integer"`
}

// DelRecordDto entity
type DelRecordDto struct {
	ID string `validate:"required" bson:"id" json:"id"`
}
