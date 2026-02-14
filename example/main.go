package main

import (
	rlfw "github.com/anrayliu/rlfw/pkg/rlfw"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	rlfw.DefaultState
}

func (g *Game) Draw(e *rlfw.Engine) {
	rl.ClearBackground(rl.Blue)
}

func main() {
	engine, err := rlfw.NewEngine(rlfw.DefaultConfig())
	if err != nil {
		panic(err)
	}

	engine.Run(&Game{})
}
