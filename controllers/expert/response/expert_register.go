package response

type ExpertRegisterResponse struct {
	Id          int                 `json:"id"`
	Username    string              `json:"username"`
	Email       string              `json:"email"`
	FullName    string              `json:"fullName"`
	Gender      string              `json:"gender"`
	Age         int                 `json:"age"`
	Experience  int                 `json:"experience"`
	Fee         int                 `json:"fee"`
	BankAccount BankAccountResponse `json:"bankAccount"`
	Expertise   ExpertiseResponse   `json:"expertise"`
}

type BankAccountResponse struct {
	AccountName   string `json:"accountName"`
	AccountNumber string `json:"accountNumber"`
}

type ExpertiseResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}