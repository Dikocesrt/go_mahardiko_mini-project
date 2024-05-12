package request

type ExpertiseRequest struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
}