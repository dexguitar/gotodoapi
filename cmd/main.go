package main

import (
	"github.com/dexguitar/gotodoapi/config"
	"github.com/dexguitar/gotodoapi/internal/app"
)

func main() {
	c, err := config.MustLoad()
	if err != nil {
		panic(err)
	}

	app, err := app.New(c)
	if err != nil {
		panic(err)
	}
	app.Router.Run()
}
