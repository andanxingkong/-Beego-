package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"math"
)

// 订单表
type OrderTable struct {
	R_id         int      `orm:"pk;auto"`
	Name         string   `json:"name"`
	Space_number string   `json:"Space"`
	Start_time   string   `json:"Start_time"`
	End_time     string   `json:"End_time"`
	Cost         float64  `json:"Cost"`
	Status       string   `json:"Status"`
	User_id      *Users   `orm:"rel(fk)"` // 定义外键关联字段
	Vehicle_id   *Vehicle `orm:"rel(fk)"` // 定义外键关联字段
}

func (orderTable *OrderTable) OrderTableToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"R_id":         orderTable.R_id,
		"Name":         orderTable.Name,
		"Space_number": orderTable.Space_number,
		"Start_time":   orderTable.Start_time,
		"End_time":     orderTable.End_time,
		"Cost":         orderTable.Cost,
		"Status":       orderTable.Status,
		"User_id":      *(orderTable.User_id),
		"Vehicle_id":   *(orderTable.Vehicle_id),
	}
	return respInfo
}
func GetOrderTable(page, pageSize int) ([]OrderTable, error) {
	o := orm.NewOrm()
	count, _ := o.QueryTable("order_table").Count()
	fmt.Println("总数量：", count)
	// 计算总页数
	totalPage := int(math.Ceil(float64(count) / float64(pageSize)))
	// 确保当前页码在有效范围内
	if page < 1 {
		page = 1
	} else if page > totalPage {
		page = totalPage
	}
	var orderTable []OrderTable
	_, err := o.QueryTable("order_table").Limit(pageSize, (page-1)*pageSize).All(&orderTable)
	return orderTable, err
}
