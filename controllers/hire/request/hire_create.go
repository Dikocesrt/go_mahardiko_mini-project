package request

type HireCreateRequest struct {
	UserId       int    `json:"user_id" form:"user_id"`
	ExpertId     int    `json:"expert_id" form:"expert_id"`
	TotalFee     int    `json:"total_fee" form:"total_fee"`
	PaymentImage string `json:"payment_image" form:"payment_image"`
}