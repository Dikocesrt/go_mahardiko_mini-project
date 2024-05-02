package request

type UserRegisterRequest struct {
	FullName string `json:"fullname"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}