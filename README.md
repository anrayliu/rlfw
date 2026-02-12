# rlfw v0.1.0
Mini Go-RayLib framework that manages some inconvenient things for you. No dependencies other than raylib. This framework is designed to assist raylib, not replace it. I tried to avoid abstraction for the sake of abstraction, and only added features if it actually made the dev experience better.

I plan on using this for my own projects, so I will update it with functionality as I discover new use cases.

## Quickstart

All the boilerplate of an application is handled by the `Engine`. 
It can be configured with the `Config` struct found in the `main` function.

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

// init logic

func (g *Game) Enter(e *rlfw.Engine) {
}

// clean up logic

func (g *Game) Exit(e *rlfw.Engine) {
}

// main app logic, called every frame

func (g *Game) Update(e *rlfw.Engine) {
}

// render app, called after Update

func (g *Game) Draw(e *rlfw.Engine) {
}

// called when window is resized

func (g *Game) Resize(e *rlfw.Engine) {
}
```
Start new states with `e.Run(&MyState{})`. Exit from the current state with `e.QuitState()`. 
This returns to the previous state in the stack. To exit from all states, use `e.QuitApp()`.

On app start, images (`.png`, `.jpg`) and fonts (`.otf`, `.ttf`) from the `assets` directory will be automatically loaded.
Inside a state, they can be accessed with `e.Resources.GetImg(name)`, `e.Resources.GetTexture(name)`, and `e.Resources.GetFont(name)`.
Further resource loading can be done with `LoadDir(path)`, `LoadImg(path)`, `LoadTexture(path)`, and `LoadFont(path)`. Use their respective
`Unload` functions to free them. Resources are automatically unloaded when all states end.

## Design Philosophy

1. Errors should be genuine errors, situations where the requested function could not complete properly.
If the result is unexpected, an error should be returned.

2. Functions should be idempotent.

Keeping these 2 points in mind, when a function loads an image, if the path does not exist, devs should receive an error.
This is unexpected behaviour, and not returning an error can be confusing. However, if the image is already loaded, subsequent
calls to load won't return an error because the function is idempotent. There is no unexpected behaviour here, so an error is not needed.

3. Functions should always try to keep expected behaviour.

Take the case when a user is trying to get a texture. If that texture does not exist, devs should be notified through
an error (in this case, a boolean instead). However, they are still expecting a texture to be returned, so a default texture is returned.
