package global

import (
	"gorm.io/gorm"
	"my-take-out/config"
)

var (
	Config *config.AllConfig // 全局Config
	DB     *gorm.DB
)
