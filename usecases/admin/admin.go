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