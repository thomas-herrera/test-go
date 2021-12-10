package calculator

import (
	"errors"
)

type Calculator struct {
	memory float64
	memoryMap map[string]float64
}

func NewCalculator() *Calculator {
	calculatorMap := make(map[string]float64)
	calculatorMap["blanco"] = 11
	calculatorMap["rojo"] = 4
	calculatorMap["azul"] = 25
	calculatorMap["negro"] = 20
	calculatorMap["amarillo"] = 7
	return &Calculator{memory: 0, memoryMap: calculatorMap}
}

func (c *Calculator) GetResult(operator string, operands []float64) (float64, error) {
	var result float64
	var error error
	switch operator {
	case "m+":
		c.memory++
	case "m-":
		c.memory--
	case "get":
		result = c.memory 
	case "reset":
		c.memory = 0 
	default:
		for i, v := range operands {
			number:= v
			if i==0 {
				result = v
				number = 0
			} else {
				switch operator {
					case "add":
						result += number
					case "substract":
						result -= number
					case "multiply":
						result *= number
					case "divide":
						if(number == 0) {
							result = 0
							error = errors.New("impossible division by zero")
						} else {
							result /= number
						}
					default:
						result = 0
						error = errors.New("undefined operator")
				}
			}
		}
	}
	return result, error
}

func (c *Calculator) GetCalculatorMemory(name string) float64{
	return c.memoryMap[name]
}

func (c *Calculator) ModifyMemory(name string, add bool, value float64) string{
	if add {
		c.memoryMap[name] += value
	} else {
		c.memoryMap[name] -= value
	}
	result := "Memory called " + name + " update successfully"
	return result
}
