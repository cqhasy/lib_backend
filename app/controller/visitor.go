package controller

import (
	"AILN/app/request"
	"AILN/app/response"
	"AILN/app/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Visitor struct{}

var visitorService *service.VisitorService

// @Summary Retrieve a value
// @Description Retrieves the value associated with a given key.
// @Accept json;multipart/form-data
// @Produce json
// @Param key query string true "Key to retrieve value" "
// @Success 200 {object} response.GetValueResponse "Value retrieved successfully"
// @Router /api/v1/visitor/value [get]
func (v *Visitor) GetValue(c *gin.Context) {
	req := &request.GetValueReq{}
	if err := c.ShouldBind(req); err != nil {
		response.FailMsg(c, fmt.Sprintf("parse params error: %v", err))
		return
	}
	value, err := visitorService.GetValue(req.Key)
	if err != nil {
		response.FailMsg(c, fmt.Sprintf("get value error: %v", err))
		return
	}
	response.OkMsgData(c, "get value success", response.GetValueResponse{Value: value})
}
