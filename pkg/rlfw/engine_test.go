package rlfw

import (
	"testing"
)

type testState struct{}

func (s *testState) Draw(e *Engine)   {}
func (s *testState) Update(e *Engine) {}
func (s *testState) Enter(e *Engine) {
	e.QuitState()
}
func (s *testState) Exit(e *Engine)   {}
func (s *testState) Resize(e *Engine) {}

type testDefaultState struct {
	DefaultState
}

func (s *testDefaultState) Enter(e *Engine) {
	e.QuitState()
}

func createTestEngine() *Engine {
	engine, err := NewEngine(DefaultConfig())
	if err != nil {
		panic(err)
	}
	return engine
}

func TestDefaultState(t *testing.T) {
	e := createTestEngine()
	e.Run(&testDefaultState{})
}

func TestCustomState(t *testing.T) {
	e := createTestEngine()
	e.Run(&testState{})
}
