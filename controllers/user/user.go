package user

import (
	"habit/controllers/user/request"
	"habit/controllers/user/response"
	userEntities "habit/entities/user"
	"habit/utilities/base"
	"net/http"

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