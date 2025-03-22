package controllers

import (
	"demo/models"
	"demo/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//个人中心更改密码

type UpPasswordController struct {
	beego.Controller
}

func (c *UpPasswordController) Get() {
	if !IsUserLoggedIn(&c.Controller) {
		// 用户未登录、返回错误信息或者跳转到登录页面
		c.Redirect("/login", 302)
		return
	}
	c.TplName = "upPassword.html"
}

func (c *UpPasswordController) UpPasswordRequest() {
	var uppasswordRequest models.UppasswordRequest
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &uppasswordRequest)
	if err != nil {
		errCode := 504
		// 使用封装的NewError函数来构建错误信息
		errorResponse := utils.NewError(errCode)
		c.Data["json"] = errorResponse
		c.ServeJSON()
		return
	}
	var user models.Users
	o := orm.NewOrm()
	var r orm.RawSeter
	error := o.QueryTable("Users").Filter("Username", USERNAME).One(&user)
	if error != nil {
		// 查询出错
		c.Data["json"] = map[string]interface{}{"error": "User not found"}
		c.ServeJSON()
		return
	}
	if user.Phone == uppasswordRequest.Phone {
		r = o.Raw("UPDATE Users SET password=? WHERE Phone = ?", utils.MD5(uppasswordRequest.Password), uppasswordRequest.Phone)
		_, error = r.Exec()
		fmt.Println(utils.MD5(uppasswordRequest.Password), uppasswordRequest.Phone)
		if error != nil {
			errCode := 505
			// 使用封装的NewError函数来构建错误信息
			errorResponse := utils.NewError(errCode)
			c.Data["json"] = errorResponse
		} else {
			c.Data["json"] = models.Response{Success: true, Message: "更新成功"}
		}
	} else {
		errCode := 508
		// 使用封装的NewError函数来构建错误信息
		errorResponse := utils.NewError(errCode)
		c.Data["json"] = errorResponse
	}
	c.ServeJSON()
}
