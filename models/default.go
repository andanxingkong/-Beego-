package models

import (
	"demo/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func Init() {
	driverName := beego.AppConfig.String("driverName")
	//注册数据库驱动
	orm.RegisterDriver(driverName, orm.DRMySQL)
	//数据库连接
	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")

	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"

	err := orm.RegisterDataBase("default", driverName, dbConn)
	if err != nil {
		utils.LogError("连接数据库出错")
		return
	}
	utils.LogInfo("连接数据库成功")

	//register model : 注册实体模型
	orm.RegisterModel(
		new(Users),
		new(ParkingSpace),
		new(RoadSegment),
		new(Vehicle),
		new(OrderTable),
	)
	//the last step: create table
	orm.RunSyncdb("default", false, true)
}
