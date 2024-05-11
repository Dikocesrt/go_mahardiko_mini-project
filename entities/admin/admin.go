package admin

import "habit/entities/expert"

type Admin struct {
	Id       int
	Username string
	Email    string
	Password string
	Token    string
}

type RepositoryInterface interface {
	Register(admin *Admin) (Admin, error)
	Login(admin *Admin) (Admin, error)
	CreateBankAccountType(bankType expert.BankAccountType) (expert.BankAccountType, error)
	GetBankAccountTypeById(bankType expert.BankAccountType) (expert.BankAccountType, error)
	GetAllBankAccountType() ([]expert.BankAccountType, error)
	UpdateBankAccountTypeById(bankType expert.BankAccountType) (expert.BankAccountType, error)
	DeleteBankAccountTypeById(bankType expert.BankAccountType) (error)
	CreateExpertise(expertise expert.Expertise) (expert.Expertise, error)
	GetExpertiseById(expertise expert.Expertise) (expert.Expertise, error)
	GetAllExpertise() ([]expert.Expertise, error)
	UpdateExpertiseById(expertise expert.Expertise) (expert.Expertise, error)
	DeleteExpertiseById(expertise expert.Expertise) (error)
}

type UseCaseInterface interface {
	Register(admin *Admin) (Admin, error)
	Login(admin *Admin) (Admin, error)
	CreateBankAccountType(bankType expert.BankAccountType) (expert.BankAccountType, error)
	GetBankAccountTypeById(bankType expert.BankAccountType) (expert.BankAccountType, error)
	GetAllBankAccountType() ([]expert.BankAccountType, error)
	UpdateBankAccountTypeById(bankType expert.BankAccountType) (expert.BankAccountType, error)
	DeleteBankAccountTypeById(bankType expert.BankAccountType) (error)
	CreateExpertise(expertise expert.Expertise) (expert.Expertise, error)
	GetExpertiseById(expertise expert.Expertise) (expert.Expertise, error)
	GetAllExpertise() ([]expert.Expertise, error)
	UpdateExpertiseById(expertise expert.Expertise) (expert.Expertise, error)
	DeleteExpertiseById(expertise expert.Expertise) (error)
}