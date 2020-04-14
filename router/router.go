package router

import (
	"fumanju/app/finance"
	"fumanju/app/middleware"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		financeCtrl := new(finance.Controller)
		group.Middleware(middleware.CORS)
		group.ALL("/", financeCtrl)
	})
}
