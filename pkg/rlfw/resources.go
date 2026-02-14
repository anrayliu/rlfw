package rlfw

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Resources stores loaded images, textures, and fonts.
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
	for name := range r.images {
		r.UnloadImg(name)
	}

	for name := range r.textures {
		r.UnloadTexture(name)
	}

	for name := range r.fonts {
		r.UnloadFont(name)
	}
}

func splitFileName(path string) (string, string, error) {
	_, err := os.Stat(path)
	if err != nil {
		return "", "", fmt.Errorf("file %s does not exist", path)
	}

	ext := filepath.Ext(path)
	filename := strings.TrimSuffix(filepath.Base(path), ext)

	return filename, ext, nil
}

// LoadImg loads an image from the given file path and stores it.
// Supported file formats are .png and .jpg.
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

	return fmt.Errorf("unsupported file format: %s", ext)
}

// UnloadImg deletes the stored image by file path or resource name.
func (r *Resources) UnloadImg(pathOrName string) error {
	img, ok := r.images[pathOrName]
	if !ok {
		base, _, err := splitFileName(pathOrName)
		if err != nil {
			return err
		}
		img, ok = r.images[base]
		pathOrName = base
		if !ok {
			// nothing to unload, but this is not an error
			// see readme for api design philosophy
			return nil
		}
	}

	rl.UnloadImage(img)
	delete(r.images, pathOrName)

	return nil
}

// LoadTexture loads a texture from the given file path and stores it.
// Supported file formats are .png and .jpg.
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

	return fmt.Errorf("unsupported file format: %s", ext)
}

// UnloadTexture deletes the stored texture by file path or resource name.
func (r *Resources) UnloadTexture(pathOrName string) error {
	texture, ok := r.textures[pathOrName]
	if !ok {
		base, _, err := splitFileName(pathOrName)
		if err != nil {
			return err
		}
		texture, ok = r.textures[base]
		pathOrName = base
		if !ok {
			// nothing to unload, but this is not an error
			// see readme for api design philosophy
			return nil
		}
	}

	rl.UnloadTexture(texture)
	delete(r.textures, pathOrName)

	return nil
}

// LoadFont loads a font from the given file path and stores it.
// Supported file formats are .ttf and .otf.
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

	return fmt.Errorf("unsupported file format: %s", ext)
}

// UnloadFont deletes the stored font by file path or resource name.
func (r *Resources) UnloadFont(pathOrName string) error {
	font, ok := r.fonts[pathOrName]
	if !ok {
		base, _, err := splitFileName(pathOrName)
		if err != nil {
			return err
		}
		font, ok = r.fonts[base]
		pathOrName = base
		if !ok {
			// nothing to unload, but this is not an error
			// see readme for api design philosophy
			return nil
		}
	}

	rl.UnloadFont(font)
	delete(r.fonts, pathOrName)

	return nil
}

// LoadDir recursively loads all supported resources from the given directory.
func (r *Resources) LoadDir(dir string) error {
	return filepath.WalkDir(dir, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			log.Printf("error while walking over %s: %s", path, err)
			return nil
		}
		if entry.IsDir() {
			return nil
		}
		ext := filepath.Ext(path)
		switch ext {
		case ".png", ".jpg":
			log.Printf("loading image: %s", path)
			err := r.LoadImg(path)
			if err != nil {
				return err
			}
			log.Printf("loading texture: %s", path)
			return r.LoadTexture(path)
		case ".ttf", ".otf":
			log.Printf("loading font: %s", path)
			return r.LoadFont(path)
		}
		return nil
	})
}

// UnloadDir recursively unloads all resources loaded from the given directory.
func (r *Resources) UnloadDir(dir string) error {
	return filepath.WalkDir(dir, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			log.Printf("error while walking over %s: %s", path, err)
			return nil
		}
		if entry.IsDir() {
			return nil
		}
		ext := filepath.Ext(path)
		// errors shouldn't be possible, so throw away
		switch ext {
		case ".png", ".jpg":
			_ = r.UnloadImg(path)
			_ = r.UnloadTexture(path)
		case ".ttf", ".otf":
			_ = r.UnloadFont(path)
		}
		return nil
	})
}

// GetTexture retrieves a texture by name.
// Returns a second value: a bool that's true if the referenced texture exists.
func (r *Resources) GetTexture(name string) (rl.Texture2D, bool) {
	texture, ok := r.textures[name]
	if !ok {
		return r.defaultTexture, false
	}
	return texture, true
}

// GetImg retrieves an image by name.
// Returns a second value: a bool that's true if the referenced image exists.
func (r *Resources) GetImg(name string) (*rl.Image, bool) {
	img, ok := r.images[name]
	if !ok {
		return r.defaultImg, false
	}
	return img, true
}

// GetFont retrieves a font by name.
// Returns a second value: a bool that's true if the referenced font exists.
func (r *Resources) GetFont(name string) (rl.Font, bool) {
	font, ok := r.fonts[name]
	if !ok {
		return rl.GetFontDefault(), false
	}
	return font, true
}
