package main

import (
	"demo/models"
	_ "demo/routers"
	"github.com/astaxie/beego"
)

func main() {
	models.Init()
	//models.CreateRoadSegment()
	//models.CreateParkingSpace()
	beego.Run()
}
