package user

import (
	"context"
	"habit/configs"
	"habit/constants"
	userEntitites "habit/entities/user"
	"habit/middlewares"
	"mime/multipart"

	"golang.org/x/crypto/bcrypt"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

type UserUseCase struct {
	repository userEntitites.RepositoryInterface
}

func NewUserUseCase(repository userEntitites.RepositoryInterface) *UserUseCase {
	return &UserUseCase{
		repository: repository,
	}
}

func (userUseCase *UserUseCase) Register(user *userEntitites.User) (userEntitites.User, error) {
	if user.FullName == "" || user.Username == "" || user.Email == "" || user.Password == "" {
		return userEntitites.User{}, constants.ErrEmptyInputRegistration
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return userEntitites.User{}, constants.ErrHashedPassword
	}

	user.Password = string(hashedPassword)
	
	newUser, err := userUseCase.repository.Register(user)
	if err != nil {
		return userEntitites.User{}, constants.ErrInsertDatabase
	}

	return newUser, nil
}

func (userUseCase *UserUseCase) Login(user *userEntitites.User) (userEntitites.User, error) {
	if user.Username == "" || user.Password == "" {
		return userEntitites.User{}, constants.ErrEmptyInputLogin
	}

	userFromDb, err := userUseCase.repository.Login(user)
	if err != nil {
		return userEntitites.User{}, constants.ErrUserNotFound
	}

	token, _ := middlewares.CreateToken(userFromDb.Id)
	userFromDb.Token = token

	return userFromDb, nil
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

func (userUseCase *UserUseCase) UpdateProfileById(user *userEntitites.User, file *multipart.FileHeader) (userEntitites.User, error) {
	if user.FullName == "" || user.Username == "" || user.Email == "" || user.Password == "" {
		return userEntitites.User{}, constants.ErrEmptyInputUpdateProfile
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return userEntitites.User{}, constants.ErrHashedPassword
	}

	user.Password = string(hashedPassword)

	if file != nil {
		SecureURL, err := uploadImage(file)
		if err != nil {
			return userEntitites.User{}, err
		}

		user.ProfilePicture = SecureURL
	}

	userFromDb, kode, err  := userUseCase.repository.UpdateProfileById(user)
	if err != nil{
		return userEntitites.User{}, constants.ErrUpdateDatabase
	}

	if kode == 2 {
		return userEntitites.User{}, constants.ErrUsernameAlreadyExist
	}

	if kode == 3 {
		return userEntitites.User{}, constants.ErrEmailAlreadyExist
	}

	return userFromDb, nil
}