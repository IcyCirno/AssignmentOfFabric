package global

import (
	"blockchain/config"
	"sync"

	"github.com/hyperledger/fabric-gateway/pkg/client"
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

	Gw       *client.Gateway
	Contract *client.Contract
	Once     sync.Once
)
