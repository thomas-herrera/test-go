package main

import (
	"main/cmd/api/handlers"

	"github.com/mercadolibre/fury_go-platform/pkg/fury"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

var MemoryMap map[string]float64

func run() error {
	app, err := fury.NewWebApplication()
	if err != nil {
		return err
	}
	
	MemoryMap = make(map[string]float64)
	MemoryMap["blanco"] = 11
	MemoryMap["rojo"] = 4
	MemoryMap["azul"] = 25
	MemoryMap["negro"] = 20
	MemoryMap["amarillo"] = 7

	handlers.Api(app, MemoryMap)

	return app.Run()
}