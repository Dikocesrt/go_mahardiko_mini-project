package expert

import (
	"context"
	"habit/configs"
	"habit/constants"
	expertEntities "habit/entities/expert"
	"habit/middlewares"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"golang.org/x/crypto/bcrypt"
)

type ExpertUseCase struct {
	repository expertEntities.RepositoryInterface
}

func NewExpertUseCase(repository expertEntities.RepositoryInterface) *ExpertUseCase {
	return &ExpertUseCase{
		repository: repository,
	}
}

func (expertUseCase *ExpertUseCase) Register(expert *expertEntities.Expert) (expertEntities.Expert, error) {
	if expert.FullName == "" || expert.Username == "" || expert.Email == "" || expert.Password == "" || expert.Gender == "" || expert.Age == 0 || expert.Fee == 0 || expert.BankAccountTypeId == 0 || expert.BankAccount.AccountName == "" || expert.BankAccount.AccountNumber == "" || expert.ExpertiseId == 0 {
		return expertEntities.Expert{}, constants.ErrEmptyInputRegistration
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(expert.Password), bcrypt.DefaultCost)
	if err != nil {
		return expertEntities.Expert{}, constants.ErrHashedPassword
	}

	expert.Password = string(hashedPassword)

	newExpert, err := expertUseCase.repository.Register(expert)
	if err != nil {
		return expertEntities.Expert{}, constants.ErrInsertDatabase
	}

	return newExpert, nil
}

func (expertUseCase *ExpertUseCase) Login(expert *expertEntities.Expert) (expertEntities.Expert, error) {
	if expert.Username == "" || expert.Password == "" {
		return expertEntities.Expert{}, constants.ErrEmptyInputLogin
	}

	expertDb, err := expertUseCase.repository.Login(expert)
	if err != nil {
		return expertEntities.Expert{}, constants.ErrUserNotFound
	}

	token, _ := middlewares.CreateTokenExpert(expertDb.Id)
	expertDb.Token = token

	return expertDb, nil
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

func (expertUseCase *ExpertUseCase) UpdateProfileExpertById(expert *expertEntities.Expert, file *multipart.FileHeader) (expertEntities.Expert, error) {
	if expert.Username == "" || expert.Email == "" || expert.FullName == "" || expert.Address == "" || expert.Password == "" || expert.PhoneNumber == "" || expert.Gender == "" || expert.Age == 0 || expert.BankAccount.AccountName == "" || expert.BankAccount.AccountNumber == "" || expert.Experience == 0 || expert.Fee == 0 || expert.BankAccountTypeId == 0 {
		return expertEntities.Expert{}, constants.ErrEmptyInputUpdateProfile
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(expert.Password), bcrypt.DefaultCost)
	if err != nil {
		return expertEntities.Expert{}, constants.ErrHashedPassword
	}

	expert.Password = string(hashedPassword)

	if file != nil {
		SecureURL, err := uploadImage(file)
		if err != nil {
			return expertEntities.Expert{}, err
		}

		expert.ProfilePicture = SecureURL
	}

	expertDb, kode, err := expertUseCase.repository.UpdateProfileExpertById(expert)
	if err != nil {
		return expertEntities.Expert{}, constants.ErrUpdateDatabase
	}

	if kode == 2 {
		return expertEntities.Expert{}, constants.ErrUsernameAlreadyExist
	}

	if kode == 3 {
		return expertEntities.Expert{}, constants.ErrEmailAlreadyExist
	}

	return expertDb, nil
}