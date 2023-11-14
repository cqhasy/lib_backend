package routers

func (r *router) useAuth() {
	authRouter := r.Group("auth")
	authRouter.POST("/login", r.auth.Login)
}
