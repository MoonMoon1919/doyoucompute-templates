package main

import (
	"os"

	"github.com/MoonMoon1919/doyoucompute-templates/internal/docs"
	"github.com/MoonMoon1919/doyoucompute-templates/pkg/bugreport"
	"github.com/MoonMoon1919/doyoucompute-templates/pkg/pullrequest"
	"github.com/MoonMoon1919/doyoucompute/pkg/app"
)

func main() {
	app := app.Default()

	bugreport, err := bugreport.New()
	if err != nil {
		panic(err)
	}

	pullrequest, err := pullrequest.New()
	if err != nil {
		panic(err)
	}

	contributing, err := docs.Contributing()
	if err != nil {
		panic(err)
	}

	readme, err := docs.ReadMe()
	if err != nil {
		panic(err)
	}

	app.Register(readme)
	app.Register(bugreport)
	app.Register(pullrequest)
	app.Register(contributing)

	app.Run(os.Args)
}
