package admin

import (
	adminEntities "habit/entities/admin"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AdminRepo struct {
	DB *gorm.DB
}

func NewAdminRepo(db *gorm.DB) *AdminRepo {
	return &AdminRepo{
		DB: db,
	}
}

func (adminRepo *AdminRepo) Register(admin *adminEntities.Admin) (adminEntities.Admin, error) {
	var adminDb Admin
	adminDb.Username = admin.Username
	adminDb.Email = admin.Email
	adminDb.Password = admin.Password
	err := adminRepo.DB.Create(&adminDb).Error
	if err != nil {
		return adminEntities.Admin{}, err
	}
	admin.Id = adminDb.Id
	return *admin, nil
}

func (adminRepo *AdminRepo) Login(admin *adminEntities.Admin) (adminEntities.Admin, error) {
	var adminDb Admin
	adminDb.Username = admin.Username
	adminDb.Password = admin.Password

	err := adminRepo.DB.Where("Username = ?", adminDb.Username).First(&adminDb).Error
	if err != nil {
		err := adminRepo.DB.Where("Email = ?", adminDb.Username).First(&adminDb).Error
		if err != nil {
			return adminEntities.Admin{}, err
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(adminDb.Password), []byte(admin.Password))
	if err != nil {
		return adminEntities.Admin{}, err
	}
	admin.Id = adminDb.Id
	admin.Username = adminDb.Username
	admin.Email = adminDb.Email
	return *admin, nil
}