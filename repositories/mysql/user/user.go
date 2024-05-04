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

func (userRepo *UserRepo) UpdateProfileById(user *userEntities.User) (userEntities.User, int64, error ) {
	userDb := FromUserEntitiesToUserDb(user)

	var userDbTemp User
	err := userRepo.DB.Where("Id = ?", userDb.Id).First(&userDbTemp).Error
	if err != nil {
		return userEntities.User{}, 1, err
	}

	if userDb.ProfilePicture == "" {
		userDb.ProfilePicture = userDbTemp.ProfilePicture
	}

	var counterUsername, counterEmail int64
	err = userRepo.DB.Model(&userDb).Where("Username = ?", userDb.Username).Count(&counterUsername).Error
	if err != nil {
		return userEntities.User{}, 1, err
	}

	if userDb.Username != userDbTemp.Username && counterUsername > 0 {
		return userEntities.User{}, 2, err
	}

	if userDb.Username == userDbTemp.Username && counterUsername > 1 {
		return userEntities.User{}, 2, err
	}
	
	err = userRepo.DB.Model(&userDb).Where("Email = ?", userDb.Email).Count(&counterEmail).Error
	if err != nil {
		return userEntities.User{}, 1, err
	}

	if userDb.Email != userDbTemp.Email && counterEmail > 0 {
		return userEntities.User{}, 3, err
	}

	if userDb.Email == userDbTemp.Email && counterEmail > 1 {
		return userEntities.User{}, 3, err
	}

	err = userRepo.DB.Save(&userDb).Error
	if err != nil {
		return userEntities.User{}, 1, err
	}

	err = userRepo.DB.Where("Id = ?", userDb.Id).First(&userDb).Error
	if err != nil {
		return userEntities.User{}, 1, err
	}

	newUser := userDb.FromUserDbToUserEntities()

	return *newUser, 0, nil
}