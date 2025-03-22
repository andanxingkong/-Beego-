package controllers

import (
	"demo/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//巡查员查询全部订单

type PatrolPersonQueryController struct {
	beego.Controller
}

func (c *PatrolPersonQueryController) Get() {
	if !IsUserLoggedIn(&c.Controller) {
		// 用户未登录、返回错误信息或者跳转到登录页面
		c.Redirect("/login", 302)
		return
	}
	c.TplName = "PatrolPerson/Query.html"
}

func (c *PatrolPersonQueryController) GetAll() {
	o := orm.NewOrm()
	var Orders []models.OrderTable
	o.QueryTable("order_table").All(&Orders)
	c.Data["json"] = Orders
	c.ServeJSON()
}
