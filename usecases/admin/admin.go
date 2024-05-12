package admin

import (
	"habit/constants"
	adminEntities "habit/entities/admin"
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
	if admin.Username == "" || admin.Email == "" || admin.Password == "" {
		return adminEntities.Admin{}, constants.ErrEmptyInputAdmin
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		return adminEntities.Admin{}, constants.ErrHashedPassword
	}

	admin.Password = string(hashedPassword)

	registeredAdmin, err := adminUseCase.repository.Register(admin)
	if err != nil {
		return adminEntities.Admin{}, constants.ErrInsertDatabase
	}
	return registeredAdmin, nil
}

func (adminUseCase *AdminUseCase) Login(admin *adminEntities.Admin) (adminEntities.Admin, error) {
	if admin.Username == "" || admin.Password == "" {
		return adminEntities.Admin{}, constants.ErrEmptyInputLogin
	}
	loginAdmin, err := adminUseCase.repository.Login(admin)
	if err != nil {
		return adminEntities.Admin{}, constants.ErrAdminNotFound
	}

	token, _ := middlewares.CreateTokenAdmin(loginAdmin.Id)

	loginAdmin.Token = token

	return loginAdmin, nil
}