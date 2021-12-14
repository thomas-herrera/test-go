package handlers

import (

	"github.com/mercadolibre/fury_go-core/pkg/web"
	"github.com/mercadolibre/fury_go-platform/pkg/fury"
)

func Api(app *fury.Application, calculator Calculator) {

	app.Router.Post("/do", Calculate(calculator), web.AcceptJSON())
}