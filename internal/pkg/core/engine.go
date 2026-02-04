package internal

import (
	"errors"
	"io/fs"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Config struct {
	WinW     int32
	WinH     int32
	Name     string
	Fps      int32
	LogLevel rl.TraceLogLevel
}

type Runnable interface {
	Init(e *Engine)
	Update(e *Engine)
	Draw(e *Engine)
}

func NewEngine(cfg Config) (*Engine, error) {
	rl.SetTraceLogLevel(cfg.LogLevel)

	win, err := newWin(cfg.WinW, cfg.WinH, cfg.Name)
	if err != nil {
		return nil, err
	}

	graphics := newGraphics()
	err = graphics.LoadDir("assets")
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		return nil, err
	}

	engine := &Engine{
		Win:      win,
		Graphics: graphics,
		Cfg:      cfg,
	}

	rl.SetTargetFPS(cfg.Fps)

	return engine, nil
}

type Engine struct {
	Win      *window
	Graphics *graphics
	Cfg      Config
}

func (e *Engine) Run(r Runnable) {
	defer e.Win.close()

	r.Init(e)

	for !rl.WindowShouldClose() {
		r.Update(e)

		rl.BeginDrawing()

		r.Draw(e)

		rl.EndDrawing()
	}
}
