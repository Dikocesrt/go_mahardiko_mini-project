package request

type ActivityTypeRequest struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
}