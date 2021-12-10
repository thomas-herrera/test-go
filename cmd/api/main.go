package main

import (
	"main/cmd/api/handlers"
	"main/internal/calculator"

	"github.com/mercadolibre/fury_go-platform/pkg/fury"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	app, err := fury.NewWebApplication()
	if err != nil {
		return err
	}

	calculator := calculator.NewCalculator()

	handlers.Api(app, calculator)

	return app.Run()
}