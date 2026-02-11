package main

import (
	app "github.com/anrayliu/rlfw/internal/app"
	rlfw "github.com/anrayliu/rlfw/pkg/rlfw"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	engine, err := rlfw.NewEngine(rlfw.Config{
		WinW:     800,
		WinH:     600,
		WinMode:  0,
		Name:     "example",
		Fps:      60,
		LogLevel: rl.LogDebug,
	})
	if err != nil {
		panic(err)
	}

	engine.Run(&app.Game{})
}
