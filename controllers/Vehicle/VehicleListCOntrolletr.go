package Vehicle

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

type VehicleListController struct {
	beego.Controller
}

func (c *VehicleListController) Get() {
	if !controllers.IsUserLoggedIn(&c.Controller) {
		// 用户未登录、返回错误信息或者跳转到登录页面
		c.Redirect("/login", 302)
		return
	}
	c.TplName = "Vehicle/VehicleList.html"
}

func (c *VehicleListController) GetAll() {
	o := orm.NewOrm()
	var vehicles []models.Vehicle
	o.QueryTable("Vehicle").All(&vehicles)
	if len(vehicles) > 0 { //查询到了用户数据
		var respList []interface{}
		for _, vehicle := range vehicles {
			respList = append(respList, vehicle.VehicleToRespDesc())
		}
		page, _ := c.GetInt(":page")
		pageSize, _ := c.GetInt("pageSize", 3)

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
func (c *VehicleListController) DeleteVehicle() {
	id, err := c.GetInt(":id") // 获取到 url 当中 id 变量的值
	if err != nil {            // 有错误就返回数据：获取参数失败
		errCode := 512 //获取参数失败
		// 使用封装的NewError函数来构建错误信息
		errorResponse := utils.NewError(errCode)
		c.Data["json"] = errorResponse
		c.ServeJSON()
		return
	}
	o := orm.NewOrm() // 创建一个orm对象
	// 调用 orm 的 Delete 方法，&models.Userinfo{Id: id} 表示删除的是哪一个跟数据库相关的模型以及限制条件
	_, err = o.Delete(&models.Vehicle{V_id: id})
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

func (c *VehicleListController) FindById() {
	id, _ := c.GetInt(":id")
	fmt.Println("find", id)
	var vehicle models.Vehicle
	o := orm.NewOrm()
	error := o.QueryTable("Vehicle").Filter("v_id", id).One(&vehicle)
	fmt.Println(vehicle)
	if error != nil {
		// 查询出错
		c.Data["json"] = map[string]interface{}{"error": "User not found"}
		c.ServeJSON()
		return
	}
	// 将用户信息以JSON格式返回给前端
	c.Data["json"] = vehicle
	c.ServeJSON()
}
func (c *VehicleListController) FindBySome() {
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

	var Vehicles []models.Vehicle
	o := orm.NewOrm()
	o.QueryTable("Vehicle").Filter("Plate_number__contains", keyword.Mohu).All(&Vehicles)
	fmt.Println("=====================", Vehicles)
	var respList []interface{}
	for _, Vehicle := range Vehicles {
		respList = append(respList, Vehicle.VehicleToRespDesc())
	}
	c.Data["json"] = respList
	c.ServeJSON()
}
