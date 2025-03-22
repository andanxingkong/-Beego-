package controllers

import "github.com/astaxie/beego"

//个人中心主页

type CenterController struct {
	beego.Controller
}

func (c *CenterController) Get() {
	if !IsUserLoggedIn(&c.Controller) {
		// 用户未登录、返回错误信息或者跳转到登录页面
		c.Redirect("/login", 302)
		return
	}
	c.TplName = "center.html"
}
