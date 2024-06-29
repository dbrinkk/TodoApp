package main

import (
	"github.com/dbrinkk/TodoApp/applayer"
	"github.com/dbrinkk/TodoApp/httplayer"
	"github.com/dbrinkk/TodoApp/storelayer"
)

func main() {
	// 	// create store layer
	storeLayer := storelayer.New()

	// 	// create app layer
	appLayer := applayer.New(storeLayer)

	// 	// create http layer
	api := httplayer.New(appLayer)

	api.Engage()
}
