package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"math"
)

// he
type Users struct {
	Id         int           `orm:"pk;auto"`
	Username   string        `json:"username"`
	Phone      string        `json:"phone"`
	Password   string        `json:"password"`
	Email      string        `json:"email"`
	Role_name  string        `json:"role"`
	OrderTable []*OrderTable `orm:"reverse(many)"` // 定义反向关联字段
}

// 表示服务器在处理登录请求后发送回客户端的响应数据格式
type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// 表示用户尝试登录时期望的请求数据格式
type Request struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type UpRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Role_name string `json:"role"`
}

type UnifiedResponse struct {
	Code    int         `json:"code"`    // 响应状态码
	Message string      `json:"message"` // 响应消息
	Data    interface{} `json:"data"`    // 响应数据
}

type UppasswordRequest struct {
	Phone    string `json:"Phone"`
	Password string `json:"password"`
}

// 表示服务器在处理登录请求后发送回客户端的响应数据格式
type Search struct {
	Mohu string `json:"mohu"`
}

func (user *Users) UserToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"Id":        user.Id,
		"Username":  user.Username,
		"Phone":     user.Phone,
		"Password":  user.Password,
		"Email":     user.Email,
		"Role_name": user.Role_name,
	}
	return respInfo
}
func GetUsers(page, pageSize int) ([]Users, error) {
	o := orm.NewOrm()
	count, _ := o.QueryTable("Users").Count()
	fmt.Println("总数量：", count)
	// 计算总页数
	totalPage := int(math.Ceil(float64(count) / float64(pageSize)))
	// 确保当前页码在有效范围内
	if page < 1 {
		page = 1
	} else if page > totalPage {
		page = totalPage
	}
	var users []Users
	_, err := o.QueryTable("Users").Limit(pageSize, (page-1)*pageSize).All(&users)
	return users, err
}

// 查询用户及其对应的所有文章
func GetUserWithPosts(username string) (*Users, error) {
	var userid Users
	o := orm.NewOrm()
	// 查询用户
	o.QueryTable("Users").Filter("Username", username).One(&userid)
	user := &Users{Id: userid.Id}
	err := o.Read(user)
	if err != nil {
		return nil, err
	}
	// 加载用户的订单
	_, err = o.LoadRelated(user, "OrderTable") //属性名
	if err != nil {
		return nil, err
	}
	return user, nil
}
