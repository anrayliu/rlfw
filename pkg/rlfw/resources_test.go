package rlfw

import (
	"testing"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/stretchr/testify/assert"
)

const applePath = "../../assets/apple.png"
const fontpath = "../../assets/arial.ttf"
const dirPath = "../../assets"

func TestNewResources(t *testing.T) {
	r := createTestEngine().Resources

	assert.Equal(t, len(r.fonts), 0)
	assert.Equal(t, len(r.images), 0)
	assert.Equal(t, len(r.textures), 0)

	assert.Equal(t, r.defaultImg.Width, int32(256))
	assert.Equal(t, r.defaultImg.Height, int32(256))

	assert.Equal(t, r.defaultTexture.Width, int32(256))
	assert.Equal(t, r.defaultTexture.Height, int32(256))
}

func TestLoadDir(t *testing.T) {
	r := createTestEngine().Resources

	r.LoadDir(dirPath)

	assert.Equal(t, len(r.images), 1)
	assert.Equal(t, len(r.textures), 1)
	assert.Equal(t, len(r.fonts), 1)

	_, ok := r.images["apple"]
	assert.True(t, ok)

	_, ok = r.textures["apple"]
	assert.True(t, ok)
}

func TestUnloadDir(t *testing.T) {
	r := createTestEngine().Resources

	r.LoadDir(dirPath)
	r.UnloadDir(dirPath)

	assert.Equal(t, len(r.images), 0)
	assert.Equal(t, len(r.textures), 0)
	assert.Equal(t, len(r.fonts), 0)
}

func TestSplitFileName(t *testing.T) {
	base, ext, err := splitFileName(applePath)
	assert.Nil(t, err)
	assert.Equal(t, base, "apple")
	assert.Equal(t, ext, ".png")

	_, _, err = splitFileName("badpath")
	assert.NotNil(t, err)
}

func TestLoadImg(t *testing.T) {
	r := createTestEngine().Resources

	r.LoadImg(applePath)

	assert.Equal(t, len(r.images), 1)

	image, ok := r.images["apple"]
	assert.True(t, ok)

	assert.Equal(t, image.Width, int32(256))
	assert.Equal(t, image.Height, int32(256))

	err := r.LoadImg(applePath)
	assert.Nil(t, err)

	err = r.LoadImg(fontpath)
	assert.NotNil(t, err)
}

func TestUnloadImg(t *testing.T) {
	r := createTestEngine().Resources

	r.LoadImg(applePath)

	err := r.UnloadImg(applePath)

	_, ok := r.images["apple"]
	assert.False(t, ok)

	assert.Equal(t, len(r.images), 0)

	r.LoadImg(applePath)

	r.UnloadImg("apple")

	_, ok = r.images["apple"]
	assert.False(t, ok)

	assert.Equal(t, len(r.images), 0)

	err = r.UnloadImg("badpath")
	assert.NotNil(t, err)
}

func TestLoadTexture(t *testing.T) {
	r := createTestEngine().Resources

	r.LoadTexture(applePath)

	assert.Equal(t, len(r.textures), 1)

	texture, ok := r.textures["apple"]
	assert.True(t, ok)

	assert.Equal(t, texture.Width, int32(256))
	assert.Equal(t, texture.Height, int32(256))

	err := r.LoadTexture(applePath)
	assert.Nil(t, err)

	err = r.LoadTexture(fontpath)
	assert.NotNil(t, err)
}

func TestUnloadTexture(t *testing.T) {
	r := createTestEngine().Resources

	r.LoadTexture(applePath)

	err := r.UnloadTexture(applePath)

	_, ok := r.textures["apple"]
	assert.False(t, ok)

	assert.Equal(t, len(r.textures), 0)

	r.LoadTexture(applePath)

	r.UnloadTexture("apple")

	_, ok = r.textures["apple"]
	assert.False(t, ok)

	assert.Equal(t, len(r.textures), 0)

	err = r.UnloadTexture("badpath")
	assert.NotNil(t, err)
}

func TestLoadFont(t *testing.T) {
	r := createTestEngine().Resources

	r.LoadFont(fontpath)

	assert.Equal(t, len(r.fonts), 1)

	font, ok := r.fonts["arial"]
	assert.True(t, ok)

	assert.Equal(t, font.CharsCount, int32(95))
}

func TestUnloadFont(t *testing.T) {
	r := createTestEngine().Resources

	r.LoadFont(fontpath)

	err := r.UnloadFont(fontpath)

	_, ok := r.fonts["arial"]
	assert.False(t, ok)

	assert.Equal(t, len(r.fonts), 0)

	r.LoadFont(fontpath)

	r.UnloadFont("arial")

	_, ok = r.fonts["arial"]
	assert.False(t, ok)

	assert.Equal(t, len(r.fonts), 0)

	err = r.UnloadFont("badpath")
	assert.NotNil(t, err)
}

func TestCleanUp(t *testing.T) {
	e := createTestEngine()

	e.Resources.LoadDir(dirPath)

	e.Run(&testDefaultState{})

	e.QuitApp()

	assert.Equal(t, len(e.Resources.images), 0)
	assert.Equal(t, len(e.Resources.textures), 0)
	assert.Equal(t, len(e.Resources.fonts), 0)
}

func TestGetImg(t *testing.T) {
	r := createTestEngine().Resources

	r.LoadImg(applePath)

	img, exists := r.GetImg("apple")

	assert.NotNil(t, img)
	assert.True(t, exists)

	img, exists = r.GetImg("nonexistent")
	assert.NotNil(t, img) // default image
	assert.False(t, exists)
}

func TestGetTexture(t *testing.T) {
	r := createTestEngine().Resources

	r.LoadTexture(applePath)

	texture, exists := r.GetTexture("apple")

	assert.NotNil(t, texture)
	assert.True(t, exists)

	texture, exists = r.GetTexture("nonexistent")
	assert.NotNil(t, texture) // default texture
	assert.False(t, exists)
}

func TestGetFont(t *testing.T) {
	r := createTestEngine().Resources

	r.LoadFont(fontpath)

	font, exists := r.GetFont("arial")

	assert.NotNil(t, font)
	assert.True(t, exists)

	font, exists = r.GetFont("nonexistent")
	assert.Equal(t, font, rl.GetFontDefault())
	assert.False(t, exists)
}
