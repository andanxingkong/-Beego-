package controllers

import (
	"demo/models"
	"demo/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//个人中心的添加车辆

type CenterVehicleController struct {
	beego.Controller
}

func (c *CenterVehicleController) Get() {
	if !IsUserLoggedIn(&c.Controller) {
		// 用户未登录、返回错误信息或者跳转到登录页面
		c.Redirect("/login", 302)
		return
	}
	c.TplName = "Vehicle/addcenterVehicle.html"
}
func (c *CenterVehicleController) AddCenterVehicle() {
	var vehicle models.Vehicle
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &vehicle)
	if err != nil {
		errCode := 504
		// 使用封装的NewError函数来构建错误信息
		errorResponse := utils.NewError(errCode)
		c.Data["json"] = errorResponse
		c.ServeJSON()
		return
	}
	om := orm.NewOrm()
	id, err := om.Insert(&vehicle)
	fmt.Println("V_id:", vehicle.V_id, "Owner_name:", vehicle.Owner_name, "Plate_number:", vehicle.Plate_number, "Status:", vehicle.Status)
	if id > 0 {
		c.Data["json"] = models.Response{Success: true, Message: "添加成功"}
	} else {
		errCode := 509 //添加失败
		errorResponse := utils.NewError(errCode)
		c.Data["json"] = errorResponse
	}
	c.ServeJSON()

}
