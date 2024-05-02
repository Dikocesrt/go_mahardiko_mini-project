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
	case constants.ErrGetAllDatabase:
		return http.StatusInternalServerError
	case constants.ErrHashedPassword:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}