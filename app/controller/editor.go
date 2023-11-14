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
// @Router /api/v1/editor/upload [post]
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
// @Param key formData string true "Key for the document"
// @Param value formData string true "Value for the document"
// @Security ApiKeyAuth
// @Success 200 {string} string "Document created successfully"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Router /api/v1/editor/document/create [post]
func (e *Editor) CreateDocument(c *gin.Context) {
	req := &request.CreateDocumentReq{}
	if err := c.ShouldBind(req); err != nil {
		response.FailMsg(c, fmt.Sprintf("parse params error: %v", err))
		return
	}
	if err := editorService.CreateDocument(req.Key, req.Value); err != nil {
		response.FailMsg(c, fmt.Sprintf("create document error: %v", err))
		return
	}
	response.OkMsg(c, "create document success")
}
