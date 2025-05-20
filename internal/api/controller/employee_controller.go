package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"my-take-out/common"
	"my-take-out/common/retcode"
	"my-take-out/global"
	"my-take-out/internal/api/request"
	"my-take-out/internal/service"
	"time"
)

type EmployeeController struct {
	service service.IEmployeeService
}

func NewEmployeeController(employeeService service.IEmployeeService) *EmployeeController {
	return &EmployeeController{service: employeeService}
}

func (ec *EmployeeController) Login(ctx *gin.Context) {
	employeeLogin := request.EmployeeLogin{}
	err := ctx.Bind(&employeeLogin)
	if err != nil {
		global.Log.Debug("EmployeeController login 解析失败")
		retcode.Fatal(ctx, err, "")
		return
	}
	resp, err := ec.service.Login(ctx, employeeLogin)
	if err != nil {
		global.Log.Warn("EmployeeController login Error:", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	retcode.OK(ctx, resp)
}

func (ec *EmployeeController) AddEmployee(ctx *gin.Context) {
	employee := request.EmployeeDTO{}
	err := ctx.Bind(&employee)
	if err != nil {
		global.Log.Debug("AddEmployee Error:", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	err = ec.service.CreateEmployee(ctx, employee)
	if err != nil {
		global.Log.Warn("AddEmployee Error", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	retcode.OK(ctx, "")
}

func (ec *EmployeeController) PageQuery(ctx *gin.Context) {
	var employeePageQueryDTO request.EmployeePageQueryDTO
	err := ctx.Bind(&employeePageQueryDTO)
	if err != nil {
		global.Log.Error("AddEmployee  invalid params err:", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	// 尝试从Redis缓存中获取数据
	cacheKey := "employee_page_query:" + employeePageQueryDTO.Name
	cachedResult, err := global.Redis.Get(cacheKey).Result()
	if err == nil {
		// 如果缓存命中，将JSON字符串反序列化为PageResult
		var pageResult common.PageResult
		err = json.Unmarshal([]byte(cachedResult), &pageResult)
		if err != nil {
			global.Log.Warn("Failed to unmarshal cached result:", err.Error())
		} else {
			global.Log.Info("Data retrieved from Redis cache")
			retcode.OK(ctx, gin.H{
				"data":   pageResult,
				"source": "redis",
			})
			return
		}
	}
	// 进行分页查询
	pageResult, err := ec.service.PageQuery(ctx, employeePageQueryDTO)
	if err != nil {
		global.Log.Warn("AddEmployee  Error:", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	// 将查询结果序列化为JSON字符串并存入Redis缓存
	jsonData, err := json.Marshal(pageResult)
	if err != nil {
		global.Log.Warn("Failed to marshal result:", err.Error())
	} else {
		err = global.Redis.Set(cacheKey, jsonData, 10*time.Minute).Err()
		if err != nil {
			global.Log.Warn("Failed to cache result:", err.Error())
		}
	}

	// 记录数据是从MySQL中获取的
	global.Log.Info("Data retrieved from MySQL database")
	retcode.OK(ctx, gin.H{
		"data":   pageResult,
		"source": "mysql",
	})
	// retcode.OK(ctx, pageResult)
}
