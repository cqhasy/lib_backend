package routers

func (r *router) useAuth() {
	authRouter := r.Group("auth")
	authRouter.GET("/email-code", r.auth.SendEmailCode)
}
