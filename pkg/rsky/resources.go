package core

import (
	"errors"
	"io/fs"
	"log"
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

func splitFileName(path string) (string, string, error) {
	_, err := os.Stat(path)
	if err != nil {
		return "", "", errors.New("file does not exist")
	}

	ext := filepath.Ext(path)
	filename := strings.TrimSuffix(filepath.Base(path), ext)

	return filename, ext, nil
}

func (r *Resources) LoadImg(path string) error {
	base, ext, err := splitFileName(path)
	if err != nil {
		return err
	}

	if ext == ".png" || ext == ".jpg" {
		r.images[base] = rl.LoadImage(path)
		r.textures[base] = rl.LoadTextureFromImage(r.images[base])
		return nil
	}

	return errors.New("unsupported file format")
}

func (r *Resources) LoadFont(path string) error {
	base, ext, err := splitFileName(path)
	if err != nil {
		return err
	}

	if ext == ".ttf" || ext == ".otf" {
		r.fonts[base] = rl.LoadFont(path)
		return nil
	}

	return errors.New("unsupported file format")
}

func (r *Resources) LoadDir(dir string) error {
	return filepath.WalkDir(dir, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			log.Printf("error while walking over %s: %s", path, err)
		}
		if entry.IsDir() {
			return nil
		}
		ext := filepath.Ext(path)
		if ext == ".png" || ext == ".jpg" {
			return r.LoadImg(path)
		}
		if ext == ".ttf" || ext == ".otf" {
			return r.LoadFont(path)
		}
		return nil
	})
}

func (r *Resources) GetTexture(name string) rl.Texture2D {
	texture, ok := r.textures[name]
	if !ok {
		log.Printf("texture not found: %s", name)
		return r.defaultTexture
	}
	return texture
}

func (r *Resources) GetImg(name string) *rl.Image {
	img, ok := r.images[name]
	if !ok {
		log.Printf("image not found: %s", name)
		return r.defaultImg
	}
	return img
}

func (r *Resources) GetFont(name string) rl.Font {
	font, ok := r.fonts[name]
	if !ok {
		log.Printf("font not found: %s", name)
		return rl.GetFontDefault()
	}
	return font
}

func (r *Resources) cleanUp() {
	for _, img := range r.images {
		rl.UnloadImage(img)
	}
	for _, texture := range r.textures {
		rl.UnloadTexture(texture)
	}
	for _, font := range r.fonts {
		rl.UnloadFont(font)
	}
}
