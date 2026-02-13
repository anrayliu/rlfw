package rlfw

import (
	"errors"
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewEngine(cfg Config) (*Engine, error) {
	rl.SetTraceLogLevel(cfg.LogLevel)
	rl.SetTargetFPS(cfg.Fps)

	rl.SetConfigFlags(cfg.WinMode)

	if cfg.WinW <= 0 || cfg.WinH <= 0 {
		return nil, errors.New("bad window size")
	}

	rl.InitWindow(cfg.WinW, cfg.WinH, cfg.Name)

	resources := newResources()
	if cfg.LoadAssets {
		err := resources.LoadDir("assets")
		if err != nil {
			log.Printf("error when loading assets folder: %s", err)
		}
	}

	return &Engine{
		Resources: resources,
		Cfg:       cfg,

		quit:    false,
		quitAll: false,

		states: []State{},
	}, nil
}

type Engine struct {
	Resources *Resources
	Cfg       Config

	states []State

	// internal state management vars
	quit    bool
	quitAll bool
}

func (e *Engine) resizeStates() {
	for i := len(e.states) - 1; i >= 0; i-- {
		e.states[i].Resize(e)
	}
}

func (e *Engine) Run(state State) {
	e.states = append(e.states, state)
	if len(e.states) == 1 {
		defer rl.CloseWindow()
		defer e.Resources.cleanUp()
	}
	defer func() {
		state.Exit(e)
		e.states = e.states[:len(e.states)-1] // pop from end of slice

		if !e.quitAll {
			e.quit = false
		}
	}()

	state.Enter(e)

	for {
		if rl.WindowShouldClose() {
			e.QuitApp()
		} else if e.quit {
			break
		}

		if rl.IsWindowResized() {
			e.resizeStates()
		}

		state.Update(e)

		rl.ClearBackground(rl.White)

		rl.BeginDrawing()

		state.Draw(e)

		rl.EndDrawing()
	}
}

func (e *Engine) QuitState() {
	e.quit = true
}

func (e *Engine) QuitApp() {
	e.quitAll = true
	e.QuitState()
}
