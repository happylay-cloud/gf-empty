package hello

import (
	"github.com/gogf/gf"
	"github.com/gogf/gf-empty/library/request"
	"github.com/gogf/gf-empty/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// Hello 是一个用于输出"Hello World!"的演示路由处理程序。
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
// @Security	ApiKeyAuth
// @Success 200 {object} response.JsonResponse      "响应结果"
// @Failure 201 {object} string
// @Router 	/version/{id} [post]
func GfVersion(r *ghttp.Request) {
	var _ = 1 / r.GetInt("id")
	g.Log("business").Line(true).Info("自定义业务日志示例。")
	response.OkWithData(r, gf.VERSION)
}

// @Tags        接口标签
// @Summary     接口摘要
// @Description 接口说明
// @Accept	application/json
// @Produce application/json
// @Param   pageReq body request.PageReq true	"请求参数"
// @Success 200 {object} response.JsonResponse	"响应结果"
// @Failure 201 {object} string
// @Router	/page [post]
func Page(r *ghttp.Request) {
	g.Dump(r.GetBody())
	var req *request.PageReq
	if err := r.Parse(&req); err != nil {
		response.FailWithEx(r, err.Error())
	}
	response.OkWithData(r, req)
}
