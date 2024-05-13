package request

type HireVerifyRequest struct {
	Id      int    `json:"id"`
	MeetUrl string `json:"meet_url"`
}