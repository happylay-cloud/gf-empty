package main

import (
	_ "github.com/gogf/gf-empty/boot"
	_ "github.com/gogf/gf-empty/router"

	"github.com/gogf/gf/frame/g"
)

// @Title GoFrame框架
// @Version v1.15.1
// @Description [`GF(Go Frame)`](https://goframe.org/)是一款模块化、高性能、企业级的Go基础开发框架。

// @Host localhost:8080
// @BasePath /
// @Schemes http https

// @SecurityDefinitions.basic BasicAuth
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
func main() {
	g.Server().Run()
}
