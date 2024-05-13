package expert

import (
	expertEntities "habit/entities/expert"

	"gorm.io/gorm"
)

type Expert struct {
	gorm.Model
	Id int `gorm:"primaryKey;autoIncrement"`
	FullName string
	Username string
	Email string
	Password string
	Address string
	Bio string
	PhoneNumber string
	Gender string `gorm:"type:ENUM('pria', 'wanita')"`
	Age int
	ProfilePicture string
	Experience int
	CustomerCount int
	Fee int
	BankAccountId int `gorm:"index"`
	ExpertiseId int `gorm:"index"`
}

type BankAccount struct {
	gorm.Model
	Id int `gorm:"primaryKey;autoIncrement"`
	AccountNumber string
	AccountName string
	BankAccountTypeId int `gorm:"index"`
}

type BankAccountType struct {
	gorm.Model
	Id int `gorm:"primaryKey;autoIncrement"`
	Name string
}

type Expertise struct {
	gorm.Model
	Id int `gorm:"primaryKey;autoIncrement"`
	Name string
	Description string
}

func FromExpertEntitiesToExpertDb(expert *expertEntities.Expert) *Expert {
	return &Expert{
		Id: expert.Id,
		FullName: expert.FullName,
		Username: expert.Username,
		Email: expert.Email,
		Password: expert.Password,
		Address: expert.Address,
		Bio: expert.Bio,
		PhoneNumber: expert.PhoneNumber,
		Gender: expert.Gender,
		Age: expert.Age,
		ProfilePicture: expert.ProfilePicture,
		Experience: expert.Experience,
		CustomerCount: expert.CustomerCount,
		Fee: expert.Fee,
		BankAccountId: expert.BankAccountId,
		ExpertiseId: expert.ExpertiseId,
	}
}

func (expert *Expert) FromExpertDbToExpertEntities() *expertEntities.Expert {
	return &expertEntities.Expert{
		Id: expert.Id,
		FullName: expert.FullName,
		Username: expert.Username,
		Email: expert.Email,
		Password: expert.Password,
		Address: expert.Address,
		Bio: expert.Bio,
		PhoneNumber: expert.PhoneNumber,
		Gender: expert.Gender,
		Age: expert.Age,
		ProfilePicture: expert.ProfilePicture,
		Experience: expert.Experience,
		CustomerCount: expert.CustomerCount,
		Fee: expert.Fee,
		BankAccountId: expert.BankAccountId,
		ExpertiseId: expert.ExpertiseId,
	}
}

func FromBankAccountEntitiesToBankAccountDb(expert *expertEntities.Expert) *BankAccount {
	return &BankAccount{
		Id: expert.BankAccount.Id,
		AccountNumber: expert.BankAccount.AccountNumber,
		AccountName: expert.BankAccount.AccountName,
		BankAccountTypeId: expert.BankAccount.BankAccountTypeId,
	}
}

func (bankAccount *BankAccount) FromBankAccountDbToBankAccountEntities() *expertEntities.BankAccount {
	return &expertEntities.BankAccount{
		Id: bankAccount.Id,
		AccountNumber: bankAccount.AccountNumber,
		AccountName: bankAccount.AccountName,
		BankAccountTypeId: bankAccount.BankAccountTypeId,
	}
}

func (expertise *Expertise) FromExpertiseDbToExpertiseEntities() *expertEntities.Expertise {
	return &expertEntities.Expertise{
		Id: expertise.Id,
		Name: expertise.Name,
		Description: expertise.Description,
	}
}