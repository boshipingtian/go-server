package global

import (
	"go-server/config"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	DB     *gorm.DB
)
