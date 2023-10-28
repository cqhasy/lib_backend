package routers

import (
	"AILN/app/controller"
	"github.com/gin-gonic/gin"
)

type router struct {
	*gin.RouterGroup
	auth *controller.Auth
}

func Load(e *gin.Engine) {
	r := &router{
		RouterGroup: &e.RouterGroup,
		auth:        &controller.Auth{},
	}
	r.useAuth()
}
