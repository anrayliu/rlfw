package core

import (
	"errors"
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type State interface {
	Enter(e *Engine)
	Exit(e *Engine)
	Update(e *Engine)
	Draw(e *Engine)
}

func NewEngine(cfg Config) (*Engine, error) {
	rl.SetTraceLogLevel(cfg.LogLevel)
	rl.SetTargetFPS(cfg.Fps)

	rl.SetConfigFlags(cfg.WinMode)

	// window dimensions don't matter if fullscreen
	if (cfg.WinW <= 0 || cfg.WinH <= 0) && (cfg.WinMode&rl.FlagFullscreenMode) == 0 {
		return nil, errors.New("bad window size")
	}

	rl.InitWindow(cfg.WinW, cfg.WinH, cfg.Name)

	resources := newResources()
	err := resources.LoadDir("assets")
	if err != nil {
		log.Printf("error when loading assets folder: %s", err)
	}

	return &Engine{
		Resources: resources,
		Cfg:       cfg,

		quit:       false,
		quitAll:    false,
		firstState: true,
	}, nil
}

type Engine struct {
	Resources *Resources
	Cfg       Config

	// internal state management vars
	quit       bool
	quitAll    bool
	firstState bool // only defer window close for first state
}

func (e *Engine) Run(state State) {
	if e.firstState {
		defer rl.CloseWindow()
		defer e.Resources.cleanUp()
		e.firstState = false
	}
	defer func() {
		state.Exit(e)
		if !e.quitAll {
			e.quit = false
		}
	}()

	state.Enter(e)

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

func (e *Engine) QuitAll() {
	e.quitAll = true
	e.Quit()
}
