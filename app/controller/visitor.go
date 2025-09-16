package controller

import (
	"AILN/app/request"
	"AILN/app/response"
	"AILN/app/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type Visitor struct{}

var visitorService *service.VisitorService

// @Summary 获取文档
// @Description Retrieves the value associated with a given key.
// @Accept json;multipart/form-data
// @Produce json
// @Param block formData string true "Block of the document" default(SomeBlock)
// @Param group formData string true "Group of the document" default(SomeGroup)
// @Success 200 {object} response.GetDocumentResponse "Value retrieved successfully"
// @Success 200 {object} response.GetSimpleDocumentResponse "Value retrieved successfully"
// @Router /api/v1/visitor/document [get]
func (v *Visitor) GetDocument(c *gin.Context) {
	req := &request.GetDocumentReq{}
	if err := c.ShouldBind(req); err != nil {
		response.FailMsg(c, fmt.Sprintf("parse params error: %v", err))
		return
	}
	docs, err := visitorService.GetDocument(req)
	if err != nil {
		response.FailMsg(c, fmt.Sprintf("get value error: %v", err))
		return
	}
	if len(docs) > 1 {
		var simpleResp response.GetSimpleDocumentResponse
		if err = copier.Copy(&simpleResp.Docs, &docs); err != nil {
			response.FailMsg(c, fmt.Sprintf("copy value error: %v", err))
			return
		}
		response.OkMsgData(c, "get document success", simpleResp)
		return
	}
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Origin", "*")
	response.OkMsgData(c, "get document success", response.GetDocumentResponse{Docs: docs})
}

// @Summary 文档detail
// @Description Retrieve details of a document by its ID
// @Accept json;multipart/form-data
// @Produce json
// @Param id formData uint true "ID of the document to retrieve details" default(123)
// @Success 200 {object} response.GetDocumentDetailResponse
// @Failure 400 {string} string "parse params error or get document detail error"
// @Router /api/v1/visitor/document/detail [get]
func (v *Visitor) GetDocumentDetail(c *gin.Context) {
	req := &request.GetDocumentDetailReq{}
	if err := c.ShouldBind(req); err != nil {
		response.FailMsg(c, fmt.Sprintf("parse params error: %v", err))
		return
	}
	docs, err := visitorService.GetDocumentDetail(req.ID)
	if err != nil {
		response.FailMsg(c, fmt.Sprintf("get document detail error: %v", err))
		return
	}
	response.OkMsgData(c, "get document detail success", response.GetDocumentDetailResponse{Docs: docs})
}

// @Summary 分页获取用户
// @Description get users wiht pagesize and page
// @Accept json;multipart/form-data
// @Produce json
// @Param  page formData int true "page number" default(1)
// @Param  pagesize formData int true "page size" default(10)
// @Success 200 {object} response.GetUsersResponse
// @Failure 400 {string} string "parse params error or get users error"
// @Router /api/v1/visitor/users [get]
func (v *Visitor) UsersInPage(c *gin.Context) {
	req := &request.GetUsersReq{}
	if err := c.ShouldBind(req); err != nil {
		response.FailMsg(c, fmt.Sprintf("parse params error: %v", err))
		return
	}
	users, err := visitorService.GetUsersWithOffset(req.PageNumber, req.PageSize)
	if err != nil {
		response.FailMsg(c, fmt.Sprintf("get users error: %v", err))
		return
	}
	var userInfos []*response.UserInfo
	for _, user := range users {
		userInfos = append(userInfos, &response.UserInfo{
			Name:        user.Username,
			Description: user.Description,
		})
	}
	response.OkMsgData(c, "get UserList success", response.GetUsersResponse{Users: userInfos})
}
