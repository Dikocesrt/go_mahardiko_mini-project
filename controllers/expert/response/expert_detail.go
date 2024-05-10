package response

type ExpertDetailResponse struct {
	Id             int                        `json:"id"`
	Username       string                     `json:"username"`
	Email          string                     `json:"email"`
	FullName       string                     `json:"fullName"`
	Address        string                     `json:"address"`
	Bio            string                     `json:"bio"`
	PhoneNumber    string                     `json:"phoneNumber"`
	Gender         string                     `json:"gender"`
	Age            int                        `json:"age"`
	ProfilePicture string                     `json:"profilePicture"`
	Experience     int                        `json:"experience"`
	Fee            int                        `json:"fee"`
	BankAccount    BankAccountProfileResponse `json:"bankAccount"`
	Expertise      ExpertiseDetailResponse    `json:"expertise"`
}

type BankAccountDetailResponse struct {
	AccountName   string `json:"accountName"`
	AccountNumber string `json:"accountNumber"`
}

type ExpertiseDetailResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}