package finance

import (
	"github.com/gogf/gf/frame/gmvc"
	"github.com/gogf/gf/net/ghttp"
)

// Controller 控制器结构体。
// 该结构体用于通过控制器方式注册，未来不再推荐使用控制器路由注册方式。
type Controller struct {
	r *ghttp.Request
	gmvc.Controller
	ws *ghttp.WebSocket
}

func (c *Controller) Index() {
	/*if c.Session.Contains("chat_name") {
		c.View.Assign("tplMain", "chat/include/chat.html")
	} else {
		c.View.Assign("tplMain", "chat/include/main.html")
	}*/
	c.View.Display("index.html")
}

func (c *Controller) List() {
	all, _ := FindAll()
	_ = c.Response.WriteJson(all)
}

func (c *Controller) AccountEntry() {
	c.View.Display("finance.html")
}

func (c *Controller) ToEdit() {
	c.View.Display("finance_edit.html")
}