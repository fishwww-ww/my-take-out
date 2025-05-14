package request

type EmployeeLogin struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type EmployeePageQueryDTO struct {
	Page     int    `json:"page" form:"page"`         //页数
	PageSize int    `json:"pageSize" form:"pageSize"` //每页容量
	Name     string `json:"name" form:"name"`
}

type EmployeeDTO struct {
	Id       uint64 `json:"id"`                          //员工id
	IdNumber string `json:"idNumber" binding:"required"` //身份证
	Name     string `json:"name" binding:"required"`     //姓名
	Phone    string `json:"phone" binding:"required"`    //手机号
	Sex      string `json:"sex" binding:"required"`      //性别
	UserName string `json:"username" binding:"required"` //用户名
}
