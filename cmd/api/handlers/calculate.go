package handlers

import (
	"main/internal/calculator"
	"main/internal/storage"
	"net/http"

	"github.com/mercadolibre/fury_go-core/pkg/web"
)

type Calculator interface {
	GetResult(operator string, operands []float64, storage calculator.Storage) (float64, error)
}


func Calculate(calculator Calculator) web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		var req RequestCalculate
		var res ResponseCalculate
		if err := web.DecodeJSON(r, &req); err != nil {
			return err
		}
		storage := storage.NewFile()
		response, err := calculator.GetResult(req.Operator, req.Operands, storage)
		if err != nil {
			return err
		}
		res.Result = response
		return web.EncodeJSON(w, res, http.StatusOK)
	}
}
