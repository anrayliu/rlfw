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
	fonts    map[string]rl.Font

	defaultImg     *rl.Image
	defaultTexture rl.Texture2D
}

func newGraphics() *graphics {
	tmp := rl.GenImageColor(256, 256, rl.Red)

	return &graphics{
		images:   map[string]*rl.Image{},
		textures: map[string]rl.Texture2D{},
		fonts:    map[string]rl.Font{},

		defaultImg:     tmp,
		defaultTexture: rl.LoadTextureFromImage(tmp),
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

func (g *graphics) GetTexture(name string) rl.Texture2D {
	texture, ok := g.textures[name]
	if !ok {
		return g.defaultTexture
	}
	return texture
}

func (g *graphics) GetImg(name string) *rl.Image {
	img, ok := g.images[name]
	if !ok {
		return g.defaultImg
	}
	return img
}

func (g *graphics) GetFont(name string) rl.Font {
	font, ok := g.fonts[name]
	if !ok {
		return rl.GetFontDefault()
	}
	return font
}
