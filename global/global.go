package global

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"my-take-out/config"
	"my-take-out/logger"
)

var (
	Config *config.AllConfig // 全局Config
	Log    logger.ILog
	DB     *gorm.DB
	Redis  *redis.Client
)
