package internal

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Resources struct {
	images   map[string]*rl.Image
	textures map[string]rl.Texture2D
	fonts    map[string]rl.Font

	defaultImg     *rl.Image
	defaultTexture rl.Texture2D
}

func newResources() *Resources {
	tmp := rl.GenImageColor(256, 256, rl.Red)
	rl.ImageDrawTextEx(tmp, rl.Vector2{X: 0, Y: 0}, rl.GetFontDefault(), "missing resource", 25, 1, rl.Black)

	return &Resources{
		images:   map[string]*rl.Image{},
		textures: map[string]rl.Texture2D{},
		fonts:    map[string]rl.Font{},

		defaultImg:     tmp,
		defaultTexture: rl.LoadTextureFromImage(tmp),
	}
}

func (r *Resources) LoadImg(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		return errors.New("file does not exist")
	}

	ext := filepath.Ext(path)
	if ext != ".png" {
		return errors.New("unsupported file format")
	}

	filename := strings.TrimSuffix(filepath.Base(path), ext)

	r.images[filename] = rl.LoadImage(path)
	r.textures[filename] = rl.LoadTextureFromImage(r.images[filename])

	return nil
}

func (r *Resources) LoadDir(dir string) error {
	return filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		return r.LoadImg(path)
	})
}

func (r *Resources) GetTexture(name string) rl.Texture2D {
	texture, ok := r.textures[name]
	if !ok {
		return r.defaultTexture
	}
	return texture
}

func (r *Resources) GetImg(name string) *rl.Image {
	img, ok := r.images[name]
	if !ok {
		return r.defaultImg
	}
	return img
}

func (r *Resources) GetFont(name string) rl.Font {
	font, ok := r.fonts[name]
	if !ok {
		return rl.GetFontDefault()
	}
	return font
}
