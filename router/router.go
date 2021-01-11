package router

import (
    "github.com/gogf/gf-empty/app/api/hello"
	"github.com/gogf/gf-empty/app/component"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/swagger"
)

func init() {
	s := g.Server()
	s.Plugin(&swagger.Swagger{})
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(component.MiddlewareCORS)
		group.ALL("/", hello.Hello)
		group.POST("/version/:id", hello.GfVersion)
	})
}

