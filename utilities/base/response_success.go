package base

type BaseSuccessResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func NewSuccessResponse(message string, data any) *BaseSuccessResponse {
	return &BaseSuccessResponse{
		Status:  true,
		Message: message,
		Data:    data,
	}
}
