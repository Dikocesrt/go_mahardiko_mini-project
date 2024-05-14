package expert

import (
	"errors"
	expertEntities "habit/entities/expert"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockExpertRepository struct {
	Err  error
	Kode int64
}

func (m *MockExpertRepository) Register(expert *expertEntities.Expert) (expertEntities.Expert, error) {
	return *expert, m.Err
}

func (m *MockExpertRepository) Login(expert *expertEntities.Expert) (expertEntities.Expert, error) {
	return *expert, m.Err
}

func (m *MockExpertRepository) UpdateProfileExpertById(expert *expertEntities.Expert) (expertEntities.Expert, int64, error) {
	return *expert, m.Kode, m.Err
}

func (m *MockExpertRepository) GetAllExperts() ([]expertEntities.Expert, error) {
	return []expertEntities.Expert{}, m.Err
}

func (m *MockExpertRepository) GetExpertById(expert *expertEntities.Expert) (expertEntities.Expert, error) {
	return *expert, m.Err
}

func (m *MockExpertRepository) CreateExpertise(expertise expertEntities.Expertise) (expertEntities.Expertise, error) {
	return expertise, m.Err
}

func (m *MockExpertRepository) GetExpertiseById(expertise expertEntities.Expertise) (expertEntities.Expertise, error) {
	return expertise, m.Err
}

func (m *MockExpertRepository) GetAllExpertise() ([]expertEntities.Expertise, error) {
	return []expertEntities.Expertise{}, m.Err
}

func (m *MockExpertRepository) UpdateExpertiseById(expertise expertEntities.Expertise) (expertEntities.Expertise, int64, error) {
	return expertise, m.Kode, m.Err
}

func (m *MockExpertRepository) DeleteExpertiseById(expertise expertEntities.Expertise) (int64, error) {
	return m.Kode, m.Err
}

func (m *MockExpertRepository) CreateBankAccountType(bankType expertEntities.BankAccountType) (expertEntities.BankAccountType, error) {
	return bankType, m.Err
}

func (m *MockExpertRepository) GetBankAccountTypeById(bankType expertEntities.BankAccountType) (expertEntities.BankAccountType, error) {
	return bankType, m.Err
}

func (m *MockExpertRepository) GetAllBankAccountType() ([]expertEntities.BankAccountType, error) {
	return []expertEntities.BankAccountType{}, m.Err
}

func (m *MockExpertRepository) UpdateBankAccountTypeById(bankType expertEntities.BankAccountType) (expertEntities.BankAccountType, int64, error) {
	return bankType, m.Kode, m.Err
}

func (m *MockExpertRepository) DeleteBankAccountTypeById(bankType expertEntities.BankAccountType) (int64, error) {
	return m.Kode, m.Err
}

func TestRegister(t *testing.T){
	t.Run("Register Success", func(t *testing.T) {
		mockUserRepository := &MockExpertRepository{Err: nil, Kode: 0}

		expertUseCase := NewExpertUseCase(mockUserRepository)

		expert := expertEntities.Expert{
			FullName: "John Doe",
			Username: "johndoe",
			Email: "johndoe@ex.com",
			Password: "password123",
			Gender: "pria",
			Age: 25,
			Fee: 50000,
			BankAccountTypeId: 1,
			ExpertiseId: 1,
			BankAccount: expertEntities.BankAccount{
				AccountName: "John Doe",
				AccountNumber: "1234567890",
			},
		}

		registeredExpert, err := expertUseCase.Register(&expert)

		assert.NoError(t, err)
		assert.Equal(t, expert, registeredExpert)
	})

	t.Run("Register Empty Input", func(t *testing.T) {
		mockUserRepository := &MockExpertRepository{Err: nil, Kode: 0}

		expertUseCase := NewExpertUseCase(mockUserRepository)

		expert := expertEntities.Expert{}

		registeredExpert, err := expertUseCase.Register(&expert)

		assert.Error(t, err)
		assert.Equal(t, expertEntities.Expert{}, registeredExpert)
	})

	t.Run("Register Failed", func(t *testing.T) {
		mockUserRepository := &MockExpertRepository{Err: errors.New("failed to register expert"), Kode: 0}

		expertUseCase := NewExpertUseCase(mockUserRepository)

		expert := expertEntities.Expert{
			FullName: "John Doe",
			Username: "johndoe",
			Email: "johndoe@ex.com",
			Password: "password123",
			Gender: "pria",
			Age: 25,
			Fee: 50000,
			BankAccountTypeId: 1,
			ExpertiseId: 1,
			BankAccount: expertEntities.BankAccount{
				AccountName: "John Doe",
				AccountNumber: "1234567890",
			},
		}

		registeredExpert, err := expertUseCase.Register(&expert)

		assert.Error(t, err)
		assert.Equal(t, expertEntities.Expert{}, registeredExpert)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Login Success", func(t *testing.T) {
		mockUserRepository := &MockExpertRepository{Err: nil, Kode: 0}

		expertUseCase := NewExpertUseCase(mockUserRepository)

		expert := expertEntities.Expert{
			Username: "johndoe",
			Password: "password123",
		}

		loggedInExpert, err := expertUseCase.Login(&expert)

		assert.NoError(t, err)
		assert.Equal(t, expert.Username, loggedInExpert.Username)
	})

	t.Run("Login Empty Input", func(t *testing.T) {
		mockUserRepository := &MockExpertRepository{Err: nil, Kode: 0}

		expertUseCase := NewExpertUseCase(mockUserRepository)

		expert := expertEntities.Expert{}

		loggedInExpert, err := expertUseCase.Login(&expert)

		assert.Error(t, err)
		assert.Equal(t, expertEntities.Expert{}, loggedInExpert)
	})

	t.Run("Login Failed", func(t *testing.T) {
		mockUserRepository := &MockExpertRepository{Err: errors.New("failed to login expert"), Kode: 0}

		expertUseCase := NewExpertUseCase(mockUserRepository)

		expert := expertEntities.Expert{
			Username: "johndoe",
			Password: "password123",
		}

		loggedInExpert, err := expertUseCase.Login(&expert)

		assert.Error(t, err)
		assert.Equal(t, expertEntities.Expert{}, loggedInExpert)
	})
}

func TestUpdateProfileExpertById(t *testing.T) {
	t.Run("Update Profile Success", func(t *testing.T) {
		mockUserRepository := &MockExpertRepository{Err: nil, Kode: 0}

		expertUseCase := NewExpertUseCase(mockUserRepository)

		expert := expertEntities.Expert{
			Id:                1,
			Username:          "johndoe",
			Email:             "johndoe@ex.com",
			FullName:          "John Doe",
			Address:           "123 Main Street",
			Password:          "password123",
			PhoneNumber:       "123456789",
			Gender:            "pria",
			Age:               25,
			Fee:               50000,
			Experience:        3,
			ExpertiseId:       1,
			BankAccountTypeId: 1,
			BankAccount: expertEntities.BankAccount{
				AccountName:  "John Doe",
				AccountNumber: "1234567890",
			},
		}

		updatedExpert, err := expertUseCase.UpdateProfileExpertById(&expert, nil)

		assert.NoError(t, err)
		assert.Equal(t, expert.Id, updatedExpert.Id)
	})

	t.Run("Update Profile Empty Input", func(t *testing.T) {
		mockUserRepository := &MockExpertRepository{Err: nil, Kode: 0}

		expertUseCase := NewExpertUseCase(mockUserRepository)

		expert := expertEntities.Expert{}

		updatedExpert, err := expertUseCase.UpdateProfileExpertById(&expert, nil)

		assert.Error(t, err)
		assert.Equal(t, expertEntities.Expert{}, updatedExpert)
	})

	t.Run("Update Profile Failed", func(t *testing.T) {
		mockUserRepository := &MockExpertRepository{Err: errors.New("failed to update expert profile"), Kode: 0}

		expertUseCase := NewExpertUseCase(mockUserRepository)

		expert := expertEntities.Expert{
			Id:                1,
			Username:          "johndoe",
			Email:             "johndoe@ex.com",
			FullName:          "John Doe",
			Address:           "123 Main Street",
			Password:          "password123",
			PhoneNumber:       "123456789",
			Gender:            "pria",
			Age:               25,
			Fee:               50000,
			Experience:        3,
			ExpertiseId:       1,
			BankAccountTypeId: 1,
			BankAccount: expertEntities.BankAccount{
				AccountName:  "John Doe",
				AccountNumber: "1234567890",
			},
		}

		updatedExpert, err := expertUseCase.UpdateProfileExpertById(&expert, nil)

		assert.Error(t, err)
		assert.Equal(t, expertEntities.Expert{}, updatedExpert)
	})

	t.Run("Update Profile Username Already Exist", func(t *testing.T) {
		mockUserRepository := &MockExpertRepository{Err: nil, Kode: 2}

		expertUseCase := NewExpertUseCase(mockUserRepository)

		expert := expertEntities.Expert{
			Id:                1,
			Username:          "johndoe",
			Email:             "johndoe@ex.com",
			FullName:          "John Doe",
			Address:           "123 Main Street",
			Password:          "password123",
			PhoneNumber:       "123456789",
			Gender:            "pria",
			Age:               25,
			Fee:               50000,
			Experience:        3,
			ExpertiseId:       1,
			BankAccountTypeId: 1,
			BankAccount: expertEntities.BankAccount{
				AccountName:  "John Doe",
				AccountNumber: "1234567890",
			},
		}

		updatedExpert, err := expertUseCase.UpdateProfileExpertById(&expert, nil)

		assert.Error(t, err)
		assert.Equal(t, expertEntities.Expert{}, updatedExpert)
	})

	t.Run("Update Profile Email Already Exist", func(t *testing.T) {
		mockUserRepository := &MockExpertRepository{Err: nil, Kode: 3}

		expertUseCase := NewExpertUseCase(mockUserRepository)

		expert := expertEntities.Expert{
			Id:                1,
			Username:          "johndoe",
			Email:             "johndoe@ex.com",
			FullName:          "John Doe",
			Address:           "123 Main Street",
			Password:          "password123",
			PhoneNumber:       "123456789",
			Gender:            "pria",
			Age:               25,
			Fee:               50000,
			Experience:        3,
			ExpertiseId:       1,
			BankAccountTypeId: 1,
			BankAccount: expertEntities.BankAccount{
				AccountName:  "John Doe",
				AccountNumber: "1234567890",
			},
		}

		updatedExpert, err := expertUseCase.UpdateProfileExpertById(&expert, nil)

		assert.Error(t, err)
		assert.Equal(t, expertEntities.Expert{}, updatedExpert)
	})
}

func TestGetAllExperts(t *testing.T) {
	t.Run("Get All Experts Success", func(t *testing.T) {
		mockUserRepository := &MockExpertRepository{Err: nil, Kode: 0}

		expertUseCase := NewExpertUseCase(mockUserRepository)

		experts, err := expertUseCase.GetAllExperts()

		assert.NoError(t, err)
		assert.Equal(t, []expertEntities.Expert{}, experts)
	})

	t.Run("Get All Experts Failed", func(t *testing.T) {
		mockUserRepository := &MockExpertRepository{Err: errors.New("failed to get all experts"), Kode: 0}

		expertUseCase := NewExpertUseCase(mockUserRepository)

		experts, err := expertUseCase.GetAllExperts()

		assert.Error(t, err)
		assert.Equal(t, []expertEntities.Expert{}, experts)
	})
}

func TestGetExpertById(t *testing.T) {
	t.Run("Get Expert by ID Success", func(t *testing.T) {
		mockUserRepository := &MockExpertRepository{Err: nil, Kode: 0}

		expertUseCase := NewExpertUseCase(mockUserRepository)

		expert := &expertEntities.Expert{Id: 1}

		result, err := expertUseCase.GetExpertById(expert)

		assert.NoError(t, err)
		assert.Equal(t, expert.Id, result.Id)
	})

	t.Run("Get Expert by ID Failed", func(t *testing.T) {
		mockUserRepository := &MockExpertRepository{Err: errors.New("failed to get expert by ID"), Kode: 0}

		expertUseCase := NewExpertUseCase(mockUserRepository)

		expert := &expertEntities.Expert{Id: 10}

		result, err := expertUseCase.GetExpertById(expert)

		assert.Error(t, err)
		assert.Equal(t, expertEntities.Expert{}, result)
	})
}

func TestCreateExpertise(t *testing.T) {
	t.Run("Create Expertise Success", func(t *testing.T) {
		mockExpertiseRepository := &MockExpertRepository{Err: nil, Kode: 0}

		expertUseCase := NewExpertUseCase(mockExpertiseRepository)

		expertise := expertEntities.Expertise{
			Name:        "Test Expertise",
			Description: "Description for Test Expertise",
		}

		result, err := expertUseCase.CreateExpertise(expertise)

		assert.NoError(t, err)
		assert.Equal(t, expertise.Name, result.Name)
	})

	t.Run("Create Expertise Empty Input", func(t *testing.T) {
		mockExpertiseRepository := &MockExpertRepository{Err: nil, Kode: 0}

		expertUseCase := NewExpertUseCase(mockExpertiseRepository)

		expertise := expertEntities.Expertise{}

		result, err := expertUseCase.CreateExpertise(expertise)

		assert.Error(t, err)
		assert.Equal(t, expertEntities.Expertise{}, result)
	})

	t.Run("Create Expertise Failed", func(t *testing.T) {
		mockExpertiseRepository := &MockExpertRepository{Err: errors.New("failed to create expertise"), Kode: 0}

		expertUseCase := NewExpertUseCase(mockExpertiseRepository)

		expertise := expertEntities.Expertise{
			Name:        "Test Expertise",
			Description: "Description for Test Expertise",
		}

		result, err := expertUseCase.CreateExpertise(expertise)

		assert.Error(t, err)
		assert.Equal(t, expertEntities.Expertise{}, result)
	})
}

func TestGetAllExpertise(t *testing.T) {
	t.Run("Get All Expertise Success", func(t *testing.T) {
		mockExpertiseRepository := &MockExpertRepository{Err: nil, Kode: 0}

		expertUseCase := NewExpertUseCase(mockExpertiseRepository)

		expertises, err := expertUseCase.GetAllExpertise()

		assert.NoError(t, err)
		assert.Equal(t, []expertEntities.Expertise{}, expertises)
	})

	t.Run("Get All Expertise Failed", func(t *testing.T) {
		mockExpertiseRepository := &MockExpertRepository{Err: errors.New("failed to get all expertise"), Kode: 0}

		expertUseCase := NewExpertUseCase(mockExpertiseRepository)

		expertises, err := expertUseCase.GetAllExpertise()

		assert.Error(t, err)
		assert.Equal(t, []expertEntities.Expertise{}, expertises)
	})
}

func TestGetExpertiseById(t *testing.T) {
	t.Run("Get Expertise by ID Success", func(t *testing.T) {
		mockExpertiseRepository := &MockExpertRepository{Err: nil, Kode: 0}

		expertUseCase := NewExpertUseCase(mockExpertiseRepository)

		expertise := expertEntities.Expertise{Id: 1}

		result, err := expertUseCase.GetExpertiseById(expertise)

		assert.NoError(t, err)
		assert.Equal(t, expertise.Id, result.Id)
	})

	t.Run("Get Expertise by ID Failed", func(t *testing.T) {
		mockExpertiseRepository := &MockExpertRepository{Err: errors.New("failed to get expertise by ID"), Kode: 0}

		expertUseCase := NewExpertUseCase(mockExpertiseRepository)

		expertise := expertEntities.Expertise{Id: 10}

		result, err := expertUseCase.GetExpertiseById(expertise)

		assert.Error(t, err)
		assert.Equal(t, expertEntities.Expertise{}, result)
	})
}

func TestUpdateExpertiseById(t *testing.T) {
	t.Run("Update Expertise by ID Success", func(t *testing.T) {
		mockExpertiseRepository := &MockExpertRepository{Err: nil, Kode: 0}

		expertUseCase := NewExpertUseCase(mockExpertiseRepository)

		expertise := expertEntities.Expertise{
			Id:          1,
			Name:        "Test Expertise",
			Description: "Description for Test Expertise",
		}

		result, err := expertUseCase.UpdateExpertiseById(expertise)

		assert.NoError(t, err)
		assert.Equal(t, expertise, result)
	})

	t.Run("Update Expertise by ID Empty Input", func(t *testing.T) {
		mockExpertiseRepository := &MockExpertRepository{Err: nil, Kode: 0}

		expertUseCase := NewExpertUseCase(mockExpertiseRepository)

		expertise := expertEntities.Expertise{}

		result, err := expertUseCase.UpdateExpertiseById(expertise)

		assert.Error(t, err)
		assert.Equal(t, expertEntities.Expertise{}, result)
	})

	t.Run("Update Expertise by ID Failed", func(t *testing.T) {
		mockExpertiseRepository := &MockExpertRepository{Err: errors.New("failed to update expertise"), Kode: 0}

		expertUseCase := NewExpertUseCase(mockExpertiseRepository)

		expertise := expertEntities.Expertise{
			Id:          1,
			Name:        "Test Expertise",
			Description: "Description for Test Expertise",
		}

		result, err := expertUseCase.UpdateExpertiseById(expertise)

		assert.Error(t, err)
		assert.Equal(t, expertEntities.Expertise{}, result)
	})

	t.Run("Update Expertise by ID Expertise Not Found", func(t *testing.T) {
		mockExpertiseRepository := &MockExpertRepository{Err: nil, Kode: 1}

		expertUseCase := NewExpertUseCase(mockExpertiseRepository)

		expertise := expertEntities.Expertise{
			Id:          100,
			Name:        "Test Expertise",
			Description: "Description for Test Expertise",
		}

		result, err := expertUseCase.UpdateExpertiseById(expertise)

		assert.Error(t, err)
		assert.Equal(t, expertEntities.Expertise{}, result)
	})
}

func TestDeleteExpertiseById(t *testing.T) {
	t.Run("Delete Expertise by ID Success", func(t *testing.T) {
		mockExpertiseRepository := &MockExpertRepository{Err: nil, Kode: 0}

		expertUseCase := NewExpertUseCase(mockExpertiseRepository)

		expertise := expertEntities.Expertise{Id: 1}

		err := expertUseCase.DeleteExpertiseById(expertise)

		assert.NoError(t, err)
	})

	t.Run("Delete Expertise by ID Failed", func(t *testing.T) {
		mockExpertiseRepository := &MockExpertRepository{Err: errors.New("failed to delete expertise"), Kode: 0}

		expertUseCase := NewExpertUseCase(mockExpertiseRepository)

		expertise := expertEntities.Expertise{Id: 1}

		err := expertUseCase.DeleteExpertiseById(expertise)

		assert.Error(t, err)
	})

	t.Run("Delete Expertise by ID Expertise Not Found", func(t *testing.T) {
		mockExpertiseRepository := &MockExpertRepository{Err: nil, Kode: 1}

		expertUseCase := NewExpertUseCase(mockExpertiseRepository)

		expertise := expertEntities.Expertise{Id: 1}

		err := expertUseCase.DeleteExpertiseById(expertise)

		assert.Error(t, err)
	})
}

func TestCreateBankAccountType(t *testing.T) {
	t.Run("Create Bank Account Type Success", func(t *testing.T) {
		mockBankAccountTypeRepository := &MockExpertRepository{Err: nil, Kode: 0}

		expertUseCase := NewExpertUseCase(mockBankAccountTypeRepository)

		bankType := expertEntities.BankAccountType{
			Name: "Savings Account",
		}

		result, err := expertUseCase.CreateBankAccountType(bankType)

		assert.NoError(t, err)
		assert.Equal(t, bankType.Name, result.Name)
	})

	t.Run("Create Bank Account Type Empty Input", func(t *testing.T) {
		mockBankAccountTypeRepository := &MockExpertRepository{Err: nil, Kode: 0}

		expertUseCase := NewExpertUseCase(mockBankAccountTypeRepository)

		bankType := expertEntities.BankAccountType{}

		result, err := expertUseCase.CreateBankAccountType(bankType)

		assert.Error(t, err)
		assert.Equal(t, expertEntities.BankAccountType{}, result)
	})

	t.Run("Create Bank Account Type Failed", func(t *testing.T) {
		mockBankAccountTypeRepository := &MockExpertRepository{Err: errors.New("failed to create bank account type"), Kode: 0}

		expertUseCase := NewExpertUseCase(mockBankAccountTypeRepository)

		bankType := expertEntities.BankAccountType{
			Name: "Savings Account",
		}

		result, err := expertUseCase.CreateBankAccountType(bankType)

		assert.Error(t, err)
		assert.Equal(t, expertEntities.BankAccountType{}, result)
	})
}

func TestGetBankAccountTypeById(t *testing.T) {
	t.Run("Get Bank Account Type by ID Success", func(t *testing.T) {
		mockBankAccountTypeRepository := &MockExpertRepository{Err: nil, Kode: 0}

		expertUseCase := NewExpertUseCase(mockBankAccountTypeRepository)

		bankType := expertEntities.BankAccountType{Id: 1}

		result, err := expertUseCase.GetBankAccountTypeById(bankType)

		assert.NoError(t, err)
		assert.Equal(t, bankType.Id, result.Id)
	})

	t.Run("Get Bank Account Type by ID Failed", func(t *testing.T) {
		mockBankAccountTypeRepository := &MockExpertRepository{Err: errors.New("failed to get bank account type by ID"), Kode: 0}

		expertUseCase := NewExpertUseCase(mockBankAccountTypeRepository)

		bankType := expertEntities.BankAccountType{Id: 1}

		result, err := expertUseCase.GetBankAccountTypeById(bankType)

		assert.Error(t, err)
		assert.Equal(t, expertEntities.BankAccountType{}, result)
	})
}

func TestGetAllBankAccountType(t *testing.T) {
	t.Run("Get All Bank Account Type Success", func(t *testing.T) {
		mockBankAccountTypeRepository := &MockExpertRepository{Err: nil, Kode: 0}

		expertUseCase := NewExpertUseCase(mockBankAccountTypeRepository)

		bankAccountTypes, err := expertUseCase.GetAllBankAccountType()

		assert.NoError(t, err)
		assert.Equal(t, []expertEntities.BankAccountType{}, bankAccountTypes)
	})

	t.Run("Get All Bank Account Type Failed", func(t *testing.T) {
		mockBankAccountTypeRepository := &MockExpertRepository{Err: errors.New("failed to get all bank account types"), Kode: 0}

		expertUseCase := NewExpertUseCase(mockBankAccountTypeRepository)

		bankAccountTypes, err := expertUseCase.GetAllBankAccountType()

		assert.Error(t, err)
		assert.Equal(t, []expertEntities.BankAccountType{}, bankAccountTypes)
	})
}

func TestUpdateBankAccountTypeById(t *testing.T) {
	t.Run("Update Bank Account Type by ID Success", func(t *testing.T) {
		mockBankAccountTypeRepository := &MockExpertRepository{Err: nil, Kode: 0}

		expertUseCase := NewExpertUseCase(mockBankAccountTypeRepository)

		bankType := expertEntities.BankAccountType{
			Id:   1,
			Name: "Savings Account",
		}

		result, err := expertUseCase.UpdateBankAccountTypeById(bankType)

		assert.NoError(t, err)
		assert.Equal(t, bankType.Id, result.Id)
	})

	t.Run("Update Bank Account Type by ID Empty Input", func(t *testing.T) {
		mockBankAccountTypeRepository := &MockExpertRepository{Err: nil, Kode: 0}

		expertUseCase := NewExpertUseCase(mockBankAccountTypeRepository)

		bankType := expertEntities.BankAccountType{}

		result, err := expertUseCase.UpdateBankAccountTypeById(bankType)

		assert.Error(t, err)
		assert.Equal(t, expertEntities.BankAccountType{}, result)
	})

	t.Run("Update Bank Account Type by ID Failed", func(t *testing.T) {
		mockBankAccountTypeRepository := &MockExpertRepository{Err: errors.New("failed to update bank account type"), Kode: 0}

		expertUseCase := NewExpertUseCase(mockBankAccountTypeRepository)

		bankType := expertEntities.BankAccountType{
			Id:   1,
			Name: "Savings Account",
		}

		result, err := expertUseCase.UpdateBankAccountTypeById(bankType)

		assert.Error(t, err)
		assert.Equal(t, expertEntities.BankAccountType{}, result)
	})

	t.Run("Update Bank Account Type by ID Bank Account Type Not Found", func(t *testing.T) {
		mockBankAccountTypeRepository := &MockExpertRepository{Err: nil, Kode: 1}

		expertUseCase := NewExpertUseCase(mockBankAccountTypeRepository)

		bankType := expertEntities.BankAccountType{
			Id:   100,
			Name: "Savings Account",
		}

		result, err := expertUseCase.UpdateBankAccountTypeById(bankType)

		assert.Error(t, err)
		assert.Equal(t, expertEntities.BankAccountType{}, result)
	})
}

func TestDeleteBankAccountTypeById(t *testing.T) {
	t.Run("Delete Bank Account Type by ID Success", func(t *testing.T) {
		mockBankAccountTypeRepository := &MockExpertRepository{Err: nil, Kode: 0}

		expertUseCase := NewExpertUseCase(mockBankAccountTypeRepository)

		bankType := expertEntities.BankAccountType{Id: 1}

		err := expertUseCase.DeleteBankAccountTypeById(bankType)

		assert.NoError(t, err)
	})

	t.Run("Delete Bank Account Type by ID Failed", func(t *testing.T) {
		mockBankAccountTypeRepository := &MockExpertRepository{Err: errors.New("failed to delete bank account type"), Kode: 0}

		expertUseCase := NewExpertUseCase(mockBankAccountTypeRepository)

		bankType := expertEntities.BankAccountType{Id: 1}

		err := expertUseCase.DeleteBankAccountTypeById(bankType)

		assert.Error(t, err)
	})

	t.Run("Delete Bank Account Type by ID Bank Account Type Not Found", func(t *testing.T) {
		mockBankAccountTypeRepository := &MockExpertRepository{Err: nil, Kode: 1}

		expertUseCase := NewExpertUseCase(mockBankAccountTypeRepository)

		bankType := expertEntities.BankAccountType{Id: 1}

		err := expertUseCase.DeleteBankAccountTypeById(bankType)

		assert.Error(t, err)
	})
}
