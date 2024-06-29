package applayer

import (
	"github.com/dbrinkk/TodoApp/storelayer"
)

type App interface {
}

type app struct {
	store storelayer.Store
}

func New(store storelayer.Store) *app {
	return &app{}
}
