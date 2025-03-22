package ParkingSpace

import (
	"demo/models"
	"demo/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type AddParkingSpaceController struct {
	beego.Controller
}

func (c *AddParkingSpaceController) Get() {
	c.TplName = "ParkingSpace/addParkingSpace.html"
}

func (c *AddParkingSpaceController) AddParkingSpace() {
	var parkingSpace models.ParkingSpace
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &parkingSpace)
	if err != nil {
		errCode := 504
		// 使用封装的NewError函数来构建错误信息
		errorResponse := utils.NewError(errCode)
		c.Data["json"] = errorResponse
		c.ServeJSON()
		return
	}

	om := orm.NewOrm()
	id, err := om.Insert(&parkingSpace)
	fmt.Println("Id:", parkingSpace.Id, "Road_id:", parkingSpace.Road_id, "Space_number:", parkingSpace.Space_number, "Status:", parkingSpace.Status)
	if id > 0 {
		c.Data["json"] = models.Response{Success: true, Message: "添加成功"}
	} else {
		errCode := 509 //添加失败
		errorResponse := utils.NewError(errCode)
		c.Data["json"] = errorResponse
	}
	c.ServeJSON()

}
