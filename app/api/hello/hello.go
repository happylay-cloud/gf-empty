package hello

import (
    "github.com/gogf/gf"
    "github.com/gogf/gf/net/ghttp"
)

// Hello是一个用于输出"Hello World!"的演示路由处理程序。
func Hello(r *ghttp.Request) {
    r.Response.Writeln("Hello World!")
}

// @Tags        接口标签
// @Summary     接口摘要
// @Description 接口说明
// @Accept      application/json
// @Produce     application/json
// @Param       id path string true                "Restful接口参数"
// @Param       field query string false           "Get请求参数"
// @Param       timestamp formData string false    "时间戳"
// @Param       Authorization header string false  "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Success 201 {object} string
// @Failure 400 {object} string
// @Router /version/{id} [post]
func GfVersion(r *ghttp.Request) {
    r.Response.Writeln(gf.VERSION)
}
