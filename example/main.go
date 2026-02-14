package main

import (
	rlfw "github.com/anrayliu/rlfw/pkg/rlfw"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	rlfw.DefaultState
}

func (g *Game) Draw(e *rlfw.Engine) {
	r := e.Resources

	rl.ClearBackground(rl.Blue)
	rl.DrawTexture(r.GetTexture("apple"), 0, 0, rl.White)
}

func main() {
	engine, err := rlfw.NewEngine(rlfw.DefaultConfig())
	if err != nil {
		panic(err)
	}

	engine.Run(&Game{})
}
