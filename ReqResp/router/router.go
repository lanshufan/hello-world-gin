package router

import (
	"ReqRep/server"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.POST("/", server.Server)
	router.POST("/json", server.ServerJsonRequest)
	return router
}
