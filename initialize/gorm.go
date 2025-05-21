package initialize

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"my-take-out/internal/model"
)

var (
	GormToManyRequestError = errors.New("gorm: to many request")
)

func InitDatabase(dsn string) *gorm.DB {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(0)
	fmt.Println("Final DSN ===>", dsn)

	// 自动迁移（建表）
	err = db.AutoMigrate(&model.Employee{})
	if err != nil {
		log.Fatalf("自动建表失败: %v", err)
	}
	log.Println("表创建成功")
	initAdminUser(db)
	return db
}

// 初始化管理员用户
func initAdminUser(db *gorm.DB) {
	admin := model.Employee{
		Username: "admin",
		Password: "e10adc3949ba59abbe56e057f20f883e", // 123456 的 MD5
		Status:   1,
	}

	// 使用 FirstOrCreate 避免重复创建
	result := db.Where(model.Employee{Username: admin.Username}).
		Attrs(model.Employee{
			Password: admin.Password,
			Status:   admin.Status,
		}).
		FirstOrCreate(&admin)

	if result.Error != nil {
		log.Printf("初始化管理员用户失败: %v", result.Error)
	} else if result.RowsAffected > 0 {
		log.Println("管理员用户创建成功")
	} else {
		log.Println("管理员用户已存在")
	}
}
