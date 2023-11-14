package common

import (
	"github.com/gookit/config/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var (
	DB        *gorm.DB
	CONFIG    *config.Config
	LOG       *zap.Logger
	StartTime time.Time
)
