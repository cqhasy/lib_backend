package main

import (
	"AILN/app/common"
	"AILN/app/core/config"
	"AILN/app/core/gorm"
	"AILN/app/core/middlewares"
	"AILN/app/core/zap"
	"AILN/app/routers"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/cliutil"
	"log"
)

func main() {
	common.CONFIG = config.New("./config.toml")
	common.LOG = zap.AddZap()
	DB, err := gorm.NewGorm()
	if err != nil {
		cliutil.Errorln(err)
		return
	}
	common.DB = DB
	engine := gin.Default()
	middlewares.Load(engine)
	routers.Load(engine)

	if err := engine.Run("0.0.0.0:30000"); err != nil {
		log.Fatal(err)
	}
}
