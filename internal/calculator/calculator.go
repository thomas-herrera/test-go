package calculator

import (
	"errors"
	"strconv"
)

func Calculator(operator string, operands []string) (CalculatorResult, error) {
	var result CalculatorResult
	var error error
	temp, err := strconv.ParseFloat(operands[0], 64)
	if err != nil {
		return CalculatorResult{}, err
	}
	for i := 1; i < len(operands); i++ {
		number, err := strconv.ParseFloat(operands[i], 64)
		if err != nil {
			return CalculatorResult{}, err
		}
		switch operator {
			case "add":
				temp += number
			case "substract":
				temp -= number
			case "divide":
				if(number == 0) {
					error = errors.New("impossible division by zero")
				}
				temp /= number
			case "multiply":
				temp *= number
			default:
				error = errors.New("undefined operator")
		}
	}
	result.Result = temp
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
