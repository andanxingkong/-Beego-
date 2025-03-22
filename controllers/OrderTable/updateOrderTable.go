package OrderTable

import (
	"demo/models"
	"demo/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type UpdateOrderTableListController struct {
	beego.Controller
}

var End_time string

// 表示用户尝试登录时期望的请求数据格式
type UpdateOrderTableRequest struct {
	Status string `json:"Status"`
}

var OrderTableId int

func (c *UpdateOrderTableListController) Getupdate() {
	c.TplName = "OrderTable/updateOrderTable.html"
	id, err := c.GetInt(":id") // 获取到 url 当中 id 变量的值
	OrderTableId = id
	if err != nil { // 有错误就返回数据：获取参数失败
		errCode := 512 //获取参数失败
		errorResponse := utils.NewError(errCode)
		c.Data["json"] = errorResponse
		c.ServeJSON()
		return
	}
}
func (c *UpdateOrderTableListController) UpdateOrderTable() {
	var updateOrderTableRequest UpdateOrderTableRequest
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &updateOrderTableRequest)
	if err != nil {
		errCode := 504
		// 使用封装的NewError函数来构建错误信息
		errorResponse := utils.NewError(errCode)
		c.Data["json"] = errorResponse
		c.ServeJSON()
		return
	}
	var orderTable models.OrderTable
	o := orm.NewOrm()
	o.QueryTable("order_table").Filter("R_id", OrderTableId).One(&orderTable)
	var error error
	if updateOrderTableRequest.Status == "已完成" {
		End_time = time.Now().Format("2006-01-02 15:5:2")
		o := orm.NewOrm()
		var r orm.RawSeter
		timeInterval := utils.CalculationTime(orderTable.Start_time, End_time)
		money := utils.CalculationMoney(timeInterval)
		r = o.Raw("UPDATE order_table SET  End_time = ?,status = ?,cost = ? WHERE R_id = ?", End_time, updateOrderTableRequest.Status, money, OrderTableId)
		_, error = r.Exec()
		fmt.Println(End_time, updateOrderTableRequest.Status, OrderTableId)
	}
	if updateOrderTableRequest.Status == "未完成" {
		o := orm.NewOrm()
		var r orm.RawSeter
		r = o.Raw("UPDATE order_table SET  End_time = ?,status = ? WHERE R_id = ?", "", updateOrderTableRequest.Status, OrderTableId)
		_, error = r.Exec()
		fmt.Println(updateOrderTableRequest.Status, OrderTableId)
	}
	if error != nil {
		errCode := 505 //更新失败
		errorResponse := utils.NewError(errCode)
		c.Data["json"] = errorResponse
	} else {
		c.Data["json"] = models.Response{Success: true, Message: "更新成功"}
	}
	c.ServeJSON()
}
