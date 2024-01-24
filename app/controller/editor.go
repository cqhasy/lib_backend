package controller

import (
	"AILN/app/common/tube"
	"AILN/app/request"
	"AILN/app/response"
	"AILN/app/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Editor struct{}

var editorService *service.EditorService

// @Summary 上传文件
// @Description Uploads a file and returns its URL upon successful upload.
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File to upload"
// @Security ApiKeyAuth
// @Success 200 {object} response.UploadFile "upload file success"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Router /api/v1/editor/file [post]
func (e *Editor) UploadFile(c *gin.Context) {
	multipartForm, err := c.MultipartForm()
	if err != nil {
		response.FailMsg(c, fmt.Sprintf("get multipart form error: %v", err))
		return
	}
	fileurls := make([]string, 0)
	for _, headers := range multipartForm.File {
		for _, header := range headers {
			file, err := header.Open()
			if err != nil {
				response.FailMsg(c, fmt.Sprintf("open file error: %v", err))
				return
			}
			fileurl, err := tube.UploadFile(header.Filename, c.GetUint("userID"), file, header.Size)
			if err != nil {
				response.FailMsg(c, fmt.Sprintf("upload file error: %v", err))
				return
			}
			fileurls = append(fileurls, fileurl)
		}
	}
	response.OkMsgData(c, "upload file success", response.UploadFile{FileURL: fileurls})
}

// @Summary 创建/更新文档
// @Description Creates a new document with a specified key and value.
// @Accept json;multipart/form-data
// @Produce json
// @Param block formData string true "Block of the document" default(SomeBlock)
// @Param group formData string true "Group of the document" default(SomeGroup)
// @Param title formData string true "Title of the document" default(SomeTitle)
// @Param create_at formData int64 true "Creation timestamp of the document" default(1636972168)
// @Param content formData string true "Content of the document" default(Some content)
// @Security ApiKeyAuth
// @Success 200 {string} string "Document created successfully"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Router /api/v1/editor/document [post]
func (e *Editor) CreateDocument(c *gin.Context) {
	req := &request.CreateDocumentReq{}
	if err := c.ShouldBind(req); err != nil {
		response.FailMsg(c, fmt.Sprintf("parse params error: %v", err))
		return
	}
	if err := editorService.CreateDocument(req); err != nil {
		response.FailMsg(c, fmt.Sprintf("create document error: %v", err))
		return
	}
	response.OkMsg(c, "create document success")
}

// @Summary 删掉文档
// @Description Delete a document by its ID
// @Accept json;multipart/form-data
// @Produce json
// @Param id formData uint true "ID of the document to delete" default(123)
// @Success 200 {string} string "delete document success"
// @Router /api/v1/editor/document [delete]
func (e *Editor) DeleteDocument(c *gin.Context) {
	req := &request.DeleteDocumentReq{}
	if err := c.ShouldBind(req); err != nil {
		response.FailMsg(c, fmt.Sprintf("parse params error: %v", err))
		return
	}
	if err := editorService.DeleteDocument(req.ID); err != nil {
		response.FailMsg(c, fmt.Sprintf("delete document error: %v", err))
		return
	}
	response.OkMsg(c, "delete document success")
}
