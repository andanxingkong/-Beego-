package controllers

import (
	"github.com/astaxie/beego"
)

type ExitController struct {
	beego.Controller
}

func (c *ExitController) Get() {
	//清除该用户登录状态的数据
	c.DelSession(NAME)
	c.Redirect("/", 302)
}
