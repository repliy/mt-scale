package dto

// StatDto Stat request dto
type StatDto struct {
	TaskID string `validate:"required" json:"task_id"`
}
