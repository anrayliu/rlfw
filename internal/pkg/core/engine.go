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

type State interface {
	Start(e *Engine)
	Update(e *Engine)
	Draw(e *Engine)
	End(e *Engine)
}

func NewEngine(cfg Config) (*Engine, error) {
	rl.SetTraceLogLevel(cfg.LogLevel)
	rl.SetTargetFPS(cfg.Fps)

	if cfg.WinW == 0 && cfg.WinH == 0 {
		rl.SetConfigFlags(rl.FlagFullscreenMode)
	} else if cfg.WinW <= 0 || cfg.WinH <= 0 {
		return nil, errors.New("bad window size")
	}

	rl.InitWindow(cfg.WinW, cfg.WinH, cfg.Name)

	resources := newResources()
	err := resources.LoadDir("assets")
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		return nil, err
	}

	return &Engine{
		Resources: resources,
		Cfg:       cfg,

		quit:       false,
		firstState: true,
	}, nil
}

type Engine struct {
	Resources *Resources
	Cfg       Config

	quit       bool
	firstState bool // only defer window close for first state
}

func (e *Engine) Run(state State) {
	if e.firstState {
		defer rl.CloseWindow()
		e.firstState = false
	}
	defer func() {
		state.End(e)
		e.quit = false
	}()

	e.quit = false

	state.Start(e)

	for !rl.WindowShouldClose() && !e.quit {
		state.Update(e)

		rl.ClearBackground(rl.White)

		rl.BeginDrawing()

		state.Draw(e)

		rl.EndDrawing()
	}
}

func (e *Engine) Quit() {
	e.quit = true
}
