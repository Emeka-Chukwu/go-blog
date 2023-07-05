package util

type DataResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func FormatErrorResponse(message string, data any) ErrorResponse {
	return ErrorResponse{
		Message: message, Data: data,
	}
}
