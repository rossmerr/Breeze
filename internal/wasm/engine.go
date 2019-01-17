package wasm

import breeze "github.com/RossMerr/Breeze"

func init() {
	breeze.RegisterEngine("wasm", breeze.EngineRegistration{
		NewFunc: newEngine,
	})
}

type Engine struct {
}

func newEngine() (breeze.Engine, error) {
	return &Engine{}, nil

}
func (s Engine) Run(widget breeze.Widget) {

}
func (s Engine) Rerender(widget breeze.Widget) {

}
