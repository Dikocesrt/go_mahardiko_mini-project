package user

import (
	"habit/controllers/user/request"
	"habit/controllers/user/response"
	userEntities "habit/entities/user"
	"habit/utilities/base"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUseCase userEntities.UseCaseInterface
}

func NewUserController(userUseCase userEntities.UseCaseInterface) *UserController {
	return &UserController{
		userUseCase: userUseCase,
	}
}

func (userController *UserController) Register(c echo.Context) error {
	var userFromRequest request.UserRegisterRequest
	c.Bind(&userFromRequest)

	userEntities := userEntities.User{
		FullName: userFromRequest.FullName,
		Username: userFromRequest.Username,
		Email:    userFromRequest.Email,
		Password: userFromRequest.Password,
	}

	newUser, err := userController.userUseCase.Register(&userEntities)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	userResponse := response.UserRegisterResponse{
		Id:    newUser.Id,
		FullName: newUser.FullName,
		Username:  newUser.Username,
		Email: newUser.Email,
	}
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Register", userResponse))
}

func (userController *UserController) Login(c echo.Context) error {
	var userFromRequest request.UserLoginRequest
	c.Bind(&userFromRequest)

	userEntities := userEntities.User{
		Username: userFromRequest.Username,
		Password: userFromRequest.Password,
	}

	userFromDb, err := userController.userUseCase.Login(&userEntities)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	userResponse := response.UserLoginResponse{
		Id:    userFromDb.Id,
		Username:  userFromDb.Username,
		Email: userFromDb.Email,
		Token: userFromDb.Token,
	}
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Login", userResponse))
}

func (userController *UserController) UpdateProfileById(c echo.Context) error {
	var userFromRequest request.UserProfileRequest
	c.Bind(&userFromRequest)

	file, _ := c.FormFile("profile_picture")

	userId := c.Param("id")
	id, _ := strconv.Atoi(userId)

	userEntities := userEntities.User{
		Id: id,
		FullName: userFromRequest.FullName,
		Username: userFromRequest.Username,
		Email:    userFromRequest.Email,
		Password: userFromRequest.Password,
		Address:  userFromRequest.Address,
		Bio:      userFromRequest.Bio,
		PhoneNumber: userFromRequest.PhoneNumber,
		Gender:   userFromRequest.Gender,
		Age:      userFromRequest.Age,
		ProfilePicture: userFromRequest.ProfilePicture,
	}

	userFromDb, err := userController.userUseCase.UpdateProfileById(&userEntities, file)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	userResponse := response.UserProfileResponse{
		Id:             userFromDb.Id,
		FullName:       userFromDb.FullName,
		Username:       userFromDb.Username,
		Email:          userFromDb.Email,
		Address:        userFromDb.Address,
		Bio:            userFromDb.Bio,
		PhoneNumber:    userFromDb.PhoneNumber,
		Gender:         userFromDb.Gender,
		Age:            userFromDb.Age,
		ProfilePicture: userFromDb.ProfilePicture,
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Update Profile", userResponse))
}

func (userController *UserController) GetUserById(c echo.Context) error {
	userId := c.Param("id")
	id, _ := strconv.Atoi(userId)

	var userEnt userEntities.User

	userEnt.Id = id

	userEnt, err := userController.userUseCase.GetUserById(&userEnt)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	userResponse := response.UserProfileResponse{
		Id:             userEnt.Id,
		FullName:       userEnt.FullName,
		Username:       userEnt.Username,
		Email:          userEnt.Email,
		Address:        userEnt.Address,
		Bio:            userEnt.Bio,
		PhoneNumber:    userEnt.PhoneNumber,
		Gender:         userEnt.Gender,
		Age:            userEnt.Age,
		ProfilePicture: userEnt.ProfilePicture,
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get User", userResponse))
}