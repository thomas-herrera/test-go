package handlers

type RequestCalculate struct {
	Operator string `json:"operator" validate:"required"`
	Operands []string `json:"operands" validate:"required"`
}

type RequestMemory struct {
	Add bool `json:"add" validate:"required"`
	Value float64 `json:"value" validate:"required"`
}