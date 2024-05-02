package user

import (
	userEntities "habit/entities/user"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		DB: db,
	}
}

func (userRepo *UserRepo) Register(user *userEntities.User) (userEntities.User, error) {
	userDb := FromUserEntitiesToUserDb(user)

	err := userRepo.DB.Create(&userDb).Error

	if err != nil {
		return userEntities.User{}, err
	}

	newUser := userDb.FromUserDbToUserEntities()

	return *newUser, nil
}

func (userRepo *UserRepo) Login(user *userEntities.User) (userEntities.User, error) {
	userDb := FromUserEntitiesToUserDb(user)

	password := userDb.Password

	err := userRepo.DB.Where("Username = ?", userDb.Username).First(&userDb).Error
	if err != nil {
		err := userRepo.DB.Where("Email = ?", userDb.Username).First(&userDb).Error
		if err != nil {
			return userEntities.User{}, err
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDb.Password), []byte(password))
	if err != nil {
		return userEntities.User{}, err
	}

	userFromDb := userDb.FromUserDbToUserEntities()

	return *userFromDb, nil
}