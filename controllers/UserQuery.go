package controllers

import (
	"demo/models"
	"demo/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

//个人中心查询订单和结束订单

type GatAllController struct {
	beego.Controller
}

func (c *GatAllController) Get() {
	if !IsUserLoggedIn(&c.Controller) {
		// 用户未登录、返回错误信息或者跳转到登录页面
		c.Redirect("/login", 302)
		return
	}
	c.TplName = "Query.html"
}

func (c *GatAllController) GetAll() {
	fmt.Println("-----" + USERNAME)
	user, err := models.GetUserWithPosts(USERNAME)
	if err != nil {
		fmt.Println("查询失败:", err)
		return
	}
	var Order []models.OrderTable
	for _, post := range user.OrderTable {
		fmt.Println(post)
		Order = append(Order, *post)
	}
	c.Data["json"] = Order
	c.ServeJSON()
}

func (c *GatAllController) UpOrderTable() {
	OrderTableId, _ := c.GetInt(":id") // 获取到 url 当中 id 变量的值
	fmt.Println(OrderTableId)
	var orderTable models.OrderTable
	o := orm.NewOrm()
	o.QueryTable("order_table").Filter("R_id", OrderTableId).One(&orderTable)
	var error error
	End_time := time.Now().Format("2006-01-02 15:04:05")
	var r orm.RawSeter
	timeInterval := utils.CalculationTime(orderTable.Start_time, End_time)
	money := utils.CalculationMoney(timeInterval)
	r = o.Raw("UPDATE order_table SET  End_time = ?,status = ?,cost = ? WHERE R_id = ?", End_time, "已完成", money, OrderTableId)
	_, error = r.Exec()
	if error != nil {
		errCode := 505
		// 使用封装的NewError函数来构建错误信息
		errorResponse := utils.NewError(errCode)
		c.Data["json"] = errorResponse
	} else {
		c.Data["json"] = models.Response{Success: true, Message: "更新成功"}
	}
	c.ServeJSON()
}
