package routers

func (r *router) userVisitor() {
	v := r.Group("/visitor")
	v.GET("/document", r.visitor.GetDocument)
	v.GET("/document/detail", r.visitor.GetDocumentDetail)
}
