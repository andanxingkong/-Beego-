package controllers

import (
	"demo/models"
	"demo/utils"
	"encoding/json"
	"github.com/astaxie/beego"
)

//泊位

type CharController struct {
	beego.Controller
}

type CharRequest struct {
	Plate_number string  `json:"plate_number"`
	Name         string  `json:"name"`
	Space_number string  `json:"space_number"`
	Start_time   string  `json:"start_time"`
	Time         string  `json:"time"`
	Cost         float64 `json:"cost"`
}

var Char CharRequest

func (c *CharController) Get() {
	if !IsUserLoggedIn(&c.Controller) {
		// 用户未登录、返回错误信息或者跳转到登录页面
		c.Redirect("/login", 302)
		return
	}
	c.TplName = "write_article.html"
}

func (c *CharController) Addchar() {
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &Char)
	if err != nil {
		errCode := 504
		errorResponse := utils.NewError(errCode)
		c.Data["json"] = errorResponse
		c.ServeJSON()
		return
	}
	if err == nil {
		c.Data["json"] = models.Response{Success: true, Message: "添加成功"}
	} else {
		errCode := 509 //添加失败
		errorResponse := utils.NewError(errCode)
		c.Data["json"] = errorResponse
	}
	c.ServeJSON()
}
