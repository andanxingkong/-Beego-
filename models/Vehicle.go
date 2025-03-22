package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"math"
)

// 车辆表
type Vehicle struct {
	V_id         int           `orm:"pk;auto"`
	Plate_number string        `json:"plate_number"`
	Owner_name   string        `json:"owner_name"`
	Status       string        `json:"status"`
	OrderTable   []*OrderTable `orm:"reverse(many)"` // 定义反向关联字段
}

// -----------数据库操作---------------
func (vehicle *Vehicle) VehicleToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"V_id":         vehicle.V_id,
		"Plate_number": vehicle.Plate_number,
		"Owner_name":   vehicle.Owner_name,
		"Status":       vehicle.Status,
	}
	return respInfo
}

func GetVehicle(page, pageSize int) ([]Vehicle, error) {
	o := orm.NewOrm()
	count, _ := o.QueryTable("Vehicle").Count()
	fmt.Println("总数量：", count)
	// 计算总页数
	totalPage := int(math.Ceil(float64(count) / float64(pageSize)))
	// 确保当前页码在有效范围内
	if page < 1 {
		page = 1
	} else if page > totalPage {
		page = totalPage
	}
	var vehicle []Vehicle
	_, err := o.QueryTable("Vehicle").Limit(pageSize, (page-1)*pageSize).All(&vehicle)
	return vehicle, err
}

// 插入  he
func InsertVehicle(vehicle Vehicle) error {
	o := orm.NewOrm()
	var r orm.RawSeter
	r = o.Raw("insert into Vehicle(plate_number,owner_name,status) values(?,?,?)", vehicle.Plate_number, vehicle.Owner_name, vehicle.Status)
	res, err := r.Exec()
	if err != nil {
		fmt.Println("插入失败")
	} else {
		fmt.Println(res.RowsAffected())
	}
	return err
}
