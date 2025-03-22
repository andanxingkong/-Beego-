package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"math"
	"strconv"
)

// 泊位表
type ParkingSpace struct {
	Id           int    `orm:"pk;auto"`
	Road_id      string `json:"Road_id"`
	Space_number string `json:"Space_number"`
	Status       string `json:"Status"`
}

func (parkingSpace *ParkingSpace) ParkingSpaceToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"Id":           parkingSpace.Id,
		"Road_id":      parkingSpace.Road_id,
		"Space_number": parkingSpace.Space_number,
		"Status":       parkingSpace.Status,
	}
	return respInfo
}

//--------------数据库操作-----------------

func GetParkingSpace(page, pageSize int) ([]ParkingSpace, error) {
	o := orm.NewOrm()
	count, _ := o.QueryTable("parking_space").Count()
	fmt.Println("总数量：", count)
	// 计算总页数
	totalPage := int(math.Ceil(float64(count) / float64(pageSize)))
	// 确保当前页码在有效范围内
	if page < 1 {
		page = 1
	} else if page > totalPage {
		page = totalPage
	}
	var parkingSpace []ParkingSpace
	_, err := o.QueryTable("parking_space").Limit(pageSize, (page-1)*pageSize).All(&parkingSpace)
	return parkingSpace, err
}

// 插入  he
func InsertParkingSpace(parkingSpace ParkingSpace) {
	o := orm.NewOrm()
	var r orm.RawSeter
	r = o.Raw("insert into parking_space(road_id,space_number,status) values (?,?,?)",
		parkingSpace.Road_id, parkingSpace.Space_number, parkingSpace.Status)
	res, err := r.Exec()
	if err != nil {
		fmt.Println("插入失败")
	} else {
		fmt.Println(res.RowsAffected())
	}
}

func CreateParkingSpace() {
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			InsertParkingSpace(ParkingSpace{Road_id: "路" + strconv.Itoa(i), Space_number: "A-00" + strconv.Itoa(j), Status: "空闲"})
		}
	}

}

// update 表名 set 字段名=更改的值 where 条件表达式
func SetParkingSpace(space, road_id string) {
	o := orm.NewOrm()
	var r orm.RawSeter
	r = o.Raw("update parking_space set status=? where space_number=? and road_id=?",
		"占用", space, road_id)
	res, err := r.Exec()
	if err != nil {
		fmt.Println("更新失败")
	} else {
		fmt.Println(res.RowsAffected())
	}
}
