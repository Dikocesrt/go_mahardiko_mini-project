package expert

import "mime/multipart"

type Expert struct {
	Id                int
	FullName          string
	Username          string
	Email             string
	Password          string
	Address           string
	Bio               string
	PhoneNumber       string
	Gender            string
	Age               int
	ProfilePicture    string
	Experience        int
	CustomerCount     int
	Fee               int
	BankAccountTypeId int
	BankAccountType   BankAccountType
	BankAccountId     int
	BankAccount       BankAccount
	ExpertiseId       int
	Expertise         Expertise
	Token             string
}

type BankAccount struct {
	Id                int
	AccountNumber     string
	AccountName       string
	BankAccountTypeId int
	BankAccountType   BankAccountType
}

type BankAccountType struct {
	Id   int
	Name string
}

type Expertise struct {
	Id          int
	Name        string
	Description string
}

type RepositoryInterface interface {
	Register(expert *Expert) (Expert, error)
	Login(expert *Expert) (Expert, error)
	UpdateProfileExpertById(expert *Expert) (Expert, int64, error)
	GetAllExperts() ([]Expert, error)
}

type UseCaseInterface interface {
	Register(expert *Expert) (Expert, error)
	Login(expert *Expert) (Expert, error)
	UpdateProfileExpertById(expert *Expert, file *multipart.FileHeader) (Expert, error)
	GetAllExperts() ([]Expert, error)
}