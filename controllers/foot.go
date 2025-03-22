package controllers

import (
	"demo/models"
	"demo/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

//核对停车信息

type RedisController struct {
	beego.Controller
}

func (c *RedisController) Get() {
	if !IsUserLoggedIn(&c.Controller) {
		// 用户未登录、返回错误信息或者跳转到登录页面
		c.Redirect("/login", 302)
		return
	}
	c.TplName = "OrderTable/addOrderTable.html"
}

func (c *RedisController) GetAll() {
	Char.Start_time = time.Now().Format("2006-01-02 15:04:05")
	Char.Cost = 0
	Char.Time = "0"
	fmt.Println(Char)
	c.Data["json"] = Char
	c.ServeJSON()
}

func (c *RedisController) Insider() {
	var roadSegment models.RoadSegment
	//插入车辆
	err := models.InsertVehicle(models.Vehicle{Plate_number: Char.Plate_number, Owner_name: USERNAME, Status: "未离开"})
	//改变路段表状况
	models.SetRoadSegment(Char.Name)
	fmt.Println("__________________________________________1")
	//按条件查询路段表的编号
	o := orm.NewOrm()
	err = o.QueryTable("road_segment").Filter("name", Char.Name).One(&roadSegment)
	fmt.Println(roadSegment)
	//来改变泊位的状态
	fmt.Println("__________________________________________2")
	models.SetParkingSpace(Char.Space_number, roadSegment.Code)
	fmt.Println("__________________________________________3")

	var user models.Users
	var vehicle models.Vehicle
	fmt.Println("__________________________________________4")
	err = o.QueryTable("Users").Filter("Username", USERNAME).One(&user)
	fmt.Println(user)
	fmt.Println("__________________________________________5")
	err1 := o.QueryTable("vehicle").Filter("plate_number", Char.Plate_number).One(&vehicle)
	fmt.Println(vehicle)
	fmt.Println("====================", err1)

	orderTable := models.OrderTable{Name: Char.Name, Space_number: Char.Space_number, Start_time: Char.Start_time, End_time: "", Cost: 0, Status: "未完成"}
	users := models.Users{Id: user.Id}
	o.Read(&users)
	orderTable.User_id = &users
	Vehicles := models.Vehicle{V_id: vehicle.V_id}
	o.Read(&Vehicles)
	orderTable.Vehicle_id = &Vehicles
	o.Insert(&orderTable)
	if err == nil {
		c.Data["json"] = models.Response{Success: true, Message: "添加成功"}
	} else {
		errCode := 509 //添加失败
		errorResponse := utils.NewError(errCode)
		c.Data["json"] = errorResponse
	}
	c.ServeJSON()
}
