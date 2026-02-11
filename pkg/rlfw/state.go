package rlfw

type State interface {
	Enter(e *Engine)
	Exit(e *Engine)
	Update(e *Engine)
	Draw(e *Engine)
	Resize(e *Engine)
}

type DefaultState struct{}

func (s *DefaultState) Enter(e *Engine) {
}

func (s *DefaultState) Exit(e *Engine) {
}

func (s *DefaultState) Update(e *Engine) {
}

func (s *DefaultState) Draw(e *Engine) {
}

func (s *DefaultState) Resize(e *Engine) {
}
