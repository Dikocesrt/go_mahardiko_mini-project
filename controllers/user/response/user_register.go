package response

type UserRegisterResponse struct {
	Id       int    `json:"id"`
	FullName string `json:"fullName"`
	Username string `json:"username"`
	Email    string `json:"email"`
}