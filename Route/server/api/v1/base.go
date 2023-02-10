package v1

type Base struct {
	ServicePublic
	ServicePrivate
}

// 处理函数入口
var BaseService = new(Base)
