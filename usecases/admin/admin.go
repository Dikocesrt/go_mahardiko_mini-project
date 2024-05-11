package admin

import (
	"habit/constants"
	adminEntities "habit/entities/admin"
	"habit/entities/expert"
	"habit/middlewares"

	"golang.org/x/crypto/bcrypt"
)

type AdminUseCase struct {
	repository adminEntities.RepositoryInterface
}

func NewAdminUseCase(repository adminEntities.RepositoryInterface) *AdminUseCase {
	return &AdminUseCase{
		repository: repository,
	}
}

func (adminUseCase *AdminUseCase) Register(admin *adminEntities.Admin) (adminEntities.Admin, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		return adminEntities.Admin{}, constants.ErrHashedPassword
	}

	admin.Password = string(hashedPassword)

	registeredAdmin, err := adminUseCase.repository.Register(admin)
	if err != nil {
		return adminEntities.Admin{}, err
	}
	return registeredAdmin, nil
}

func (adminUseCase *AdminUseCase) Login(admin *adminEntities.Admin) (adminEntities.Admin, error) {
	loginAdmin, err := adminUseCase.repository.Login(admin)
	if err != nil {
		return adminEntities.Admin{}, err
	}

	token, _ := middlewares.CreateTokenAdmin(loginAdmin.Id)

	loginAdmin.Token = token

	return loginAdmin, nil
}


func (adminUseCase *AdminUseCase) CreateBankAccountType(bankType expert.BankAccountType) (expert.BankAccountType, error) {
	bankType, err := adminUseCase.repository.CreateBankAccountType(bankType)
	if err != nil {
		return expert.BankAccountType{}, err
	}
	return bankType, nil
}

func (adminUseCase *AdminUseCase) GetBankAccountTypeById(bankType expert.BankAccountType) (expert.BankAccountType, error) {
	bankType, err := adminUseCase.repository.GetBankAccountTypeById(bankType)
	if err != nil {
		return expert.BankAccountType{}, err
	}
	return bankType, nil
}

func (adminUseCase *AdminUseCase) GetAllBankAccountType() ([]expert.BankAccountType, error) {
	bankTypes, err := adminUseCase.repository.GetAllBankAccountType()
	if err != nil {
		return []expert.BankAccountType{}, err
	}
	return bankTypes, nil
}

func (adminUseCase *AdminUseCase) UpdateBankAccountTypeById(bankType expert.BankAccountType) (expert.BankAccountType, error) {
	bankType, err := adminUseCase.repository.UpdateBankAccountTypeById(bankType)
	if err != nil {
		return expert.BankAccountType{}, err
	}
	return bankType, nil
}

func (adminUseCase *AdminUseCase) DeleteBankAccountTypeById(bankType expert.BankAccountType) error {
	err := adminUseCase.repository.DeleteBankAccountTypeById(bankType)
	if err != nil {
		return err
	}
	return nil
}


func (adminUseCase *AdminUseCase) CreateExpertise(expertise expert.Expertise) (expert.Expertise, error) {
	newExpertise, err := adminUseCase.repository.CreateExpertise(expertise)
	if err != nil {
		return expert.Expertise{}, err
	}
	return newExpertise, nil
}

func (adminUseCase *AdminUseCase) GetAllExpertise() ([]expert.Expertise, error) {
	newExpertise, err := adminUseCase.repository.GetAllExpertise()
	if err != nil {
		return []expert.Expertise{}, err
	}
	return newExpertise, nil
}

func (adminUseCase *AdminUseCase) GetExpertiseById(expertise expert.Expertise) (expert.Expertise, error) {
	newExpertise, err := adminUseCase.repository.GetExpertiseById(expertise)
	if err != nil {
		return expert.Expertise{}, err
	}
	return newExpertise, nil
}

func (adminUseCase *AdminUseCase) UpdateExpertiseById(expertise expert.Expertise) (expert.Expertise, error) {
	newExpertise, err := adminUseCase.repository.UpdateExpertiseById(expertise)
	if err != nil {
		return expert.Expertise{}, err
	}
	return newExpertise, nil
}

func (adminUseCase *AdminUseCase) DeleteExpertiseById(expertise expert.Expertise) error {
	err := adminUseCase.repository.DeleteExpertiseById(expertise)
	if err != nil {
		return err
	}
	return nil
}