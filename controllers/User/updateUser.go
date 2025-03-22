package User

import (
	"demo/models"
	"demo/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UpdateController struct {
	beego.Controller
}

var Id int

func (c *UpdateController) Getupdate() {
	c.TplName = "User/updateUser.html"
	id, err := c.GetInt(":id") // 获取到 url 当中 id 变量的值
	Id = id
	if err != nil { // 有错误就返回数据：获取参数失败
		errCode := 512 //获取参数失败
		errorResponse := utils.NewError(errCode)
		c.Data["json"] = errorResponse
		c.ServeJSON()
		return
	}
}
func (c *UpdateController) UpdateUser() {
	var upRequest models.UpRequest
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &upRequest)
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
	r = o.Raw("UPDATE Users SET  username = ?,password=?,Role_name=? WHERE id = ?", upRequest.Username, utils.MD5(upRequest.Password), upRequest.Role_name, Id)
	_, error := r.Exec()
	fmt.Println(upRequest.Username, utils.MD5(upRequest.Password), upRequest.Role_name, Id)
	if error != nil {
		errCode := 505 //更新失败
		errorResponse := utils.NewError(errCode)
		c.Data["json"] = errorResponse
	} else {
		c.Data["json"] = models.Response{Success: true, Message: "更新成功"}
	}
	c.ServeJSON()
}
