package rlfw

import rl "github.com/gen2brain/raylib-go/raylib"

// Config stores the application settings.
// Used during engine creation.
type Config struct {
	// WinW is the window width in pixels.
	WinW int32
	// WinH is the window height in pixels.
	WinH int32
	// WinMode contains raylib window flags for display configuration (e.g rl.FlagBorderlessWindowedMode)
	WinMode uint32
	// Name is the window title.
	Name string
	// Fps is the target frames per second for the game loop.
	Fps int32
	// LogLevel sets the raylib logging verbosity level.
	LogLevel rl.TraceLogLevel
	// LoadAssets specifies whether to automatically load assets on engine creation.
	// If true, the "assets" directory (relative to the working directory) is loaded.
	LoadAssets bool
}

// DefaultConfig returns a Config struct with default values.
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
