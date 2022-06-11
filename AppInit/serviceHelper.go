package appinit

import (
	"context"

	"github.com/gin-gonic/gin"
)

//业务最终函数原型
type Endpoint func(ctx *gin.Context, request interface{}) (response interface{}, err error)

//怎么取参数
type EncodeRequestFunc func(*gin.Context) (interface{}, error)

//怎么处理业务结果
type DecodeResponseFunc func(*gin.Context, interface{}) error

func RegisterHandler(endpoint Endpoint, encodeFunc EncodeRequestFunc, decodeFunc DecodeResponseFunc) func(ctx *gin.Context) {
	
	return func(ctx *gin.Context) {

	}
}
