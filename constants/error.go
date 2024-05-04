package constants

import "errors"

var ErrInsertDatabase error = errors.New("invalid insert data in database")
var ErrEmptyInputRegistration error = errors.New("fullname, username, email and password cannot be empty")
var ErrEmptyInputLogin error = errors.New("username and password cannot be empty")
var ErrUserNotFound error = errors.New("user not found")
var ErrGetDatabase error = errors.New("failed get data from database")
var ErrUpdateDatabase error = errors.New("failed update data in database")
var ErrDeleteDatabase error = errors.New("failed delete data in database")
var ErrHashedPassword error = errors.New("invalid hashed password")
var ErrEmptyInputCreateActivity error = errors.New("title, activity start, activity finish and user id cannot be empty")
var ErrEmptyInputUpdateProfile error = errors.New("fullname, username, email and password cannot be empty")
var ErrUsernameAlreadyExist error = errors.New("username already exist")
var ErrEmailAlreadyExist error = errors.New("email already exist")
var ErrCloudinary error = errors.New("cloudinary url not found")