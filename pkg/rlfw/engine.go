package rlfw

import (
	"errors"
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// NewEngine creates and returns an Engine with the given configuration.
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

// Engine is the core driver of rlfw.
// It handles application boilerplate, loads and stores resources, and manages state.
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

// Run places the given state on the stack and passes control to it.
func (e *Engine) Run(state State) {
	e.states = append(e.states, state)
	if len(e.states) == 1 {
		defer rl.CloseWindow()
		defer e.Resources.cleanUp()
	}
	defer func() {
		state.Exit(e)
		e.states = e.states[:len(e.states)-1]

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

// QuitState exits from the current state, passing control back to the previous state.
func (e *Engine) QuitState() {
	e.quit = true
}

// QuitApp exits from all states, terminating the engine.
func (e *Engine) QuitApp() {
	e.quitAll = true
	e.QuitState()
}
