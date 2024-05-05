package request

type ExpertRegisterRequest struct {
	FullName          string `json:"fullname" form:"fullname"`
	Username          string `json:"username" form:"username"`
	Email             string `json:"email" form:"email"`
	Password          string `json:"password" form:"password"`
	Gender            string `json:"gender" form:"gender"`
	Age               int    `json:"age" form:"age"`
	Experience        int    `json:"experience" form:"experience"`
	Fee               int    `json:"fee" form:"fee"`
	BankAccountTypeId int    `json:"bank_account_type_id" form:"bank_account_type_id"`
	AccountNumber     string `json:"account_number" form:"account_number"`
	AccountName       string `json:"account_name" form:"account_name"`
	ExpertiseId       int    `json:"expertise_id" form:"expertise_id"`
}