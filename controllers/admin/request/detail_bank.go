package request

type BankTypeRequest struct {
	Name string `json:"name" form:"name"`
}