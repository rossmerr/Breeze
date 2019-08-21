package breeze

import (
	"errors"

	"github.com/sirupsen/logrus"
)

var ErrEngineNotRegistred = errors.New("This Engine is not registered.")

type NewEngineFunc func() (Engine, error)

type EngineRegistration struct {
	NewFunc NewEngineFunc
}

type EngineRegistry struct {
}

var engineRegistry = make(map[string]EngineRegistration)

func RegisterEngine(name string, register EngineRegistration) {
	if register.NewFunc == nil {
		logrus.Panic("NewFunc must not be nil")
	}

	if _, found := engineRegistry[name]; found {
		logrus.Panicf("Already registered Engine %q.", name)
	}
	engineRegistry[name] = register
}

func NewEngine(name string) (Engine, error) {
	r, registered := engineRegistry[name]
	if !registered {
		return nil, ErrEngineNotRegistred
	}

	return r.NewFunc()
}
