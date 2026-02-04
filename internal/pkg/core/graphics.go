package internal

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type graphics struct {
	images   map[string]*rl.Image
	textures map[string]rl.Texture2D
}

func newGraphics() *graphics {
	return &graphics{
		images:   map[string]*rl.Image{},
		textures: map[string]rl.Texture2D{},
	}
}

func (g *graphics) LoadImg(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		return errors.New("file does not exist")
	}

	ext := filepath.Ext(path)
	if ext != ".png" {
		return errors.New("unsupported file format")
	}

	filename := strings.TrimSuffix(filepath.Base(path), ext)

	g.images[filename] = rl.LoadImage(path)
	g.textures[filename] = rl.LoadTextureFromImage(g.images[filename])

	return nil
}

func (g *graphics) LoadDir(dir string) error {
	return filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		return g.LoadImg(path)
	})
}

func (g *graphics) DrawImg(image string, x int32, y int32) error {
	texture, ok := g.textures[image]
	if !ok {
		return errors.New("image not found")
	}

	rl.DrawTexture(texture, x, y, rl.White)

	return nil
}

func (g *graphics) DrawImgRect(image string, sourceRect rl.Rectangle) error {
	texture, ok := g.textures[image]
	if !ok {
		return errors.New("image not found")
	}

	rl.DrawTexturePro(texture,
		rl.Rectangle{0, 0, float32(texture.Width), float32(texture.Height)},
		sourceRect,
		rl.Vector2{0, 0},
		0.0,
		rl.White,
	)

	return nil
}
