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

	err := r.LoadDir(dirPath)
	assert.Nil(t, err)

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

	err := r.LoadDir(dirPath)
	assert.Nil(t, err)

	err = r.UnloadDir(dirPath)
	assert.Nil(t, err)

	assert.Equal(t, len(r.images), 0)
	assert.Equal(t, len(r.textures), 0)
	assert.Equal(t, len(r.fonts), 0)
}

func TestSplitFileName(t *testing.T) {
	base, ext, err := splitFileNameIfExists(applePath)
	assert.Nil(t, err)
	assert.Equal(t, base, "apple")
	assert.Equal(t, ext, ".png")

	_, _, err = splitFileNameIfExists("badpath")
	assert.NotNil(t, err)
}

func TestLoadImg(t *testing.T) {
	r := createTestEngine().Resources

	err := r.LoadImg(applePath)
	assert.Nil(t, err)

	assert.Equal(t, len(r.images), 1)

	image, ok := r.images["apple"]
	assert.True(t, ok)

	assert.Equal(t, image.Width, int32(256))
	assert.Equal(t, image.Height, int32(256))

	err = r.LoadImg(fontpath)
	assert.NotNil(t, err)
}

func TestUnloadImg(t *testing.T) {
	r := createTestEngine().Resources

	err := r.LoadImg(applePath)
	assert.Nil(t, err)

	err = r.UnloadImg(applePath)
	assert.Nil(t, err)

	_, ok := r.images["apple"]
	assert.False(t, ok)

	assert.Equal(t, len(r.images), 0)

	err = r.LoadImg(applePath)
	assert.Nil(t, err)

	err = r.UnloadImg("apple")
	assert.Nil(t, err)

	_, ok = r.images["apple"]
	assert.False(t, ok)

	assert.Equal(t, len(r.images), 0)

	err = r.UnloadImg("badpath")
	assert.NotNil(t, err)
}

func TestLoadTexture(t *testing.T) {
	r := createTestEngine().Resources

	err := r.LoadTexture(applePath)
	assert.Nil(t, err)

	assert.Equal(t, len(r.textures), 1)

	texture, ok := r.textures["apple"]
	assert.True(t, ok)

	assert.Equal(t, texture.Width, int32(256))
	assert.Equal(t, texture.Height, int32(256))

	err = r.LoadTexture(fontpath)
	assert.NotNil(t, err)
}

func TestUnloadTexture(t *testing.T) {
	r := createTestEngine().Resources

	err := r.LoadTexture(applePath)
	assert.Nil(t, err)

	err = r.UnloadTexture(applePath)
	assert.Nil(t, err)

	_, ok := r.textures["apple"]
	assert.False(t, ok)

	assert.Equal(t, len(r.textures), 0)

	err = r.LoadTexture(applePath)
	assert.Nil(t, err)

	err = r.UnloadTexture("apple")
	assert.Nil(t, err)

	_, ok = r.textures["apple"]
	assert.False(t, ok)

	assert.Equal(t, len(r.textures), 0)

	err = r.UnloadTexture("badpath")
	assert.NotNil(t, err)
}

func TestLoadFont(t *testing.T) {
	r := createTestEngine().Resources

	err := r.LoadFont(fontpath)
	assert.Nil(t, err)

	assert.Equal(t, len(r.fonts), 1)

	font, ok := r.fonts["arial"]
	assert.True(t, ok)

	assert.Equal(t, font.CharsCount, int32(95))
}

func TestUnloadFont(t *testing.T) {
	r := createTestEngine().Resources

	err := r.LoadFont(fontpath)
	assert.Nil(t, err)

	err = r.UnloadFont(fontpath)
	assert.Nil(t, err)

	_, ok := r.fonts["arial"]
	assert.False(t, ok)

	assert.Equal(t, len(r.fonts), 0)

	err = r.LoadFont(fontpath)
	assert.Nil(t, err)

	err = r.UnloadFont("arial")
	assert.Nil(t, err)

	_, ok = r.fonts["arial"]
	assert.False(t, ok)

	assert.Equal(t, len(r.fonts), 0)

	err = r.UnloadFont("badpath")
	assert.NotNil(t, err)
}

func TestCleanUp(t *testing.T) {
	e := createTestEngine()

	err := e.Resources.LoadDir(dirPath)
	assert.Nil(t, err)

	e.Run(&testDefaultState{})

	e.QuitApp()

	assert.Equal(t, len(e.Resources.images), 0)
	assert.Equal(t, len(e.Resources.textures), 0)
	assert.Equal(t, len(e.Resources.fonts), 0)
}

func TestGetImg(t *testing.T) {
	r := createTestEngine().Resources

	err := r.LoadImg(applePath)
	assert.Nil(t, err)

	img := r.GetImg("apple")
	assert.NotNil(t, img)

	img = r.GetImg("nonexistent")
	assert.NotNil(t, img) // default image
}

func TestGetTexture(t *testing.T) {
	r := createTestEngine().Resources

	err := r.LoadTexture(applePath)
	assert.Nil(t, err)

	texture := r.GetTexture("apple")
	assert.NotNil(t, texture)

	texture = r.GetTexture("nonexistent")
	assert.NotNil(t, texture) // default texture
}

func TestGetFont(t *testing.T) {
	r := createTestEngine().Resources

	err := r.LoadFont(fontpath)
	assert.Nil(t, err)

	font := r.GetFont("arial")
	assert.NotNil(t, font)

	font = r.GetFont("nonexistent")
	assert.Equal(t, font, rl.GetFontDefault())
}
