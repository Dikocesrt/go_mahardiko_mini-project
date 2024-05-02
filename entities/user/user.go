package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id             string
	FullName       string
	Username       string
	Email          string
	Password       string
	Address        string
	Bio            string
	PhoneNumber    string
	Gender         string
	Age            int
	ProfilePicture string
	Token          string
}

type RepositoryInterface interface {
	Register(user *User) (User, error)
}

type UseCaseInterface interface {
	Register(user *User) (User, error)
}