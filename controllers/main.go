package controllers

import "github.com/astaxie/beego"

//管理员主页

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	if !IsUserLoggedIn(&c.Controller) {
		// 用户未登录、返回错误信息或者跳转到登录页面
		c.Redirect("/login", 302)
		return
	}
	c.TplName = "managerpage.html"
}
