package main

import (
	"github.com/mateo-tavera/max-inventory/config"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(config.New), //pass all that return a struct
		fx.Invoke(),            //command to execute before starting,
	)

	app.Run()
}
