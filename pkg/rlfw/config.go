package rlfw

import rl "github.com/gen2brain/raylib-go/raylib"

type Config struct {
	WinW       int32
	WinH       int32
	WinMode    uint32 // rl window flags
	Name       string
	Fps        int32
	LogLevel   rl.TraceLogLevel
	LoadAssets bool
}

func DefaultConfig() Config {
	return Config{
		WinW:       800,
		WinH:       600,
		WinMode:    0,
		Name:       "example",
		Fps:        60,
		LogLevel:   rl.LogDebug,
		LoadAssets: true,
	}
}
