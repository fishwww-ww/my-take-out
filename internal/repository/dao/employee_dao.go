package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"my-take-out/common"
	"my-take-out/common/e"
	"my-take-out/common/retcode"
	"my-take-out/global"
	"my-take-out/internal/api/request"
	"my-take-out/internal/model"
	"time"
)

type EmployeeDao struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewEmployeeDao(db *gorm.DB, redis *redis.Client) *EmployeeDao {
	return &EmployeeDao{
		db:    db,
		redis: redis,
	}
}

func (d *EmployeeDao) GetByUserName(ctx context.Context, username string) (*model.Employee, error) {
	var employee model.Employee
	err := d.db.WithContext(ctx).Where("username=?", username).First(&employee).Error
	if err != nil {
		return nil, retcode.NewError(e.MysqlERR, "Get employee failed")
	}
	return &employee, err
}

func (d *EmployeeDao) Insert(ctx context.Context, entity model.Employee) error {
	err := d.db.WithContext(ctx).Create(&entity).Error
	if err != nil {
		return retcode.NewError(e.MysqlERR, "Insert employee failed")
	}
	return nil
}

func (d *EmployeeDao) PageQuery(ctx context.Context, dto request.EmployeePageQueryDTO) (*common.PageResult, error) {
	// 将 context.Context 转换为 *gin.Context
	ginCtx, ok := ctx.(*gin.Context)
	if !ok {
		return nil, fmt.Errorf("failed to convert context to *gin.Context")
	}

	var err error
	// 尝试从Redis缓存中获取数据
	cacheKey := fmt.Sprintf("employee_page_query:name:%s:page:%d:size:%d",
		dto.Name,
		dto.Page,
		dto.PageSize)
	cachedResult, err := global.Redis.Get(cacheKey).Result()
	if err == nil {
		// 如果缓存命中，将JSON字符串反序列化为PageResult
		var pageResult common.PageResult
		err = json.Unmarshal([]byte(cachedResult), &pageResult)
		if err != nil {
			global.Log.Warn("Failed to unmarshal cached result:", err.Error())
		} else {
			global.Log.Info("Data retrieved from Redis cache")
			retcode.OK(ginCtx, gin.H{
				"data":   pageResult,
				"source": "redis",
			})
			return &pageResult, nil
		}
	}

	// 从mysql中查询
	var result common.PageResult      // 存储查询结果
	var employeeList []model.Employee // 存储员工记录
	query := d.db.WithContext(ctx).Model(&model.Employee{})
	if dto.Name != "" {
		// 模糊查询
		query = query.Where("name LIKE ?", "%"+dto.Name+"%")
	}
	// 查询记录总数
	if err = query.Count(&result.Total).Error; err != nil {
		return nil, retcode.NewError(e.MysqlERR, "Get employee count failed")
	}
	err = query.Scopes(result.Paginate(&dto.Page, &dto.PageSize)).Find(&employeeList).Error
	if err != nil {
		return nil, retcode.NewError(e.MysqlERR, "Get employee list failed")
	}
	result.Records = employeeList
	// 记录数据是从MySQL中获取的
	global.Log.Info("Data retrieved from MySQL database")
	retcode.OK(ginCtx, gin.H{
		"data":   result,
		"source": "mysql",
	})

	// 将查询结果序列化为JSON字符串并存入Redis缓存
	jsonData, err := json.Marshal(result)
	if err != nil {
		global.Log.Warn("Failed to marshal result:", err.Error())
	} else {
		err = global.Redis.Set(cacheKey, jsonData, 10*time.Minute).Err()
		if err != nil {
			global.Log.Warn("Failed to cache result:", err.Error())
		}
	}

	return &result, nil
}
