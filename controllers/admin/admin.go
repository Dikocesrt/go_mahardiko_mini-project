package admin

import (
	"habit/controllers/admin/request"
	"habit/controllers/admin/response"
	adminEntities "habit/entities/admin"
	"habit/utilities/base"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AdminController struct {
	adminUseCase adminEntities.UseCaseInterface
}

func NewAdminController(adminUseCase adminEntities.UseCaseInterface) *AdminController {
	return &AdminController{
		adminUseCase: adminUseCase,
	}
}

func (adminController *AdminController) Register(c echo.Context) error {
	var registerRequest request.AdminRegisterRequest
	c.Bind(&registerRequest)

	var adminEnt adminEntities.Admin
	adminEnt.Username = registerRequest.Username
	adminEnt.Email = registerRequest.Email
	adminEnt.Password = registerRequest.Password

	adminEnt, err := adminController.adminUseCase.Register(&adminEnt)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	adminRes := response.DetailAdmin{
		Id:       adminEnt.Id,
		Username: adminEnt.Username,
		Email:    adminEnt.Email,
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Register", adminRes))
}

func (adminController *AdminController) Login(c echo.Context) error {
	var loginRequest request.AdminLoginRequest
	c.Bind(&loginRequest)

	var adminEnt adminEntities.Admin
	adminEnt.Username = loginRequest.Username
	adminEnt.Password = loginRequest.Password

	adminEnt, err := adminController.adminUseCase.Login(&adminEnt)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	adminRes := response.AdminLoginResponse{
		Id:       adminEnt.Id,
		Username: adminEnt.Username,
		Email:    adminEnt.Email,
		Token:    adminEnt.Token,
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Login", adminRes))
}