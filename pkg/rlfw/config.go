package rlfw

import rl "github.com/gen2brain/raylib-go/raylib"

type Config struct {
	WinW     int32
	WinH     int32
	WinMode  uint32 // rl window flags
	Name     string
	Fps      int32
	LogLevel rl.TraceLogLevel
}
