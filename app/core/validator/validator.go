package validator

import (
	"AILN/app/core/container"
	"AILN/app/core/validator/checker"
	"AILN/app/lib/xerr"
	"github.com/gin-gonic/gin"
	"log"
)

type Validator interface {
	CheckParams(context *gin.Context)
}

func Get(key string) func(context *gin.Context) {

	if value := container.NewContainer().Get(key); value != nil {
		if val, isOk := value.(Validator); isOk {
			return val.CheckParams
		}
	}
	log.Fatal(xerr.ErrorsValidatorNotExists + ", 验证器模块：" + key)
	return nil
}

func init() {
	//创建容器
	containers := container.NewContainer()

	var key string

	key = "SendEmailCode"
	containers.Set(key, &checker.SendEmailChecker{})

}
