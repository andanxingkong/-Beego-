package RoadSegment

import (
	"demo/models"
	"demo/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type AddRoadSegmentController struct {
	beego.Controller
}

func (c *AddRoadSegmentController) Get() {
	c.TplName = "RoadSegment/addRoadSegment.html"
}

func (c *AddRoadSegmentController) AddRoadSegment() {
	var roadSegment models.RoadSegment
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &roadSegment)
	if err != nil {
		errCode := 504
		// 使用封装的NewError函数来构建错误信息
		errorResponse := utils.NewError(errCode)
		c.Data["json"] = errorResponse
		c.ServeJSON()
		return
	}

	om := orm.NewOrm()
	id, err := om.Insert(&roadSegment)
	fmt.Println("Road_id:", roadSegment.Road_id, "Code:", roadSegment.Code, "Name:", roadSegment.Name, "Status:", roadSegment.Status)
	if id > 0 {
		c.Data["json"] = models.Response{Success: true, Message: "添加成功"}
	} else {
		errCode := 509 //添加失败
		errorResponse := utils.NewError(errCode)
		c.Data["json"] = errorResponse
	}
	c.ServeJSON()

}
