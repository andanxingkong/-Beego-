package Vehicle

import (
	"demo/models"
	"demo/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UpdateVehicleListController struct {
	beego.Controller
}

// 表示用户尝试登录时期望的请求数据格式
type UpdateVehicleRequest struct {
	Status string `json:"Status"`
}

var VehicleId int

func (c *UpdateVehicleListController) Getupdate() {
	c.TplName = "Vehicle/updateVehicle.html"
	id, err := c.GetInt(":id") // 获取到 url 当中 id 变量的值
	VehicleId = id
	if err != nil { // 有错误就返回数据：获取参数失败
		errCode := 512 //获取参数失败
		errorResponse := utils.NewError(errCode)
		c.Data["json"] = errorResponse
		c.ServeJSON()
		return
	}
}
func (c *UpdateVehicleListController) UpdateVehicle() {
	var updateVehicleRequest UpdateVehicleRequest
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &updateVehicleRequest)
	if err != nil {
		errCode := 504 //解析请求失败
		errorResponse := utils.NewError(errCode)
		c.Data["json"] = errorResponse
		c.ServeJSON()
		return
	}

	o := orm.NewOrm()
	var r orm.RawSeter
	r = o.Raw("UPDATE vehicle SET  status = ? WHERE v_id = ?", updateVehicleRequest.Status, VehicleId)
	_, error := r.Exec()
	fmt.Println(updateVehicleRequest.Status, VehicleId)
	if error != nil {
		errCode := 505 //更新失败
		errorResponse := utils.NewError(errCode)
		c.Data["json"] = errorResponse
	} else {
		c.Data["json"] = models.Response{Success: true, Message: "更新成功"}
	}
	c.ServeJSON()
}
