package vo

// StatBoxWeightVo Vo
type StatBoxWeightVo struct {
	Type   string  `json:"type"`
	Weight float32 `json:"weight"`
	Name   string  `json:"name"`
}

// StatSpecWeightVo Statistical species weight vo
type StatSpecWeightVo struct {
	Name   string  `json:"name"`
	Weight float32 `json:"weight"`
}

// StatTotalWeightVo Statis total weight
type StatTotalWeightVo struct {
	Weight float32 `json:"weight" bson:"weight"`
}

// StatWeightVo Statis weight
type StatWeightVo struct {
	Species []StatSpecWeightVo `json:"species"`
	Box     []StatBoxWeightVo  `json:"box"`
	Total   StatTotalWeightVo  `json:"total"`
}
