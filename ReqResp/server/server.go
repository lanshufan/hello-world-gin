package server

import (
	"ReqRep/model/request"
	"ReqRep/model/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

type dataMsg struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func Server(c *gin.Context) {
	var data = dataMsg{
		Id:   1,
		Name: "haha",
	}

	str := c.PostForm("keyword")

	switch str {
	case "err":
		response.ErrResp(c)
	case "errmsg":
		response.ErrWithMsg("错误", c)
	case "success":
		response.OKResp(c)
	case "successmsg":
		response.OKWithMsgResp("正确", c)
	case "successdata":
		response.OKWithDataResp(data, c)
	case "successall":
		response.OKWithAllResp(data, "正确", c)
	case "successnil":
		response.OKWithDataResp(map[string]string{}, c)
	case "successnothing":
		response.OKNothing(c)
	default:
		fmt.Println("错误")
	}
}

func ServerJsonRequest(c *gin.Context) {
	var ids request.IdInfo
	// 解析json请求
	if err := c.ShouldBind(&ids); err != nil {
		response.ErrResp(c)
		return
	}
	response.OKWithDataResp(map[string]int{"id": ids.ID}, c)
}
