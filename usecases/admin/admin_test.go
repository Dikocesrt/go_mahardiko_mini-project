package admin

import (
	"errors"
	adminEntities "habit/entities/admin"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockAdminRepository struct {
	Err  error
}

func (m *MockAdminRepository) Register(admin *adminEntities.Admin) (adminEntities.Admin, error) {
	return *admin, m.Err
}

func (m *MockAdminRepository) Login(admin *adminEntities.Admin) (adminEntities.Admin, error) {
	return *admin, m.Err
}

func TestRegister(t *testing.T) {
	t.Run("Register Success", func(t *testing.T) {
		mockAdminRepository := &MockAdminRepository{Err: nil}

		adminUseCase := NewAdminUseCase(mockAdminRepository)

		admin := adminEntities.Admin{
			Username: "admin",
			Email:    "admin@example.com",
			Password: "password123",
		}

		result, err := adminUseCase.Register(&admin)

		assert.NoError(t, err)
		assert.Equal(t, admin, result)
	})

	t.Run("Register Empty Input", func(t *testing.T) {
		mockAdminRepository := &MockAdminRepository{Err: nil}

		adminUseCase := NewAdminUseCase(mockAdminRepository)

		admin := adminEntities.Admin{}

		result, err := adminUseCase.Register(&admin)

		assert.Error(t, err)
		assert.Equal(t, adminEntities.Admin{}, result)
	})

	t.Run("Register Database Insert Error", func(t *testing.T) {
		mockAdminRepository := &MockAdminRepository{Err: errors.New("database insert failed")}

		adminUseCase := NewAdminUseCase(mockAdminRepository)

		admin := adminEntities.Admin{
			Username: "admin",
			Email:    "admin@example.com",
			Password: "password123",
		}

		result, err := adminUseCase.Register(&admin)

		assert.Error(t, err)
		assert.Equal(t, adminEntities.Admin{}, result)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Login Success", func(t *testing.T) {
		mockAdminRepository := &MockAdminRepository{Err: nil}

		adminUseCase := NewAdminUseCase(mockAdminRepository)

		admin := adminEntities.Admin{
			Username: "admin",
			Password: "password123",
		}

		result, err := adminUseCase.Login(&admin)

		assert.NoError(t, err)
		assert.Equal(t, admin.Username, result.Username)
	})

	t.Run("Login Empty Input", func(t *testing.T) {
		mockAdminRepository := &MockAdminRepository{Err: nil}

		adminUseCase := NewAdminUseCase(mockAdminRepository)

		admin := adminEntities.Admin{}

		result, err := adminUseCase.Login(&admin)

		assert.Error(t, err)
		assert.Equal(t, adminEntities.Admin{}, result)
	})

	t.Run("Login Admin Not Found", func(t *testing.T) {
		mockAdminRepository := &MockAdminRepository{Err: errors.New("admin not found")}

		adminUseCase := NewAdminUseCase(mockAdminRepository)

		admin := adminEntities.Admin{
			Username: "admin",
			Password: "password123",
		}

		result, err := adminUseCase.Login(&admin)

		assert.Error(t, err)
		assert.Equal(t, adminEntities.Admin{}, result)
	})
}