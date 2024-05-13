package constants

import "errors"

var ErrEmptyInputUser error = errors.New("fullname, username, email or password cannot be empty")

var ErrHashedPassword error = errors.New("error hashing password")

var ErrInsertDatabase error = errors.New("failed insert data in database")

var ErrEmptyInputLogin error = errors.New("username or password cannot be empty")

var ErrUserNotFound error = errors.New("user not found")

var ErrUploadImage error = errors.New("failed upload image")

var ErrUsernameAlreadyExist error = errors.New("username already exist")

var ErrEmailAlreadyExist error = errors.New("email already exist")

var ErrEmptyInputActivity error = errors.New("title, activity start or activity finish cannot be empty")

var ErrGetActivitiesByUserId error = errors.New("no activity data found for the user with the specified id")

var ErrActivityNotFound error = errors.New("activity not found")

var ErrEmptyInputActivityType error = errors.New("name or description cannot be empty")

var ErrGetAllData error = errors.New("failed get data from database")

var ErrUpdateData error = errors.New("failed update data in database")

var ErrDeleteData error = errors.New("failed delete data in database")

var ErrActivityTypeNotFound error = errors.New("activity type not found")

var ErrEmptyInputExpert error = errors.New("fullname, username, email, password, gender, age, fee, bank account type id, bank account name, bank account number or expertise id cannot be empty")

var ErrExpertNotFound error = errors.New("expert not found")

var ErrExpertiseNotFound error = errors.New("expertise not found")

var ErrEmptyInputBankAccountType error = errors.New("name cannot be empty")

var ErrBankAccountTypeNotFound error = errors.New("bank account type not found")

var ErrEmptyInputHire error = errors.New("meet day, meet time, total fee, payment status, payment image, user id or expert id cannot be empty")

var ErrGetHiresByExpertId error = errors.New("no hire data found for the expert with the specified id")

var ErrGetHiresByUserId error = errors.New("no hire data found for the user with the specified id")

var ErrHireNotFound error = errors.New("hire not found")

var ErrEmptyInputVerifyPayment error = errors.New("payment status or meet url cannot be empty")

var ErrEmptyInputAdmin error = errors.New("username, email or password cannot be empty")

var ErrAdminNotFound error = errors.New("admin not found")

var ErrCloudinary error = errors.New("cloudinary url not found")