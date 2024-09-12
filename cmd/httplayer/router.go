package httplayer

import (
	"github.com/dbrinkk/TodoApp/applayer"
	"github.com/gin-gonic/gin"
)

type httpApi struct {
	engine *gin.Engine
	app    applayer.App
}

func New(appLayer applayer.App) *httpApi {
	return &httpApi{}
}

func (self *httpApi) Engage() {
	self.engine.Run()
}
