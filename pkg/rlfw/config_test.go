package rlfw

import (
	"reflect"
	"testing"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/stretchr/testify/assert"
)

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()

	assert.Equal(t, cfg.WinW, int32(800))
	assert.Equal(t, cfg.WinH, int32(600))
	assert.Equal(t, cfg.WinMode, uint32(0))
	assert.Equal(t, cfg.Name, "example")
	assert.Equal(t, cfg.Fps, int32(60))
	assert.Equal(t, cfg.LogLevel, rl.LogDebug)
	assert.Equal(t, cfg.LoadAssets, true)

	assert.Equal(t, reflect.TypeOf(cfg).NumField(), 7)
}
