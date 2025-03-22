package Vehicle

import (
	"demo/models"
	"demo/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type AddVehicleController struct {
	beego.Controller
}

func (c *AddVehicleController) Get() {
	c.TplName = "Vehicle/addVehicle.html"
}

func (c *AddVehicleController) AddVehicle() {
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
