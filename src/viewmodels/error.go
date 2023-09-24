package viewmodels

const (
	errorBadRequest int32 = 400
	errorNotFound   int32 = 404
)

type errorResponse struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

func NewBadRequest(message string) *errorResponse {
	return &errorResponse{
		Code:    errorBadRequest,
		Message: message,
	}
}

func NewNotFound(message string) *errorResponse {
	return &errorResponse{
		Code:    errorNotFound,
		Message: message,
	}
}
