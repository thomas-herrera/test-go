package calculator

import (
	"errors"
)

func GetResult(operator string, operands []float64) (float64, error) {
	var result float64
	var error error
	for i, v := range operands {
		number:= v
		if i==0 {
			result = v
			number = 0
		}
		switch operator {
			case "add":
				result += number
			case "substract":
				result -= number
			case "divide":
				if(number == 0) {
					error = errors.New("impossible division by zero")
				}
				result /= number
			case "multiply":
				result *= number
			default:
				error = errors.New("undefined operator")
		}
	}
	return result, error
}

func GetCalculatorMemory(name string, memoryMap map[string]float64) float64{
	return memoryMap[name]
}

func CalculatorMemory(memoryMap map[string]float64, name string, add bool, value float64) string{
	if add {
		memoryMap[name] += value
	} else {
		memoryMap[name] -= value
	}
	result := "Memory called " + name + " update successfully"
	return result
}
