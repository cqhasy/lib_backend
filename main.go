package main

import (
	"AILN/app/config"
	"AILN/app/routers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	config.Info = config.New("./config/config.yaml")
	engine := gin.Default()
	routers.Load(engine)
	if err := engine.Run("0.0.0.0:8080"); err != nil {
		log.Fatal(err)
	}
}
