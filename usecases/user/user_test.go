package user

import (
	"errors"
	"habit/constants"
	userEntities "habit/entities/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockUserRepository struct {
	Err error
	Kode int64
}

func (m *MockUserRepository) Register(user *userEntities.User) (userEntities.User, int64, error) {
	return *user, m.Kode, m.Err
}

func (m *MockUserRepository) Login(user *userEntities.User) (userEntities.User, error) {
	return *user, m.Err
}

func (m *MockUserRepository) UpdateProfileById(user *userEntities.User) (userEntities.User, int64, error) {
	return *user, m.Kode, m.Err
}

func (m *MockUserRepository) GetUserById(user *userEntities.User) (userEntities.User, error) {
	return *user, m.Err
}

func TestRegister(t *testing.T) {
	t.Run("Register Success", func(t *testing.T) {
		mockUserRepository := &MockUserRepository{Err: nil, Kode: 0}

		userUseCase := NewUserUseCase(mockUserRepository)

		user := userEntities.User{
			FullName: "John Doe",
			Username: "johndoe",
			Email: "johndoe@ex.com",
			Password: "password123",
		}

		registeredUser, err := userUseCase.Register(&user)

		assert.NoError(t, err)
		assert.Equal(t, user, registeredUser)
	})

	t.Run("Register Empty Input", func(t *testing.T) {
		mockUserRepository := &MockUserRepository{Err: nil, Kode: 0}

		userUseCase := NewUserUseCase(mockUserRepository)

		user := userEntities.User{}

		registeredUser, err := userUseCase.Register(&user)

		assert.Error(t, err)
		assert.Equal(t, constants.ErrEmptyInputUser, err)
		assert.Equal(t, userEntities.User{}, registeredUser)
	})

	t.Run("Register Database Error", func(t *testing.T) {
		mockUserRepository := &MockUserRepository{Err: errors.New("database error"), Kode: 0}

		userUseCase := NewUserUseCase(mockUserRepository)

		user := userEntities.User{
			FullName: "John Doe",
			Username: "johndoe",
			Email: "johndoe@ex.com",
			Password: "password123",
		}

		registeredUser, err := userUseCase.Register(&user)
		assert.Error(t, err)
		assert.Equal(t, constants.ErrInsertDatabase, err)
		assert.Equal(t, userEntities.User{}, registeredUser)
	})

	t.Run("Register Username Already Exist", func(t *testing.T) {
		mockUserRepository := &MockUserRepository{Err: nil, Kode: 1}

		userUseCase := NewUserUseCase(mockUserRepository)

		user := userEntities.User{
			FullName: "John Doe",
			Username: "johndoe",
			Email: "johndoe@ex.com",
			Password: "password123",
		}

		registeredUser, err := userUseCase.Register(&user)
		assert.Error(t, err)
		assert.Equal(t, constants.ErrUsernameAlreadyExist, err)
		assert.Equal(t, userEntities.User{}, registeredUser)
	})

	t.Run("Register Email Already Exist", func(t *testing.T) {
		mockUserRepository := &MockUserRepository{Err: nil, Kode: 2}

		userUseCase := NewUserUseCase(mockUserRepository)

		user := userEntities.User{
			FullName: "John Doe",
			Username: "johndoe",
			Email: "johndoe@ex.com",
			Password: "password123",
		}

		registeredUser, err := userUseCase.Register(&user)
		assert.Error(t, err)
		assert.Equal(t, constants.ErrEmailAlreadyExist, err)
		assert.Equal(t, userEntities.User{}, registeredUser)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Login Success", func(t *testing.T) {
		mockUserRepository := &MockUserRepository{Err: nil}

		userUseCase := NewUserUseCase(mockUserRepository)

		user := userEntities.User{
			Username: "johndoe",
			Password: "password123",
		}

		loginUser, err := userUseCase.Login(&user)

		assert.NoError(t, err)
		assert.Equal(t, user.Username, loginUser.Username)
	})

	t.Run("Login Empty Input", func(t *testing.T) {
		mockUserRepository := &MockUserRepository{Err: nil}

		userUseCase := NewUserUseCase(mockUserRepository)

		user := userEntities.User{}

		loginUser, err := userUseCase.Login(&user)

		assert.Error(t, err)
		assert.Equal(t, constants.ErrEmptyInputLogin, err)
		assert.Equal(t, userEntities.User{}, loginUser)
	})

	t.Run("Login Database Error", func(t *testing.T) {
		mockUserRepository := &MockUserRepository{Err: errors.New("database error")}

		userUseCase := NewUserUseCase(mockUserRepository)

		user := userEntities.User{
			Username: "johndoe",
			Password: "password123",
		}

		loginUser, err := userUseCase.Login(&user)

		assert.Error(t, err)
		assert.Equal(t, constants.ErrUserNotFound, err)
		assert.Equal(t, userEntities.User{}, loginUser)
	})
}

func TestUpdateProfileById(t *testing.T){

	t.Run("Update Profile Success", func(t *testing.T) {
		mockUserRepository := &MockUserRepository{Err: nil}

		userUseCase := NewUserUseCase(mockUserRepository)

		user := userEntities.User{
			FullName: "John Doe",
			Username: "johndoe",
			Email: "johndoe@ex.com",
			Password: "password123",
		}

		updatedUser, err := userUseCase.UpdateProfileById(&user, nil)

		assert.NoError(t, err)
		assert.Equal(t, user, updatedUser)
	})

	t.Run("Update Profile Empty Input", func(t *testing.T) {
		mockUserRepository := &MockUserRepository{Err: nil}

		userUseCase := NewUserUseCase(mockUserRepository)

		user := userEntities.User{}

		updatedUser, err := userUseCase.UpdateProfileById(&user, nil)

		assert.Error(t, err)
		assert.Equal(t, userEntities.User{}, updatedUser)
	})

	t.Run("Update Profile Database Error", func(t *testing.T) {
		mockUserRepository := &MockUserRepository{Err: errors.New("database error")}

		userUseCase := NewUserUseCase(mockUserRepository)

		user := userEntities.User{
			FullName: "John Doe",
			Username: "johndoe",
			Email: "johndoe@ex.com",
			Password: "password123",
		}

		updatedUser, err := userUseCase.UpdateProfileById(&user, nil)

		assert.Error(t, err)
		assert.Equal(t, userEntities.User{}, updatedUser)
	})

	t.Run("Update Profile Usernamer Already Exist", func(t *testing.T) {
		mockUserRepository := &MockUserRepository{Err: nil, Kode: 2}

		userUseCase := NewUserUseCase(mockUserRepository)
	
		user := userEntities.User{
			FullName: "John Doe",
			Username: "johndoe",
			Email: "johndoe@ex.com",
			Password: "password123",
		}
	
		updatedUser, err := userUseCase.UpdateProfileById(&user, nil)
	
		assert.Error(t, err)
		assert.Equal(t, constants.ErrUsernameAlreadyExist, err)
		assert.Equal(t, userEntities.User{}, updatedUser)
	})

	t.Run("Update Profile Email Already Exist", func(t *testing.T) {
		mockUserRepository := &MockUserRepository{Err: nil, Kode: 3}

		userUseCase := NewUserUseCase(mockUserRepository)
	
		user := userEntities.User{
			FullName: "John Doe",
			Username: "johndoe",
			Email: "johndoe@ex.com",
			Password: "password123",
		}
	
		updatedUser, err := userUseCase.UpdateProfileById(&user, nil)
	
		assert.Error(t, err)
		assert.Equal(t, constants.ErrEmailAlreadyExist, err)
		assert.Equal(t, userEntities.User{}, updatedUser)
	})
}


func TestGetUserById(t *testing.T) {

	t.Run("Get User Success", func(t *testing.T) {
		mockUserRepository := &MockUserRepository{Err: nil}

		userUseCase := NewUserUseCase(mockUserRepository)

		user := userEntities.User{
			Id: 1,
		}

		getUser, err := userUseCase.GetUserById(&user)

		assert.NoError(t, err)
		assert.Equal(t, user, getUser)
	})

	t.Run("Get User Database Error", func(t *testing.T) {
		mockUserRepository := &MockUserRepository{Err: errors.New("database error")}

		userUseCase := NewUserUseCase(mockUserRepository)

		user := userEntities.User{
			Id: 1,
		}

		getUser, err := userUseCase.GetUserById(&user)

		assert.Error(t, err)
		assert.Equal(t, userEntities.User{}, getUser)
	})
}