package routers

func (r *router) userVisitor() {
	v := r.Group("/visitor")
	v.GET("/value", r.visitor.GetValue)
}
