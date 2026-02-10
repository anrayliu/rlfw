package main

import (
	app "anray/raylib-game/internal/app/raylib_game"
	core "anray/raylib-game/internal/pkg/core"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	engine, err := core.NewEngine(core.Config{
		WinW:     800,
		WinH:     600,
		WinMode:  0,
		Name:     "test",
		Fps:      60,
		LogLevel: rl.LogDebug,
	})
	if err != nil {
		panic(err)
	}

	engine.Run(&app.Game{})
}
