package finance

import (
	"fumanju/library/response"
	idworker "github.com/gitstliu/go-id-worker"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
)

// Controller 控制器结构体。
// 该结构体用于通过控制器方式注册，未来不再推荐使用控制器路由注册方式。
type Controller struct {
	r *ghttp.Request
	gmvc.Controller
	ws *ghttp.WebSocket
}

type RegisterReq struct {
	Id             int64       `p:"id"`
	FinanceName    string      `p:"financeName"`    // 财务单名称
	FinanceProject string      `p:"financeProject"` // 财务单项目
	FinanceClass   string      `p:"financeClass"`   // 财务单类别
	FinancePrice   float64     `p:"financePrice"`   // 财务单金额
	Createby       string      `p:"createby"`       // 负责人
	Create         *gtime.Time `p:"create"`         // 创建时间
	UpdateTime     *gtime.Time `p:"updateTime"`     // 修改时间
	Remark         string      `p:"remark"`         // 备注

	Pass  string `p:"password1"`
	Pass2 string `p:"password2"`
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
	entity := &[]Entity{}
	all, _ := g.DB().Table("finance").Limit(c.Request.GetInt("page"), c.Request.GetInt("limit")).All()
	_ = all.Structs(entity)
	var a = make(map[string]interface{})
	a["list"] = entity
	a["count"] = len(all)
	_ = c.Response.WriteJson(a)
}

func (c *Controller) AccountEntry() {
	c.View.Display("finance.html")
}

func (c *Controller) ToEdit() {
	c.View.Display("finance_edit.html")
}

func (c *Controller) DoEdit() {
	var entity *Entity
	c.Request.Parse(&entity)

	if entity.Id == 0 {
		currWoker := &idworker.IdWorker{}
		_ = currWoker.InitIdWorker(100, 1)
		newId, _ := currWoker.NextId()
		entity.Id = newId
		Save(entity)
	} else {
		Update(entity)
	}
	//c.Response.WriteJson("{'msg': '新增、修改成功！' 'code': '0'}")
	response.JsonExit(c.Request, 0, "新增、修改成功！")
}
