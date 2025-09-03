package global

import (
	"blockchain/config"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	TOKEN_NAME   = "Authorization"
	TOKEN_PREFIX = "Bearer: "
)

var (
	DB          *gorm.DB
	RedisClient *config.RedisClient
	Logger      *zap.SugaredLogger
)
