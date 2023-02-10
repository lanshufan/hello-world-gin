package sys

import (
	v1 "Route/server/api/v1"
	"github.com/gin-gonic/gin"
)

type PublicRoute struct{}

// public POST请求方法
func (p *PublicRoute) PublicRouteGroup(group *gin.RouterGroup) {
	publicServer := v1.BaseService.ServicePublic
	group.POST("public", publicServer.GetMsg)
}
