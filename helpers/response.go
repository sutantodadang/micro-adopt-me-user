package helpers

type ApiResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseJson(msg string, data interface{}) ApiResponse {
	res := &ApiResponse{
		Message: msg,
		Data:    data,
	}

	return *res
}
