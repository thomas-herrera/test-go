package handlers

type RequestCalculate struct {
	Operator string `json:"operator" validate:"required"`
	Operands []float64 `json:"operands" validate:"required"`
}

type RequestMemory struct {
	Add bool `json:"add" validate:"required"`
	Value float64 `json:"value" validate:"required"`
}

type ResponseCalculate struct {
	Result float64 `json:"result"`
}