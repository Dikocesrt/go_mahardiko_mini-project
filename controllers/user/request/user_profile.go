package request

type UserProfileRequest struct {
	FullName       string `json:"fullname" form:"fullname"`
	Username       string `json:"username" form:"username"`
	Email          string `json:"email" form:"email"`
	Password       string `json:"password" form:"password"`
	Address        string `json:"address" form:"address"`
	Bio            string `json:"bio" form:"bio"`
	PhoneNumber    string `json:"phone_number" form:"phone_number"`
	Gender         string `json:"gender" form:"gender"`
	Age            int    `json:"age" form:"age"`
	ProfilePicture string `json:"profile_picture" form:"profile_picture"`
}