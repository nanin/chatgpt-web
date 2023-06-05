package bootstrap

import (
	"chatgpt-web/routes"
	"sync"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine
var once sync.Once

func SetUpRoute() {
	once.Do(func() {
		router = gin.Default()
		routes.RegisterWebRoutes(router)
	})
}
