package dto

// StatDto Stat request dto
type StatDto struct {
	TaskID string `form:"task_id" validate:"required"`
}
