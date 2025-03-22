package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"math"
	"strconv"
)

// 路段表
type RoadSegment struct {
	Road_id int    `orm:"pk;auto"`
	Name    string `json:"name"`
	Code    string `json:"code"`
	Status  string `json:"Status"`
}

func (roadSegment *RoadSegment) RoadSegmentToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"Road_id": roadSegment.Road_id,
		"Name":    roadSegment.Name,
		"Code":    roadSegment.Code,
		"Status":  roadSegment.Status,
	}
	return respInfo
}
func GetRoadSegment(page, pageSize int) ([]RoadSegment, error) {
	o := orm.NewOrm()
	count, _ := o.QueryTable("road_segment").Count()
	fmt.Println("总数量：", count)
	// 计算总页数
	totalPage := int(math.Ceil(float64(count) / float64(pageSize)))
	// 确保当前页码在有效范围内
	if page < 1 {
		page = 1
	} else if page > totalPage {
		page = totalPage
	}
	var roadSegment []RoadSegment
	_, err := o.QueryTable("road_segment").Limit(pageSize, (page-1)*pageSize).All(&roadSegment)
	return roadSegment, err
}

// 插入  he
func InsertRoadSegment(roadSegment RoadSegment) {
	o := orm.NewOrm()
	var r orm.RawSeter
	r = o.Raw("insert into road_segment(name,code,status) values (?,?,?)", roadSegment.Name, roadSegment.Code, roadSegment.Status)
	res, err := r.Exec()
	if err != nil {
		fmt.Println("插入失败")
	} else {
		fmt.Println(res.RowsAffected())
	}
}
func CreateRoadSegment() {
	for i := 1; i <= 3; i++ {
		InsertRoadSegment(RoadSegment{Name: "科技" + strconv.Itoa(i) + "路", Code: "路" + strconv.Itoa(i), Status: "可用"})
	}
}

// 更新数据  he
func SetRoadSegment(roadSegment string) {
	o := orm.NewOrm()
	var r orm.RawSeter
	r = o.Raw("update road_segment set status=? where name=?", "不可用", roadSegment)
	res, err := r.Exec()
	if err != nil {
		fmt.Println("更新失败")
	} else {
		fmt.Println(res.RowsAffected())
	}
}
