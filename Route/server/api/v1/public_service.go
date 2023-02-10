package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ServicePublic struct{}

// /v1/public返回方法
func (s *ServicePublic) GetMsg(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "public service"})
}
