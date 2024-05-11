package response

type UserProfileResponse struct {
	Id             int    `json:"id"`
	FullName       string `json:"fullname"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	Address        string `json:"address"`
	Bio            string `json:"bio"`
	PhoneNumber    string `json:"phone_number"`
	Gender         string `json:"gender"`
	Age            int    `json:"age"`
	ProfilePicture string `json:"profile_picture"`
}