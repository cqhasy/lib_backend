package routers

import "AILN/app/core/middlewares"

func (r *router) userEditor() {
	e := r.Group("/editor")
	middlewares.UseJwt(e)
	e.POST("/file", r.editor.UploadFile)
	e.POST("/doc", r.editor.CreateDocument)
}
