package constants

import "errors"

var ErrInsertDatabase error = errors.New("invalid insert data in database")
var ErrEmptyInputRegistration error = errors.New("fullname, username, email and password cannot be empty")
var ErrUserNotFound error = errors.New("user not found")
var ErrGetAllDatabase error = errors.New("failed get all data from database")
var ErrHashedPassword error = errors.New("invalid hashed password")