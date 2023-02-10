package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ServicePrivate struct{}

// /v1/private返回方法
func (s *ServicePrivate) GetMsg(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "private service"})
}
