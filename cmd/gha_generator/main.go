package main

import (
	app "anray/gha-generator/internal/app/gha_generator"
	core "anray/gha-generator/internal/pkg/core"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	engine, err := core.NewEngine(core.Config{
		WinW:     800,
		WinH:     600,
		Name:     "test",
		Fps:      60,
		LogLevel: rl.LogDebug,
	})
	if err != nil {
		panic(err)
	}

	engine.Run(&app.Game{})
}
