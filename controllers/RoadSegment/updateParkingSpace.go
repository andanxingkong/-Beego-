package RoadSegment

import (
	"demo/models"
	"demo/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UpdateRoadSegmentController struct {
	beego.Controller
}

// 表示用户尝试登录时期望的请求数据格式
type UpdateRoadSegmentRequest struct {
	Status string `json:"Status"`
}

var RoadSegmentId int

func (c *UpdateRoadSegmentController) Getupdate() {
	c.TplName = "RoadSegment/UpDateRoadSegment.html"
	id, err := c.GetInt(":id") // 获取到 url 当中 id 变量的值
	RoadSegmentId = id
	if err != nil { // 有错误就返回数据：获取参数失败
		errCode := 512 //获取参数失败
		errorResponse := utils.NewError(errCode)
		c.Data["json"] = errorResponse
		c.ServeJSON()
		return
	}
}
func (c *UpdateRoadSegmentController) UpdateRoadSegment() {
	var updateRoadSegmentRequest UpdateRoadSegmentRequest
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &updateRoadSegmentRequest)
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
	r = o.Raw("UPDATE road_segment SET  status = ? WHERE road_id = ?", updateRoadSegmentRequest.Status, RoadSegmentId)
	_, error := r.Exec()
	fmt.Println(updateRoadSegmentRequest.Status, RoadSegmentId)
	if error != nil {
		errCode := 505 //更新失败
		errorResponse := utils.NewError(errCode)
		c.Data["json"] = errorResponse
	} else {
		c.Data["json"] = models.Response{Success: true, Message: "更新成功"}
	}
	c.ServeJSON()
}
