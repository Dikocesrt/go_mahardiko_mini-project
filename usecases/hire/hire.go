package hire

import (
	"context"
	"habit/configs"
	hireEntities "habit/entities/hire"
	"mime/multipart"
	"time"

	"habit/constants"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

type HireUseCase struct {
	repository hireEntities.RepositoryInterface
}

func NewHireUseCase(repository hireEntities.RepositoryInterface) *HireUseCase {
	return &HireUseCase{
		repository: repository,
	}
}

func uploadImage(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	cloudinaryURL := configs.InitConfigCloudinary()
	if cloudinaryURL == "" {
		return "", constants.ErrCloudinary
	}

	cld, err := cloudinary.NewFromURL(cloudinaryURL)
	if err != nil {
		return "", err
	}

	uploadResult, err := cld.Upload.Upload(context.Background(), src, uploader.UploadParams{})
	if err != nil {
		return "", err
	}

	return uploadResult.SecureURL, nil
}

func (hireUseCase *HireUseCase) CreateHire(hire *hireEntities.Hire, file *multipart.FileHeader) (hireEntities.Hire, error) {
	if file == nil {
		return hireEntities.Hire{}, constants.ErrEmptyImageInput
	}

	if file != nil {
		SecureURL, err := uploadImage(file)
		if err != nil {
			return hireEntities.Hire{}, constants.ErrUploadImage
		}

		hire.PaymentImage = SecureURL
	}

	hire.PaymentStatus = "pending"

	hireFromDb, err := hireUseCase.repository.CreateHire(hire)
	if err != nil {
		return hireEntities.Hire{}, constants.ErrInsertDatabase
	}

	return hireFromDb, nil
}

func (hireUseCase *HireUseCase) GetHiresByExpertId(id int) ([]hireEntities.Hire, error) {
	hiresFromDb, err := hireUseCase.repository.GetHiresByExpertId(id)
	if err != nil {
		return []hireEntities.Hire{}, constants.ErrGetDatabase
	}

	return hiresFromDb, nil
}

func (hireUseCase *HireUseCase) GetHiresByUserId(id int) ([]hireEntities.Hire, error) {
	hiresFromDb, err := hireUseCase.repository.GetHiresByUserId(id)
	if err != nil {
		return []hireEntities.Hire{}, constants.ErrGetDatabase
	}

	return hiresFromDb, nil
}

func (hireUseCase *HireUseCase) GetHireById(id int) (hireEntities.Hire, error) {
	hireFromDb, err := hireUseCase.repository.GetHireById(id)
	if err != nil {
		return hireEntities.Hire{}, constants.ErrGetDatabase
	}

	return hireFromDb, nil
}

func (hireUseCase *HireUseCase) VerifyPayment(hire *hireEntities.Hire) (hireEntities.Hire, error) {
	hire.HireStart = time.Now()
	hire.HireEnd = time.Now().AddDate(0, 0, 30)
	hire.PaymentStatus = "paid"
	hireFromDb, err := hireUseCase.repository.VerifyPayment(hire)
	if err != nil {
		return hireEntities.Hire{}, constants.ErrGetDatabase
	}

	return hireFromDb, nil
}