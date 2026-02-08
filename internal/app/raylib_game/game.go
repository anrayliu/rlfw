package raylib_game

import (
	core "anray/raylib-game/internal/pkg/core"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct{}

func (g *Game) Enter(e *core.Engine) {
	e.Run(&Foo{})
}

func (g *Game) Exit(e *core.Engine) {
}

func (g *Game) Update(e *core.Engine) {
}

func (g *Game) Draw(e *core.Engine) {
	rl.ClearBackground(rl.Red)
}

type Foo struct{}

func (g *Foo) Enter(e *core.Engine) {
}

func (g *Foo) Exit(e *core.Engine) {
}

func (g *Foo) Update(e *core.Engine) {
	if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
		e.QuitAll()
	}
}

func (g *Foo) Draw(e *core.Engine) {
	rl.ClearBackground(rl.Blue)
}
