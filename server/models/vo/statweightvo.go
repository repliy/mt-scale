package vo

// StatWeightVo Statis weight
type StatWeightVo struct {
	Species []StatSpecWeightVo `json:"species"`
	Box     []StatBoxWeightVo  `json:"box"`
	Total   StatTotalWeightVo  `json:"total"`
}
