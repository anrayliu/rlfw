package internal

import (
	"errors"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func newWin(w int32, h int32, name string) (*window, error) {
	if w == 0 && h == 0 {
		rl.SetConfigFlags(rl.FlagFullscreenMode)
	} else if w <= 0 || h <= 0 {
		return nil, errors.New("bad window size")
	}

	rl.InitWindow(w, h, name)

	return &window{
		W: int32(rl.GetScreenWidth()),
		H: int32(rl.GetScreenHeight()),
	}, nil
}

type window struct {
	W int32
	H int32
}

func (win *window) close() {
	rl.CloseWindow()
}
