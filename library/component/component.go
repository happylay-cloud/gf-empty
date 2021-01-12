package component

import (
	"github.com/gogf/gf-empty/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// 跨域中间件
func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

// 全局后置中间件捕获异常
func MiddlewareErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()
	if err := r.GetError(); err != nil {
		// 记录到自定义错误日志文件
		g.Log("exception").Error(err)
		// 清除系统异常响应
		r.Response.ClearBuffer()
		// 返回自定义异常响应
		response.FailWithEx(r, err.Error())
	}
}
