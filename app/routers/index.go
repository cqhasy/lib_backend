package routers

import (
	"AILN/app/controller"
	"github.com/gin-gonic/gin"
)

type router struct {
	*gin.RouterGroup
	auth    *controller.Auth
	editor  *controller.Editor
	visitor *controller.Visitor
}

func Load(e *gin.Engine) {
	r := &router{
		RouterGroup: &e.RouterGroup,
		auth:        &controller.Auth{},
	}
	r.RouterGroup = r.Group("/api/v1")
	r.useAuth()
	r.userEditor()
	r.userVisitor()
}
