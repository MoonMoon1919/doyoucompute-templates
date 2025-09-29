package main

import (
	"os"

	"github.com/MoonMoon1919/doyoucompute"
	"github.com/MoonMoon1919/doyoucompute-templates/internal/docs"
	"github.com/MoonMoon1919/doyoucompute-templates/pkg/bugreport"
	"github.com/MoonMoon1919/doyoucompute-templates/pkg/contributing"
	"github.com/MoonMoon1919/doyoucompute-templates/pkg/pullrequest"
	"github.com/MoonMoon1919/doyoucompute/pkg/app"
)

func main() {
	// TODO: Vend "DefaultApp" that offers "WithService" in main pkg
	svc, err := doyoucompute.DefaultService()
	if err != nil {
		panic(err)
	}

	app := app.New(svc)

	bugreport, err := bugreport.New()
	if err != nil {
		panic(err)
	}

	pullrequest, err := pullrequest.New()
	if err != nil {
		panic(err)
	}

	contributing, err := contributing.New(
		"https://github.com/MoonMoon1919/doyoucompute-templates",
		"https://github.com/MoonMoon1919/doyoucompute-templates/issues",
	)
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
