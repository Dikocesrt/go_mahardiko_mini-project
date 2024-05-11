package request

type HireVerifyRequest struct {
	Id            int       `json:"id"`
	PaymentStatus string    `json:"payment_status"`
	MeetUrl       string    `json:"meet_url"`
}