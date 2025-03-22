package User

import (
	"demo/models"
	"demo/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type AdduserController struct {
	beego.Controller
}

func (c *AdduserController) Get() {
	c.TplName = "User/adduser.html"
}

func (c *AdduserController) Adduser() {
	var user models.Users
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	if err != nil {
		errCode := 504
		// 使用封装的NewError函数来构建错误信息
		errorResponse := utils.NewError(errCode)
		c.Data["json"] = errorResponse
		c.ServeJSON()
		return
	}
	user.Password = utils.MD5(user.Password)
	fmt.Println("md5后：", user.Password)
	om := orm.NewOrm()
	id, err := om.Insert(&user)
	fmt.Println("Id:", user.Id, "Username:", user.Username, "password:", user.Password, "Phone:", user.Phone, "Email:", user.Email, "Role_name:", user.Role_name)
	if id > 0 {
		c.Data["json"] = models.Response{Success: true, Message: "添加成功"}
	} else {
		errCode := 509 //添加失败
		errorResponse := utils.NewError(errCode)
		c.Data["json"] = errorResponse
	}
	c.ServeJSON()

}
