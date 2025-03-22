package routers

import (
	"demo/controllers"
	"demo/controllers/OrderTable"
	"demo/controllers/ParkingSpace"
	"demo/controllers/RoadSegment"
	"demo/controllers/User"
	"demo/controllers/Vehicle"
	"github.com/astaxie/beego"
)

func init() {
	//用户主页
	beego.Router("/", &controllers.HomeController{})
	//管理员主页
	beego.Router("/main", &controllers.MainController{})
	//登录
	beego.Router("/login", &controllers.UserController{})
	beego.Router("/dologin", &controllers.UserController{}, "POST:Login")
	//用户订单和结束订单
	beego.Router("/UserOrderTableGetAll", &controllers.GatAllController{})
	beego.Router("/doUserOrderTableGetAll", &controllers.GatAllController{}, "POST:GetAll")
	beego.Router("/UPUserOrderTableGetAll/:id", &controllers.GatAllController{}, "PUT:UpOrderTable")
	//巡查员查询全部订单
	beego.Router("/PatrolPersonQuery", &controllers.PatrolPersonQueryController{})
	beego.Router("/doPatrolPersonQuery", &controllers.PatrolPersonQueryController{}, "POST:GetAll")
	//退出
	beego.Router("/exit", &controllers.ExitController{})
	//个人中心
	beego.Router("/center", &controllers.CenterController{})
	beego.Router("/centerVehicle", &controllers.CenterVehicleController{})
	beego.Router("/doAddCenterVehicle", &controllers.CenterVehicleController{}, "POST:AddCenterVehicle")
	beego.Router("/upPassword", &controllers.UpPasswordController{})
	beego.Router("/doupPassword", &controllers.UpPasswordController{}, "PUT:UpPasswordRequest")
	//注册
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/doregister", &controllers.RegisterController{}, "POST:Register")
	//停车
	beego.Router("/char", &controllers.CharController{})
	beego.Router("/dochar", &controllers.CharController{}, "POST:Addchar")
	//核对停车信息
	beego.Router("/confirm", &controllers.RedisController{})
	beego.Router("/doConfirm", &controllers.RedisController{}, "GET:GetAll")
	beego.Router("/Insider", &controllers.RedisController{}, "POST:Insider")
	//用户管理的查询和修改和删除和添加
	beego.Router("/userlist", &User.UesrListController{})
	beego.Router("/getall/:page", &User.UesrListController{}, "GET:GetAll")
	beego.Router("/userGetOne/:id", &User.UesrListController{}, "GET:FindById")
	beego.Router("/userGetSome", &User.UesrListController{}, "POST:FindBySome")
	beego.Router("/deleteuser/:id", &User.UesrListController{}, "DELETE:DeleteUser")
	beego.Router("/updateUser/:id", &User.UpdateController{}, "GET:Getupdate")
	beego.Router("/doupdateUser", &User.UpdateController{}, "PUT:UpdateUser")
	beego.Router("/adduser", &User.AdduserController{})
	beego.Router("/doadduser", &User.AdduserController{}, "POST:Adduser")
	//路段管理的查询和修改和删除和添加
	beego.Router("/roadsegmentlist", &RoadSegment.RoadSegmentlist{})
	beego.Router("/RoadSegmentGetAll/:page", &RoadSegment.RoadSegmentlist{}, "GET:GetAll")
	beego.Router("/RoadSegmentGetOne/:id", &RoadSegment.RoadSegmentlist{}, "GET:FindById")
	beego.Router("/RoadSegmentGetSome", &RoadSegment.RoadSegmentlist{}, "POST:FindBySome")
	beego.Router("/deleteRoadSegment/:id", &RoadSegment.RoadSegmentlist{}, "DELETE:DeleteRoadSegment")
	beego.Router("/updateRoadSegment/:id", &RoadSegment.UpdateRoadSegmentController{}, "GET:Getupdate")
	beego.Router("/doupdateRoadSegment", &RoadSegment.UpdateRoadSegmentController{}, "PUT:UpdateRoadSegment")
	beego.Router("/addRoadSegment", &RoadSegment.AddRoadSegmentController{})
	beego.Router("/doAddRoadSegment", &RoadSegment.AddRoadSegmentController{}, "POST:AddRoadSegment")
	//泊位管理的查询和修改和删除和添加
	beego.Router("/ParkingSpaceList", &ParkingSpace.ParkingSpaceListController{})
	beego.Router("/ParkingSpaceGetAll/:page", &ParkingSpace.ParkingSpaceListController{}, "GET:GetAll")
	beego.Router("/ParkingSpaceGetOne/:id", &ParkingSpace.ParkingSpaceListController{}, "GET:FindById")
	beego.Router("/ParkingSpaceGetSome", &ParkingSpace.ParkingSpaceListController{}, "POST:FindBySome")
	beego.Router("/deleteParkingSpace/:id", &ParkingSpace.ParkingSpaceListController{}, "DELETE:DeleteParkingSpace")
	beego.Router("/updateParkingSpace/:id", &ParkingSpace.UpdateParkingSpaceController{}, "GET:Getupdate")
	beego.Router("/doupdateParkingSpace", &ParkingSpace.UpdateParkingSpaceController{}, "PUT:UpdateParkingSpace")
	beego.Router("/addParkingSpace", &ParkingSpace.AddParkingSpaceController{})
	beego.Router("/doAddRParkingSpace", &ParkingSpace.AddParkingSpaceController{}, "POST:AddParkingSpace")
	//车辆管理的查询和修改和删除和添加
	beego.Router("/VehicleList", &Vehicle.VehicleListController{})
	beego.Router("/VehicleGetAll/:page", &Vehicle.VehicleListController{}, "GET:GetAll")
	beego.Router("/VehicleGetOne/:id", &Vehicle.VehicleListController{}, "GET:FindById")
	beego.Router("/VehicleGetSome", &Vehicle.VehicleListController{}, "POST:FindBySome")
	beego.Router("/deleteVehicle/:id", &Vehicle.VehicleListController{}, "DELETE:DeleteVehicle")
	beego.Router("/updateVehicle/:id", &Vehicle.UpdateVehicleListController{}, "GET:Getupdate")
	beego.Router("/doupdateVehicle", &Vehicle.UpdateVehicleListController{}, "PUT:UpdateVehicle")
	beego.Router("/addVehicle", &Vehicle.AddVehicleController{})
	beego.Router("/doAddVehicle", &Vehicle.AddVehicleController{}, "POST:AddVehicle")
	//订单管理的查询和修改和删除
	beego.Router("/OrderTableList", &OrderTable.OrderTableListController{})
	beego.Router("/OrderTableGetAll/:page", &OrderTable.OrderTableListController{}, "GET:GetAll")
	beego.Router("/OrderTableGetOne/:id", &OrderTable.OrderTableListController{}, "GET:FindById")
	beego.Router("/OrderTableGetSome", &OrderTable.OrderTableListController{}, "POST:FindBySome")
	beego.Router("/deleteOrderTable/:id", &OrderTable.OrderTableListController{}, "DELETE:DeleteOrderTable")
	beego.Router("/updateOrderTable/:id", &OrderTable.UpdateOrderTableListController{}, "GET:Getupdate")
	beego.Router("/doupdateOrderTable", &OrderTable.UpdateOrderTableListController{}, "PUT:UpdateOrderTable")
}
