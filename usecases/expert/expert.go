package expert

import (
	"habit/constants"
	expertEntities "habit/entities/expert"
	"habit/middlewares"

	"golang.org/x/crypto/bcrypt"
)

type ExpertUseCase struct {
	repository expertEntities.RepositoryInterface
}

func NewExpertUseCase(repository expertEntities.RepositoryInterface) *ExpertUseCase {
	return &ExpertUseCase{
		repository: repository,
	}
}

func (expertUseCase *ExpertUseCase) Register(expert *expertEntities.Expert) (expertEntities.Expert, error) {
	if expert.FullName == "" || expert.Username == "" || expert.Email == "" || expert.Password == "" || expert.Gender == "" || expert.Age == 0 || expert.Fee == 0 || expert.BankAccountTypeId == 0 || expert.BankAccount.AccountName == "" || expert.BankAccount.AccountNumber == "" || expert.ExpertiseId == 0 {
		return expertEntities.Expert{}, constants.ErrEmptyInputRegistration
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(expert.Password), bcrypt.DefaultCost)
	if err != nil {
		return expertEntities.Expert{}, constants.ErrHashedPassword
	}

	expert.Password = string(hashedPassword)

	newExpert, err := expertUseCase.repository.Register(expert)
	if err != nil {
		return expertEntities.Expert{}, constants.ErrInsertDatabase
	}

	return newExpert, nil
}

func (expertUseCase *ExpertUseCase) Login(expert *expertEntities.Expert) (expertEntities.Expert, error) {
	if expert.Username == "" || expert.Password == "" {
		return expertEntities.Expert{}, constants.ErrEmptyInputLogin
	}

	expertDb, err := expertUseCase.repository.Login(expert)
	if err != nil {
		return expertEntities.Expert{}, constants.ErrUserNotFound
	}

	token, _ := middlewares.CreateTokenExpert(expertDb.Id)
	expertDb.Token = token

	return expertDb, nil
}