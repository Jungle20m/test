package common

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Error []*Response `json:"errors"`
}

func HttpResponse(http_code int, message string, data interface{}) *Response {
	return &Response{
		Code:    http_code,
		Message: message,
		Data:    data,
	}
}

func HttpErrorResponse(error_object []*Response) *ErrorResponse {
	var errors []*Response
	for index := 0; index < len(error_object); index++ {
		errors = append(errors, error_object[index])
	}
	return &ErrorResponse{
		Error: errors,
	}
}
