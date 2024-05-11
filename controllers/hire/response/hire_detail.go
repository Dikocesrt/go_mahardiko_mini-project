package response

type HireDetailResponse struct {
	Id            int                  `json:"id"`
	HireStart     string               `json:"hire_start"`
	HireEnd       string               `json:"hire_end"`
	TotalFee      int                  `json:"total_fee"`
	PaymentStatus string               `json:"payment_status"`
	PaymentImage  string               `json:"payment_image"`
	MeetTime      string               `json:"meet_time"`
	MeetDay       string               `json:"meet_day"`
	MeetUrl       string               `json:"meet_url"`
	User          UserDetailResponse   `json:"user"`
	Expert        ExpertDetailResponse `json:"expert"`
}

type UserDetailResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

type ExpertDetailResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}