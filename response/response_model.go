package response

type Response struct {
	Code    uint64      `json:"code" example:"200"`
	Message string      `json:"message" example:"Success."`
	Data    interface{} `json:"data,omitempty"`
}

type ErrResponse struct {
	Code    uint64 `json:"code" example:"200"`
	Message string `json:"message" example:"Success."`
	Error   string `json:"-"`
}

func NewResponse(code uint64, message string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func NewErrResponse(resp ErrResponse, err string) *ErrResponse {
	return &ErrResponse{
		Code:    resp.Code,
		Message: resp.Message,
		Error:   err,
	}
}
