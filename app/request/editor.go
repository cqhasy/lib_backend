package request

type CreateDocumentReq struct {
	Key   string `json:"key" form:"key" bing:"required"`
	Value string `json:"value" form:"value" bing:"required"`
}
