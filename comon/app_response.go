package comon

type successResponse struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

func NewSuccessResponse(data, paging, filter interface{}) *successResponse {
	return &successResponse{data, paging, filter}
}
func SimpleSuccessResponse(data interface{}) *successResponse {
	return &successResponse{data, nil, nil}
}
