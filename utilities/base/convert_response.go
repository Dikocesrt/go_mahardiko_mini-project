package base

import (
	"habit/constants"
	"net/http"
)

func ConvertResponseCode(err error) int {
	switch err {
		case constants.ErrEmptyInputUser:
			return http.StatusBadRequest

		case constants.ErrHashedPassword:
			return http.StatusInternalServerError

		case constants.ErrInsertDatabase:
			return http.StatusInternalServerError

		case constants.ErrEmptyInputLogin:
			return http.StatusBadRequest

		case constants.ErrUserNotFound:
			return http.StatusNotFound

		case constants.ErrUploadImage:
			return http.StatusInternalServerError

		case constants.ErrActivityNotFound:
			return http.StatusNotFound

		case constants.ErrUsernameAlreadyExist:
			return http.StatusBadRequest

		case constants.ErrEmailAlreadyExist:
			return http.StatusBadRequest

		case constants.ErrEmptyInputActivity:
			return http.StatusBadRequest

		case constants.ErrGetActivitiesByUserId:
			return http.StatusNotFound

		case constants.ErrEmptyInputActivityType:
			return http.StatusBadRequest

		case constants.ErrGetAllData:
			return http.StatusInternalServerError

		case constants.ErrActivityTypeNotFound:
			return http.StatusNotFound

		case constants.ErrEmptyInputExpert:
			return http.StatusBadRequest

		case constants.ErrExpertNotFound:
			return http.StatusNotFound

		case constants.ErrExpertiseNotFound:
			return http.StatusNotFound

		case constants.ErrEmptyInputBankAccountType:
			return http.StatusBadRequest

		case constants.ErrBankAccountTypeNotFound:
			return http.StatusNotFound

		case constants.ErrEmptyInputHire:
			return http.StatusBadRequest

		case constants.ErrGetHiresByExpertId:
			return http.StatusNotFound

		case constants.ErrGetHiresByUserId:
			return http.StatusNotFound

		case constants.ErrHireNotFound:
			return http.StatusNotFound

		case constants.ErrEmptyInputVerifyPayment:
			return http.StatusBadRequest

		case constants.ErrEmptyInputAdmin:
			return http.StatusBadRequest

		case constants.ErrAdminNotFound:
			return http.StatusNotFound

		case constants.ErrCloudinary:
			return http.StatusInternalServerError

		case constants.ErrUpdateData:
			return http.StatusInternalServerError

		case constants.ErrDeleteData:
			return http.StatusInternalServerError

		default:
			return http.StatusInternalServerError
	}
}