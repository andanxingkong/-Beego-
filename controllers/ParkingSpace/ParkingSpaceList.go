package ParkingSpace

import (
	"demo/controllers"
	"demo/models"
	"demo/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"math"
)

type ParkingSpaceListController struct {
	beego.Controller
}

func (c *ParkingSpaceListController) Get() {
	if !controllers.IsUserLoggedIn(&c.Controller) {
		// 用户未登录、返回错误信息或者跳转到登录页面
		c.Redirect("/login", 302)
		return
	}
	c.TplName = "ParkingSpace/ParkingSpaceList.html"
}

func (c *ParkingSpaceListController) GetAll() {
	o := orm.NewOrm()
	var parkingSpace []models.ParkingSpace
	o.QueryTable("parking_space").All(&parkingSpace)
	if len(parkingSpace) > 0 { //查询到了用户数据
		var respList []interface{}
		for _, parkingSpaces := range parkingSpace {
			respList = append(respList, parkingSpaces.ParkingSpaceToRespDesc())
		}
		page, _ := c.GetInt(":page")
		pageSize, _ := c.GetInt("pageSize", 2)

		if float64(page) > math.Ceil(float64(len(respList))/float64((pageSize))) {
			page = int(math.Ceil(float64(len(respList)) / float64(pageSize)))
		}

		startIndex := (page - 1) * pageSize
		endIndex := startIndex + pageSize

		if endIndex > len(respList) {
			endIndex = len(respList)
		}

		result := respList[startIndex:endIndex]
		fmt.Println("=======================================================================", page)
		c.Data["json"] = result
		c.ServeJSON()
	}
}

func (c *ParkingSpaceListController) DeleteParkingSpace() {
	id, err := c.GetInt(":id") // 获取到 url 当中 id 变量的值
	if err != nil {            // 有错误就返回数据：获取参数失败
		errCode := 512 //获取参数失败
		errorResponse := utils.NewError(errCode)
		c.Data["json"] = errorResponse
		c.ServeJSON()
		return
	}
	o := orm.NewOrm() // 创建一个orm对象
	// 调用 orm 的 Delete 方法，&models.Userinfo{Id: id} 表示删除的是哪一个跟数据库相关的模型以及限制条件
	_, err = o.Delete(&models.ParkingSpace{Id: id})
	if err != nil { // 如果添加错误就返回学信息："删除数据失败"
		errCode := 510 //删除数据失败
		errorResponse := utils.NewError(errCode)
		c.Data["json"] = errorResponse
		c.ServeJSON()
		return
	}
	c.Data["json"] = models.Response{Success: true, Message: "删除成功"}
	c.ServeJSON()
}

func (c *ParkingSpaceListController) FindById() {
	id, _ := c.GetInt(":id")
	var parkingSpace models.ParkingSpace
	o := orm.NewOrm()
	error := o.QueryTable("parking_space").Filter("Id", id).One(&parkingSpace)
	if error != nil {
		// 查询出错
		c.Data["json"] = map[string]interface{}{"error": "road_segment not found"}
		c.ServeJSON()
		return
	}
	// 将用户信息以JSON格式返回给前端
	c.Data["json"] = parkingSpace
	c.ServeJSON()
}

func (c *ParkingSpaceListController) FindBySome() {
	var keyword models.Search
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &keyword)
	if err != nil {
		fmt.Println(err)
		errCode := 504
		// 使用封装的NewError函数来构建错误信息
		errorResponse := utils.NewError(errCode)
		c.Data["json"] = errorResponse
		c.ServeJSON()
		return
	}
	fmt.Println("=====================uuii=", keyword.Mohu)

	var ParkingSpaces []models.ParkingSpace
	o := orm.NewOrm()
	o.QueryTable("parking_space").Filter("Road_id__contains", keyword.Mohu).All(&ParkingSpaces)
	fmt.Println("=====================", ParkingSpaces)
	var respList []interface{}
	for _, ParkingSpace := range ParkingSpaces {
		respList = append(respList, ParkingSpace.ParkingSpaceToRespDesc())
	}
	c.Data["json"] = respList
	c.ServeJSON()
}
