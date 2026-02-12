package main

import (
	app "github.com/anrayliu/rlfw/internal/app"
	rlfw "github.com/anrayliu/rlfw/pkg/rlfw"
)

func main() {
	engine, err := rlfw.NewEngine(rlfw.DefaultConfig())
	if err != nil {
		panic(err)
	}

	engine.Run(&app.Game{})
}
