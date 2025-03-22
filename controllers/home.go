package controllers

import (
	"demo/utils"
	"github.com/astaxie/beego"
)

//用户主页

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	redisUtil := utils.NewRedisCache()
	username, _ := redisUtil.Get(LOGIN)
	if username != nil {
		// 用户名存在，返回给前端页面
		c.Data["Username"] = username
	} else {
		// 用户名不存在，可以返回错误信息或者处理匿名用户的情况
		c.Data["Username"] = "Anonymous"
	}
	c.TplName = "home.html"
}
