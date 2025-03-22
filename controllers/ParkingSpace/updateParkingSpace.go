package ParkingSpace

import (
	"demo/models"
	"demo/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UpdateParkingSpaceController struct {
	beego.Controller
}

// 表示用户尝试登录时期望的请求数据格式
type UpdateParkingSpaceRequest struct {
	Status string `json:"role"`
}

var ParkingSpaceId int

func (c *UpdateParkingSpaceController) Getupdate() {
	c.TplName = "ParkingSpace/UpDateParkingSpace.html"
	id, err := c.GetInt(":id") // 获取到 url 当中 id 变量的值
	ParkingSpaceId = id
	if err != nil { // 有错误就返回数据：获取参数失败
		errCode := 512 //获取参数失败
		errorResponse := utils.NewError(errCode)
		c.Data["json"] = errorResponse
		c.ServeJSON()
		return
	}
}
func (c *UpdateParkingSpaceController) UpdateParkingSpace() {
	var updateParkingSpaceRequest UpdateParkingSpaceRequest
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &updateParkingSpaceRequest)
	if err != nil {
		errCode := 504
		// 使用封装的NewError函数来构建错误信息
		errorResponse := utils.NewError(errCode)
		c.Data["json"] = errorResponse
		c.ServeJSON()
		return
	}
	o := orm.NewOrm()
	var r orm.RawSeter
	r = o.Raw("UPDATE parking_space SET  status = ? WHERE id = ?", updateParkingSpaceRequest.Status, ParkingSpaceId)
	_, error := r.Exec()
	fmt.Println(updateParkingSpaceRequest.Status, ParkingSpaceId)
	if error != nil {
		errCode := 505 //更新失败
		errorResponse := utils.NewError(errCode)
		c.Data["json"] = errorResponse
	} else {
		c.Data["json"] = models.Response{Success: true, Message: "更新成功"}
	}
	c.ServeJSON()
}
