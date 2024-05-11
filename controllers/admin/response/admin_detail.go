package response

type DetailAdmin struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}