package controllers

import (
	"demo/models"
	"demo/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//登录

type UserController struct {
	beego.Controller
}

const (
	NAME  = "login"
	LOGIN = "loginuser"
)

var USERNAME string

func (c *UserController) Get() {
	c.TplName = "login.html"
}

func IsUserLoggedIn(c *beego.Controller) bool {
	// 获取 session 中的用户名
	username := c.GetSession(NAME)
	// 判断用户名是否存在
	if username == nil {
		return false
	}
	return true
}
func (c *UserController) Login() {
	var loginRequest models.Request
	var user models.Users
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &loginRequest)
	if err != nil {
		errCode := 504
		// 使用封装的NewError函数来构建错误信息
		errorResponse := utils.NewError(errCode)
		c.Data["json"] = errorResponse
		c.ServeJSON()
		return
	}
	o := orm.NewOrm()
	o.QueryTable("Users").Filter("password", utils.MD5(loginRequest.Password)).Filter("phone", loginRequest.Phone).One(&user)
	if user.Id > 0 {
		fmt.Println(user.Role_name)
		c.SetSession(NAME, user.Username)
		redisCache := utils.NewRedisCache()
		redisCache.Set(LOGIN, user.Username)
		USERNAME1, _ := redisCache.Get("loginuser")
		USERNAME = USERNAME1.(string)
		fmt.Println(USERNAME)
		if user.Role_name == "用户" {
			c.Data["json"] = models.Response{Success: true, Message: "用户"}
		} else if user.Role_name == "巡查员" {
			c.Data["json"] = models.Response{Success: true, Message: "巡查员"}
		} else if user.Role_name == "管理员" {
			c.Data["json"] = models.Response{Success: true, Message: "管理员"}
		}
	} else {
		errCode := 502
		// 使用封装的NewError函数来构建错误信息
		errorResponse := utils.NewError(errCode)
		c.Data["json"] = errorResponse
	}
	c.ServeJSON()
}
