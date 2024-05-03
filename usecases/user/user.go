package user

import (
	"habit/constants"
	userEntitites "habit/entities/user"
	"habit/middlewares"

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

func (userUseCase *UserUseCase) Login(user *userEntitites.User) (userEntitites.User, error) {
	if user.Username == "" || user.Password == "" {
		return userEntitites.User{}, constants.ErrEmptyInputLogin
	}

	userFromDb, err := userUseCase.repository.Login(user)
	if err != nil {
		return userEntitites.User{}, constants.ErrUserNotFound
	}

	token, _ := middlewares.CreateToken(userFromDb.Id)
	userFromDb.Token = token

	return userFromDb, nil
}