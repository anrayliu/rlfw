# RSky
Mini Go-RayLib framework that manages some inconvenient things for you. Based on a similar project I made a few years ago for Pygame.

I plan on using this for my own projects, so I want to update it with functionality as I discover use cases.

## Quickstart

All the boilerplate of an application is handled by the `Engine`. 
It can be configured with the `Config` struct found in the `main` function.

```
type Config struct {
	WinW     int32
	WinH     int32
	Name     string
	Fps      int32
	LogLevel rl.TraceLogLevel
}
```
All application logic can be placed in `State` structs, which have the following structure:
```
type Game struct{}

// init logic

func (g *Game) Start(e *core.Engine) {
}

// clean up logic

func (g *Game) End(e *core.Engine) {
}

// main app logic, called every frame

func (g *Game) Update(e *core.Engine) {
}

// render app, called after Update

func (g *Game) Draw(e *core.Engine) {
}
```
Start new states with `e.Run(&MyState{})`. This places a new state on the stack and gives control to it.
Exit from the current state with `e.Quit()`. This returns to the previous state in the stack.
