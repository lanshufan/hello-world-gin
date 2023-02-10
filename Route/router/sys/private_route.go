package sys

import (
	v1 "Route/server/api/v1"
	"github.com/gin-gonic/gin"
)

type PrivateRoute struct{}

// private POST请求方法
func (p *PrivateRoute) PrivateRouteGroup(group *gin.RouterGroup) {
	privateServer := v1.BaseService.ServicePrivate
	group.POST("private", privateServer.GetMsg)
}
