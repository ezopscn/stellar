package middleware

import (
	"net/http"
	"stellar/pkg/response"

	"github.com/gin-gonic/gin"
)

// 为了避免因为处理异常导致程序 panic 退出，所以需要对 panic 进行封装，响应也通过 panic 方法处理
func Exception(ctx *gin.Context) {
	defer func() {
		err := recover()
		if err != nil {
			// 使用断言判断错误是用户定义的响应异常还是系统抛出的异常
			resp, ok := err.(response.Response)
			if !ok {
				// 系统异常就响应用户服务器错误，避免值直接 panic 导致程序退出
				resp = response.Response{
					Code:    response.RequestInternalServerError,
					Message: response.ResponseMessage[response.RequestInternalServerError],
					Data:    map[string]interface{}{},
				}
			}
			// 无论请求是否成功，响应的状态码都应该是 200，都会有返回数据
			ctx.JSON(http.StatusOK, resp)
			ctx.Abort()
			return
		}
	}()
	ctx.Next()
}
