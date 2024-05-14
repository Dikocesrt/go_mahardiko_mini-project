package hire

import (
	"bytes"
	"errors"
	hireEntities "habit/entities/hire"
	"io"
	"mime/multipart"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockHireRepository struct {
	Err  error
}

func (m *MockHireRepository) CreateHire(hire *hireEntities.Hire) (hireEntities.Hire, error) {
	return *hire, m.Err
}

func (m *MockHireRepository) GetHiresByExpertId(id int) ([]hireEntities.Hire, error) {
	return []hireEntities.Hire{}, m.Err
}

func (m *MockHireRepository) GetHiresByUserId(id int) ([]hireEntities.Hire, error) {
	return []hireEntities.Hire{}, m.Err
}

func (m *MockHireRepository) GetHireById(id int) (hireEntities.Hire, error) {
	var hire hireEntities.Hire
	hire.Id = id
	return hire, m.Err
}

func (m *MockHireRepository) VerifyPayment(hire *hireEntities.Hire) (hireEntities.Hire, error) {
	return *hire, m.Err
}

func createDummyFile(t *testing.T) *multipart.FileHeader {
	t.Helper()

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	fileWriter, err := writer.CreateFormFile("file", "dummy.jpg")
	if err != nil {
		t.Fatalf("Failed to create form file: %v", err)
	}

	_, _ = io.Copy(fileWriter, bytes.NewReader([]byte("dummy file content")))
	_ = writer.Close()

	return &multipart.FileHeader{
		Filename: "dummy.jpg",
		Size:     int64(buf.Len()),
		Header:   map[string][]string{"Content-Type": {writer.FormDataContentType()}},
	}
}

func TestCreateHire(t *testing.T) {
	t.Run("Create Hire Empty Input", func(t *testing.T) {
		mockHireRepository := &MockHireRepository{Err: nil}

		hireUseCase := NewHireUseCase(mockHireRepository)

		hire := hireEntities.Hire{}

		result, err := hireUseCase.CreateHire(&hire, nil)

		assert.Error(t, err)
		assert.Equal(t, hireEntities.Hire{}, result)
	})

	t.Run("Create Hire Database Insert Failed", func(t *testing.T) {
		mockHireRepository := &MockHireRepository{Err: errors.New("failed to insert hire into database")}

		hireUseCase := NewHireUseCase(mockHireRepository)

		hire := hireEntities.Hire{
			UserId:     1,
			ExpertId:   2,
			MeetDay:    "Monday",
			MeetTime:   "10:00",
			TotalFee:   50000,
		}

		file := createDummyFile(t)

		result, err := hireUseCase.CreateHire(&hire, file)

		assert.Error(t, err)
		assert.Equal(t, hireEntities.Hire{}, result)
	})
}

func TestGetHiresByExpertId(t *testing.T) {
	t.Run("Get Hires by Expert ID Success", func(t *testing.T) {
		mockHireRepository := &MockHireRepository{Err: nil}

		hireUseCase := NewHireUseCase(mockHireRepository)

		result, err := hireUseCase.GetHiresByExpertId(2)

		assert.NoError(t, err)
		assert.Equal(t, []hireEntities.Hire{}, result)
	})

	t.Run("Get Hires by Expert ID Failed", func(t *testing.T) {
		mockHireRepository := &MockHireRepository{Err: errors.New("failed to get hires by expert ID")}

		hireUseCase := NewHireUseCase(mockHireRepository)

		result, err := hireUseCase.GetHiresByExpertId(2)

		assert.Error(t, err)
		assert.Equal(t, []hireEntities.Hire{}, result)
	})
}

func TestGetHiresByUserId(t *testing.T) {
	t.Run("Get Hires by User ID Success", func(t *testing.T) {
		mockHireRepository := &MockHireRepository{Err: nil}

		hireUseCase := NewHireUseCase(mockHireRepository)

		result, err := hireUseCase.GetHiresByUserId(1)

		assert.NoError(t, err)
		assert.Equal(t, []hireEntities.Hire{}, result)
	})

	t.Run("Get Hires by User ID Failed", func(t *testing.T) {
		mockHireRepository := &MockHireRepository{Err: errors.New("failed to get hires by user ID")}

		hireUseCase := NewHireUseCase(mockHireRepository)

		result, err := hireUseCase.GetHiresByUserId(1)

		assert.Error(t, err)
		assert.Equal(t, []hireEntities.Hire{}, result)
	})
}

func TestGetHireById(t *testing.T) {
	t.Run("Get Hire by ID Success", func(t *testing.T) {
		mockHireRepository := &MockHireRepository{Err: nil}

		hireUseCase := NewHireUseCase(mockHireRepository)

		hire := hireEntities.Hire{
			Id:       1,
		}

		result, err := hireUseCase.GetHireById(hire.Id)

		assert.NoError(t, err)
		assert.Equal(t, hire, result)
	})

	t.Run("Get Hire by ID Failed", func(t *testing.T) {
		mockHireRepository := &MockHireRepository{Err: errors.New("failed to get hire by ID")}

		hireUseCase := NewHireUseCase(mockHireRepository)

		result, err := hireUseCase.GetHireById(1)

		assert.Error(t, err)
		assert.Equal(t, hireEntities.Hire{}, result)
	})
}

func TestVerifyPayment(t *testing.T) {
	t.Run("Verify Payment Success", func(t *testing.T) {
		mockHireRepository := &MockHireRepository{Err: nil}

		hireUseCase := NewHireUseCase(mockHireRepository)

		hire := hireEntities.Hire{
			Id:       1,
			UserId:   1,
			ExpertId: 2,
			MeetUrl:  "meet_url",
		}

		result, err := hireUseCase.VerifyPayment(&hire)

		assert.NoError(t, err)
		assert.Equal(t, hire, result)
	})

	t.Run("Verify Payment Empty Input", func(t *testing.T) {
		mockHireRepository := &MockHireRepository{Err: nil}

		hireUseCase := NewHireUseCase(mockHireRepository)

		hire := hireEntities.Hire{}

		result, err := hireUseCase.VerifyPayment(&hire)

		assert.Error(t, err)
		assert.Equal(t, hireEntities.Hire{}, result)
	})

	t.Run("Verify Payment Database Update Failed", func(t *testing.T) {
		mockHireRepository := &MockHireRepository{Err: errors.New("failed to verify payment")}

		hireUseCase := NewHireUseCase(mockHireRepository)

		hire := hireEntities.Hire{
			Id:       1,
			UserId:   1,
			ExpertId: 2,
			MeetUrl:  "meet_url",
		}

		result, err := hireUseCase.VerifyPayment(&hire)

		assert.Error(t, err)
		assert.Equal(t, hireEntities.Hire{}, result)
	})
}
