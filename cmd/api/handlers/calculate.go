package handlers

import (
	"main/internal/calculator"
	"net/http"

	"github.com/mercadolibre/fury_go-core/pkg/web"
)

func Calculate() web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		var req RequestCalculate
		if err := web.DecodeJSON(r, &req); err != nil {
			return err
		}
		response, err := calculator.Calculator(req.Operator, req.Operands)
		if err != nil {
			return err
		}

		return web.EncodeJSON(w, response, http.StatusOK)
	}
}

func CalculateMemory(memoryMap map[string]float64) web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		var req RequestMemory
		name := web.Params(r)["name"]
		if err := web.DecodeJSON(r, &req); err != nil {
			return err
		}
		response := calculator.CalculatorMemory(memoryMap, name, req.Add, req.Value)
		return web.EncodeJSON(w, response, http.StatusOK)
	}
}

func GetMemory(memoryMap map[string]float64) web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		name := web.Params(r)["name"]
		response := calculator.GetCalculatorMemory(name, memoryMap)
		return web.EncodeJSON(w, response, http.StatusOK)
	}
}