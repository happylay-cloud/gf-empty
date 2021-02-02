package router

import (
	"github.com/gogf/gf-empty/app/api/hello"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/swagger"

	"github.com/happylay-cloud/gf-extend/compoent/hweb"
)

func init() {
	s := g.Server()
	s.Plugin(&swagger.Swagger{})
	s.Use(hweb.MiddlewareErrorHandler)
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(hweb.MiddlewareCORS)
		group.ALL("/", hello.Hello)
		group.POST("/version/:id", hello.GfVersion)
		group.POST("/page", hello.Page)
	})
}
