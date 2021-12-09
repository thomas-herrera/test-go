package handlers

import (
	"main/internal/calculator"
	"net/http"

	"github.com/mercadolibre/fury_go-core/pkg/web"
)

var FuncCalculator = calculator.GetResult
var FuncCalculatorMemory = calculator.CalculatorMemory
var FuncGetCalculatorMemory = calculator.GetCalculatorMemory

func Calculate() web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		var req RequestCalculate
		var res ResponseCalculate
		if err := web.DecodeJSON(r, &req); err != nil {
			return err
		}
		response, err := FuncCalculator(req.Operator, req.Operands)
		if err != nil {
			return err
		}
		res.Result = response
		return web.EncodeJSON(w, res, http.StatusOK)
	}
}

func CalculateMemory(memoryMap map[string]float64) web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		var req RequestMemory
		name := web.Params(r)["name"]
		if err := web.DecodeJSON(r, &req); err != nil {
			return err
		}
		response := FuncCalculatorMemory(memoryMap, name, req.Add, req.Value)
		return web.EncodeJSON(w, response, http.StatusOK)
	}
}

func GetMemory(memoryMap map[string]float64) web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		name := web.Params(r)["name"]
		response := FuncGetCalculatorMemory(name, memoryMap)
		return web.EncodeJSON(w, response, http.StatusOK)
	}
}