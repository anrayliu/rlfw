# rlfw
Mini raylib-go framework that manages some inconvenient things for you. No extra dependencies. This framework is designed to assist raylib, not replace it. I avoided abstractions for the sake of abstraction, and only added features if it actually made the dev experience better.

I plan on using this for my own projects, so I will update it with functionality as I discover new use cases.

## Quickstart
See `example/main.go` for a super quick starter template.

All the boilerplate of an application is handled by the `Engine`. 
It can be configured with a `Config` struct found in the `main` function.

```
engine, err := rlfw.NewEngine(rlfw.Config{
	WinW:     800,
	WinH:     600,
	WinMode:  rl.FlagFullscreenMode,
	Name:     "example",
	Fps:      60,
	LogLevel: rl.LogDebug,
	LoadAssets: true,
})
if err != nil {
	panic(err)
}
```
`rlfw` uses a stack-based state manager. States have the following interface:
```
type Game struct{}

func (g *Game) Enter(e *rlfw.Engine) {
	// init logic
}

func (g *Game) Exit(e *rlfw.Engine) {
	// clean up logic
}

func (g *Game) Update(e *rlfw.Engine) {
	// main app logic, called every frame
}

func (g *Game) Draw(e *rlfw.Engine) {
	// render app, called after Update
}

func (g *Game) Resize(e *rlfw.Engine) {
	// called when window is resized
}
```
Start new states with `e.Run(&MyState{})`. Exit from the current state with `e.QuitState()`. 
This returns control to the previous state in the stack. To exit from all states, use `e.QuitApp()`.

On app start, images (`.png`, `.jpg`) and fonts (`.otf`, `.ttf`) from the `assets` directory will be automatically loaded
into `engine.Resources`. Since each state function has access to the engine via the parameter `e`, you can easily
access the loaded resources with `e.Resources.GetImg("apple")`. Note: for convenience, you don't need the extension when
accessing loaded resources. For example, load a texture with `e.Resources.LoadTexture("pictures/dog.png")`, but
access it with `e.Resources.GetTexture("dog")`.

Full list of resource functions:

 - `LoadImg(path) error` - Load image file as a raylib image
 - `LoadTexture(path) error` - Load image file as a raylib texture
 - `LoadFont(path) error` - Load font file
 - `LoadDir(path) error` - Automatically load images and fonts in directory
 - `UnloadImg(nameOrPath) error` - Frees resource stored in memory
 - `UnloadTexture(nameOrPath) error` - Frees resource stored in memory
 - `UnloadFont(nameOrPath) error` - Frees resource stored in memory
 - `UnloadDir(path) error` - Unloads any resources that were loaded from this directory
 - `GetImg(name) *rl.Image` - Return stored image
 - `GetTexture(name) rl.Texture2D` - Return stored texture
 - `GetFont(name) rl.Font` - Return stored font

Loaded resources are automatically cleaned up when there are no more states.

## Design Philosophy

1. Errors should be genuine errors, situations where the requested function could not complete properly.
If the result is unexpected, an error should be returned.

2. Functions should be idempotent.

Keeping these 2 points in mind, when a function loads an image, if the path does not exist, devs should receive an error.
This is unexpected behaviour, and not returning an error can be confusing. However, if the image is already loaded, subsequent
calls to load won't return an error because the function is idempotent. There is no unexpected behaviour here, so an error is not needed.

3. Functions should always try to keep expected behaviour.

Take the case when a user is trying to get a texture. Even if that texture does not exist, devs are still expecting a texture 
to be returned, so a default texture is returned.
