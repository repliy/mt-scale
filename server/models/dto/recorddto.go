package dto

// QueryRecordDto query weight record
type QueryRecordDto struct {
	BoxID     string `form:"box_id"`
	SpeciesID string `form:"species_id"`
	Index     int    `validate:"integer" form:"index"`
	PageNum   int    `validate:"integer" form:"page_num"`
	PageSize  int    `validate:"integer" form:"page_size"`
}
