package rlfw

import (
	"bytes"
	"log"
	"strings"
	"testing"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/stretchr/testify/assert"
)

func createTestEngineWithConfig(cfg Config) *Engine {
	engine, err := NewEngine(cfg)
	if err != nil {
		panic(err)
	}
	return engine
}

func createTestEngine() *Engine {
	cfg := DefaultConfig()
	cfg.LoadAssets = false
	cfg.LogLevel = rl.LogFatal
	return createTestEngineWithConfig(cfg)
}

func TestWindowCreation(t *testing.T) {
	cfg := DefaultConfig()
	cfg.LogLevel = rl.LogFatal
	cfg.WinMode = rl.FlagWindowResizable
	cfg.LoadAssets = false

	createTestEngineWithConfig(cfg)

	assert.Equal(t, cfg.WinW, int32(rl.GetScreenWidth()))
	assert.Equal(t, cfg.WinH, int32(rl.GetScreenHeight()))
	assert.True(t, rl.IsWindowState(cfg.WinMode))

	// not all the window information can be retrived for testing
}

func TestBadWindowSize(t *testing.T) {
	cfg := DefaultConfig()
	cfg.LoadAssets = false
	cfg.LogLevel = rl.LogFatal

	cfg.WinW = 0
	_, err := NewEngine(cfg)
	assert.NotNil(t, err)

	cfg.WinW = 100
	cfg.WinH = 0
	_, err = NewEngine(cfg)
	assert.NotNil(t, err)

	cfg.WinW = 100
	cfg.WinH = 100
	_, err = NewEngine(cfg)
	assert.Nil(t, err)

	cfg.WinW = 0
	cfg.WinH = 0
	_, err = NewEngine(cfg)
	assert.NotNil(t, err)
}

func TestEngineInitValues(t *testing.T) {
	engine := createTestEngine()
	assert.False(t, engine.quit)
	assert.False(t, engine.quitAll)
	assert.Equal(t, len(engine.states), 0)
}

type testDefaultState struct {
	DefaultState
}

func (s *testDefaultState) Enter(e *Engine) {
	e.QuitState()
}

func TestDefaultState(t *testing.T) {
	e := createTestEngine()
	e.Run(&testDefaultState{})
}

func TestEngineLoadAssets(t *testing.T) {
	var buf bytes.Buffer

	orig := log.Writer()
	log.SetOutput(&buf)
	defer log.SetOutput(orig)

	_ = createTestEngine()

	assert.Equal(t, buf.Len(), 0)

	cfg := DefaultConfig()
	cfg.LogLevel = rl.LogFatal
	cfg.LoadAssets = true

	_ = createTestEngineWithConfig(cfg)

	assert.NotEqual(t, buf.Len(), 0)
	assert.True(t, strings.HasSuffix(buf.String(), "error while walking over assets: lstat assets: no such file or directory\n"))
}

func TestQuitState(t *testing.T) {
	engine := createTestEngine()
	engine.QuitState()
	assert.True(t, engine.quit)
	assert.False(t, engine.quitAll)
}

func TestQuitApp(t *testing.T) {
	engine := createTestEngine()
	engine.QuitApp()
	assert.True(t, engine.quit)
	assert.True(t, engine.quitAll)
}

type testState struct {
	eventsTriggered int
}

func (s *testState) Draw(e *Engine) {
	s.eventsTriggered++
	e.QuitState()
}
func (s *testState) Update(e *Engine) {
	s.eventsTriggered++
	e.resizeStates()
}
func (s *testState) Enter(e *Engine) {
	s.eventsTriggered++
}
func (s *testState) Exit(e *Engine) {
	s.eventsTriggered++
}
func (s *testState) Resize(e *Engine) {
	s.eventsTriggered++
}

func TestCallbacks(t *testing.T) {
	engine := createTestEngine()
	state := &testState{}

	engine.Run(state)

	assert.Equal(t, state.eventsTriggered, 5)
}
