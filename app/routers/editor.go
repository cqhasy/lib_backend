package routers

import "AILN/app/core/middlewares"

func (r *router) userEditor() {
	e := r.Group("/editor")
	middlewares.UseJwt(e)
	e.POST("/file", r.editor.UploadFile)
	e.POST("/document", r.editor.CreateDocument)
	e.DELETE("/document", r.editor.DeleteDocument)
}
