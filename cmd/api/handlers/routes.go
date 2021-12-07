package handlers

import (

	"github.com/mercadolibre/fury_go-core/pkg/web"
	"github.com/mercadolibre/fury_go-platform/pkg/fury"
)

func Api(app *fury.Application, memoryMap map[string]float64) {

	app.Router.Post("/do", Calculate(), web.AcceptJSON())
	app.Router.Post("/memory/{name}", CalculateMemory(memoryMap), web.AcceptJSON())
	app.Router.Get("/memory/{name}", GetMemory(memoryMap), web.AcceptJSON())
}