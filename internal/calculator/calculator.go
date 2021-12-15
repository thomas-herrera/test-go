package calculator

import (
	"errors"
	"main/internal/storage"
)

type Calculator struct {
	memory float64
	storage Storage
}

func New() *Calculator {
	storage := storage.NewFile()
	return &Calculator{memory: 0, storage: storage}
}

type Storage interface {
	Save(result float64) (bool, error)
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
	c.storage.Save(result)
	return result, error
}
