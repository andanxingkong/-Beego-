package controllers

import (
	"demo/models"
	"demo/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//注册

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Get() {
	c.TplName = "register.html"
}

func (c *RegisterController) Register() {
	var user models.Users
	var user2 models.Users
	o := orm.NewOrm()
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	if err != nil {
		errCode := 504
		// 使用封装的NewError函数来构建错误信息
		errorResponse := utils.NewError(errCode)
		c.Data["json"] = errorResponse
		c.ServeJSON()
		return
	}

	error := o.QueryTable("Users").Filter("Phone", user.Phone).One(&user2)
	if error == nil {
		errCode := 507
		// 使用封装的NewError函数来构建错误信息
		errorResponse := utils.NewError(errCode)
		c.Data["json"] = errorResponse
	} else {
		user.Password = utils.MD5(user.Password)
		fmt.Println("md5后：", user.Password)
		id, _ := o.Insert(&user)
		fmt.Println("Id:", user.Id, "Username:", user.Username, "password:", user.Password, "Phone:", user.Phone, "Email:", user.Email, "Role_name:", user.Role_name)
		if id > 0 {
			c.Data["json"] = models.Response{Success: true, Message: "注册成功"}
		} else {
			errCode := 506
			errorResponse := utils.NewError(errCode)
			c.Data["json"] = errorResponse
		}
	}
	c.ServeJSON()
}
