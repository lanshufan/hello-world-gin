package router

import (
	"Route/router/sys"
	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	r := gin.Default()

	// 路由组v1
	publicRouterGroup := sys.RouteInit.PublicRoute
	privateRouterGroup := sys.RouteInit.PrivateRoute
	routerGroup := r.Group("v1")
	{
		publicRouterGroup.PublicRouteGroup(routerGroup)
		privateRouterGroup.PrivateRouteGroup(routerGroup)
	}

	return r
}
