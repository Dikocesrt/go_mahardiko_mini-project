package user

import (
	"habit/constants"
	userEntitites "habit/entities/user"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	repository userEntitites.RepositoryInterface
}

func NewUserUseCase(repository userEntitites.RepositoryInterface) *UserUseCase {
	return &UserUseCase{
		repository: repository,
	}
}

func (userUseCase *UserUseCase) Register(user *userEntitites.User) (userEntitites.User, error) {
	if user.FullName == "" || user.Username == "" || user.Email == "" || user.Password == "" {
		return userEntitites.User{}, constants.ErrEmptyInputRegistration
	}

	id := uuid.New()
	user.Id = id.String()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return userEntitites.User{}, constants.ErrHashedPassword
	}

	user.Password = string(hashedPassword)
	
	newUser, err := userUseCase.repository.Register(user)
	if err != nil {
		return userEntitites.User{}, constants.ErrInsertDatabase
	}

	return newUser, nil
}