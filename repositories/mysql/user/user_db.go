package user

import (
	"gorm.io/gorm"

	userEntities "habit/entities/user"
)

type User struct {
	gorm.Model
	Id             int `gorm:"primaryKey:autoIncrement"`
	FullName       string
	Username       string
	Email          string
	Password       string
	Address        string
	Bio            string
	PhoneNumber    string
	Gender         string `gorm:"type:ENUM('pria', 'wanita')"`
	Age            int
	ProfilePicture string
}

func FromUserEntitiesToUserDb(user *userEntities.User) *User {
	return &User{
		Id:        user.Id,
		FullName:  user.FullName,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		Address:   user.Address,
		Bio:       user.Bio,
		PhoneNumber: user.PhoneNumber,
		Gender:    user.Gender,
		Age:       user.Age,
		ProfilePicture: user.ProfilePicture,
	}
}

func (user *User) FromUserDbToUserEntities() *userEntities.User {
	return &userEntities.User{
		Id:        user.Id,
		FullName:  user.FullName,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		Address:   user.Address,
		Bio:       user.Bio,
		PhoneNumber: user.PhoneNumber,
		Gender:    user.Gender,
		Age:       user.Age,
		ProfilePicture: user.ProfilePicture,
	}
}