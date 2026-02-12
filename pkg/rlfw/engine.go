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

	// window dimensions don't matter if fullscreen
	if (cfg.WinW <= 0 || cfg.WinH <= 0) && (cfg.WinMode&rl.FlagFullscreenMode) == 0 {
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
		states:  NewStack[State](),
	}, nil
}

type Engine struct {
	Resources *Resources
	Cfg       Config

	states *Stack[State]

	// internal state management vars
	quit    bool
	quitAll bool
}

func (e *Engine) resizeStates() {
	for i := e.states.Len() - 1; i >= 0; i-- {
		e.states.GetSlice()[i].Resize(e)
	}
}

func (e *Engine) Run(state State) {
	e.states.Push(state)
	if e.states.Len() == 1 {
		defer rl.CloseWindow()
		defer e.Resources.cleanUp()
	}
	defer func() {
		state.Exit(e)
		e.states.Pop()

		if !e.quitAll {
			e.quit = false
		}
	}()

	state.Enter(e)

	for {
		if rl.WindowShouldClose() {
			e.QuitAll()
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

func (e *Engine) Quit() {
	e.quit = true
}

func (e *Engine) QuitAll() {
	e.quitAll = true
	e.Quit()
}
