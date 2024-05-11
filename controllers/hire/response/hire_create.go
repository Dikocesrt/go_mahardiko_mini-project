package response

type HireCreateResponse struct {
	Id            int           `json:"id"`
	TotalFee      int           `json:"total_fee"`
	PaymentStatus string        `json:"payment_status"`
	PaymentImage  string        `json:"payment_image"`
	User          UserCreateResponse     `json:"user"`
	Expert        ExpertCreateResponse `json:"expert"`
}

type UserCreateResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

type ExpertCreateResponse struct {
	Id             int    `json:"id"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	FullName       string `json:"full_name"`
}