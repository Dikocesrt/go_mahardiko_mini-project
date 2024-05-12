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
		return expertEntities.Expert{}, constants.ErrEmptyInputExpert
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
		return expertEntities.Expert{}, constants.ErrExpertNotFound
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
		return expertEntities.Expert{}, constants.ErrEmptyInputExpert
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(expert.Password), bcrypt.DefaultCost)
	if err != nil {
		return expertEntities.Expert{}, constants.ErrHashedPassword
	}

	expert.Password = string(hashedPassword)

	if file != nil {
		SecureURL, err := uploadImage(file)
		if err != nil {
			return expertEntities.Expert{}, constants.ErrUploadImage
		}

		expert.ProfilePicture = SecureURL
	}

	expertDb, kode, err := expertUseCase.repository.UpdateProfileExpertById(expert)
	if err != nil {
		return expertEntities.Expert{}, constants.ErrExpertNotFound
	}

	if kode == 2 {
		return expertEntities.Expert{}, constants.ErrUsernameAlreadyExist
	}

	if kode == 3 {
		return expertEntities.Expert{}, constants.ErrEmailAlreadyExist
	}

	return expertDb, nil
}

func (expertUseCase *ExpertUseCase) GetAllExperts() ([]expertEntities.Expert, error) {
	expertDb, err := expertUseCase.repository.GetAllExperts()
	if err != nil {
		return []expertEntities.Expert{}, constants.ErrGetAllData
	}

	return expertDb, nil
}

func (expertUseCase *ExpertUseCase) GetExpertById(expert *expertEntities.Expert) (expertEntities.Expert, error) {
	expertDb, err := expertUseCase.repository.GetExpertById(expert)
	if err != nil {
		return expertEntities.Expert{}, constants.ErrExpertNotFound
	}

	return expertDb, nil
}

func (expertUseCase *ExpertUseCase) CreateExpertise(expertise expertEntities.Expertise) (expertEntities.Expertise, error) {
	if expertise.Name == "" || expertise.Description == "" {
		return expertEntities.Expertise{}, constants.ErrEmptyInputActivityType
	}
	newExpertise, err := expertUseCase.repository.CreateExpertise(expertise)
	if err != nil {
		return expertEntities.Expertise{}, constants.ErrInsertDatabase
	}
	return newExpertise, nil
}

func (expertUseCase *ExpertUseCase) GetAllExpertise() ([]expertEntities.Expertise, error) {
	newExpertise, err := expertUseCase.repository.GetAllExpertise()
	if err != nil {
		return []expertEntities.Expertise{}, constants.ErrGetAllData
	}
	return newExpertise, nil
}

func (expertUseCase *ExpertUseCase) GetExpertiseById(expertise expertEntities.Expertise) (expertEntities.Expertise, error) {
	newExpertise, err := expertUseCase.repository.GetExpertiseById(expertise)
	if err != nil {
		return expertEntities.Expertise{}, constants.ErrExpertiseNotFound
	}
	return newExpertise, nil
}

func (expertUseCase *ExpertUseCase) UpdateExpertiseById(expertise expertEntities.Expertise) (expertEntities.Expertise, error) {
	if expertise.Name == "" || expertise.Description == "" {
		return expertEntities.Expertise{}, constants.ErrEmptyInputActivityType
	}
	newExpertise, err := expertUseCase.repository.UpdateExpertiseById(expertise)
	if err != nil {
		return expertEntities.Expertise{}, constants.ErrExpertiseNotFound
	}
	return newExpertise, nil
}

func (expertUseCase *ExpertUseCase) DeleteExpertiseById(expertise expertEntities.Expertise) error {
	err := expertUseCase.repository.DeleteExpertiseById(expertise)
	if err != nil {
		return constants.ErrExpertiseNotFound
	}
	return nil
}

func (expertUseCase *ExpertUseCase) CreateBankAccountType(bankType expertEntities.BankAccountType) (expertEntities.BankAccountType, error) {
	if bankType.Name == "" {
		return expertEntities.BankAccountType{}, constants.ErrEmptyInputBankAccountType
	}
	bankType, err := expertUseCase.repository.CreateBankAccountType(bankType)
	if err != nil {
		return expertEntities.BankAccountType{}, constants.ErrInsertDatabase
	}
	return bankType, nil
}

func (expertUseCase *ExpertUseCase) GetBankAccountTypeById(bankType expertEntities.BankAccountType) (expertEntities.BankAccountType, error) {
	bankType, err := expertUseCase.repository.GetBankAccountTypeById(bankType)
	if err != nil {
		return expertEntities.BankAccountType{}, constants.ErrBankAccountTypeNotFound
	}
	return bankType, nil
}

func (expertUseCase *ExpertUseCase) GetAllBankAccountType() ([]expertEntities.BankAccountType, error) {
	bankTypes, err := expertUseCase.repository.GetAllBankAccountType()
	if err != nil {
		return []expertEntities.BankAccountType{}, constants.ErrGetAllData
	}
	return bankTypes, nil
}

func (expertUseCase *ExpertUseCase) UpdateBankAccountTypeById(bankType expertEntities.BankAccountType) (expertEntities.BankAccountType, error) {
	if bankType.Name == "" {
		return expertEntities.BankAccountType{}, constants.ErrEmptyInputBankAccountType
	}
	bankType, err := expertUseCase.repository.UpdateBankAccountTypeById(bankType)
	if err != nil {
		return expertEntities.BankAccountType{}, constants.ErrBankAccountTypeNotFound
	}
	return bankType, nil
}

func (expertUseCase *ExpertUseCase) DeleteBankAccountTypeById(bankType expertEntities.BankAccountType) error {
	err := expertUseCase.repository.DeleteBankAccountTypeById(bankType)
	if err != nil {
		return constants.ErrBankAccountTypeNotFound
	}
	return nil
}