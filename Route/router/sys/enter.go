package sys

type routeInit struct {
	PublicRoute
	PrivateRoute
}

// 路由入口
var RouteInit = new(routeInit)
