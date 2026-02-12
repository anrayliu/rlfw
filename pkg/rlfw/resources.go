package rlfw

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
	// generate default image for missing assets
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

	_, ok := r.images[base]
	if ok {
		// already loaded, but this is not an error
		// see readme for api design philosophy
		return nil
	}

	if ext == ".png" || ext == ".jpg" {
		r.images[base] = rl.LoadImage(path)
		return nil
	}

	return errors.New("unsupported file format")
}

func (r *Resources) UnloadImg(pathOrName string) error {
	img, ok := r.images[pathOrName]
	if !ok {
		base, _, err := splitFileName(pathOrName)
		if err != nil {
			return err
		}
		img, ok = r.images[base]
		if !ok {
			// nothing to unload, but this is not an error
			// see readme for api design philosophy
			return nil
		}
	}

	rl.UnloadImage(img)

	return nil
}

func (r *Resources) LoadTexture(path string) error {
	base, ext, err := splitFileName(path)
	if err != nil {
		return err
	}

	_, ok := r.textures[base]
	if ok {
		// already loaded, but this is not an error
		// see readme for api design philosophy
		return nil
	}

	if ext == ".png" || ext == ".jpg" {
		img, ok := r.images[base]
		if !ok {
			// temporarily load image to create texture with
			img = rl.LoadImage(path)
			defer rl.UnloadImage(img)
		}

		r.textures[base] = rl.LoadTextureFromImage(img)
		return nil
	}

	return errors.New("unsupported file format")
}

func (r *Resources) UnloadTexture(pathOrName string) error {
	texture, ok := r.textures[pathOrName]
	if !ok {
		base, _, err := splitFileName(pathOrName)
		if err != nil {
			return err
		}
		texture, ok = r.textures[base]
		if !ok {
			// nothing to unload, but this is not an error
			// see readme for api design philosophy
			return nil
		}
	}

	rl.UnloadTexture(texture)

	return nil
}

func (r *Resources) LoadFont(path string) error {
	base, ext, err := splitFileName(path)
	if err != nil {
		return err
	}

	_, ok := r.fonts[base]
	if ok {
		// already loaded, but this is not an error
		// see readme for api design philosophy
		return nil
	}

	if ext == ".ttf" || ext == ".otf" {
		r.fonts[base] = rl.LoadFont(path)
		return nil
	}

	return errors.New("unsupported file format")
}

func (r *Resources) UnloadFont(pathOrName string) error {
	font, ok := r.fonts[pathOrName]
	if !ok {
		base, _, err := splitFileName(pathOrName)
		if err != nil {
			return err
		}
		font, ok = r.fonts[base]
		if !ok {
			// nothing to unload, but this is not an error
			// see readme for api design philosophy
			return nil
		}
	}

	rl.UnloadFont(font)

	return nil
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
			log.Printf("loading image: %s", path)
			err := r.LoadImg(path)
			if err != nil {
				return err
			}
			log.Printf("loading texture: %s", path)
			return r.LoadTexture(path)
		}
		if ext == ".ttf" || ext == ".otf" {
			log.Printf("loading font: %s", path)
			return r.LoadFont(path)
		}
		return nil
	})
}

func (r *Resources) GetTexture(name string) (rl.Texture2D, bool) {
	texture, ok := r.textures[name]
	if !ok {
		return r.defaultTexture, false
	}
	return texture, true
}

func (r *Resources) GetImg(name string) (*rl.Image, bool) {
	img, ok := r.images[name]
	if !ok {
		return r.defaultImg, false
	}
	return img, true
}

func (r *Resources) GetFont(name string) (rl.Font, bool) {
	font, ok := r.fonts[name]
	if !ok {
		return rl.GetFontDefault(), false
	}
	return font, true
}
