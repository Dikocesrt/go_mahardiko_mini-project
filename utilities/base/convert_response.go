package base

import (
	"habit/constants"
	"net/http"
)

func ConvertResponseCode(err error) int {
	switch err {
	case constants.ErrInsertDatabase:
		return http.StatusInternalServerError
	case constants.ErrEmptyInputRegistration:
		return http.StatusBadRequest
	case constants.ErrUserNotFound:
		return http.StatusNotFound
	case constants.ErrGetDatabase:
		return http.StatusInternalServerError
	case constants.ErrUpdateDatabase:
		return http.StatusInternalServerError
	case constants.ErrDeleteDatabase:
		return http.StatusInternalServerError
	case constants.ErrHashedPassword:
		return http.StatusInternalServerError
	case constants.ErrEmptyInputLogin:
		return http.StatusBadRequest
	case constants.ErrEmptyInputCreateActivity:
		return http.StatusBadRequest
	case constants.ErrEmptyInputUpdateProfile:
		return http.StatusBadRequest
	case constants.ErrUsernameAlreadyExist:
		return http.StatusBadRequest
	case constants.ErrEmailAlreadyExist:
		return http.StatusBadRequest
	case constants.ErrCloudinary:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}