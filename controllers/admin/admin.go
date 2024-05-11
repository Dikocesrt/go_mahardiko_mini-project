package admin

import (
	"habit/controllers/admin/request"
	"habit/controllers/admin/response"
	adminEntities "habit/entities/admin"
	"habit/entities/expert"
	"habit/utilities/base"
	"net/http"
	"strconv"

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

func (adminController *AdminController) CreateBankAccountType(c echo.Context) error {
	var createBankRequest request.BankTypeRequest
	c.Bind(&createBankRequest)

	var bankType expert.BankAccountType

	bankType.Name = createBankRequest.Name

	bankType, err := adminController.adminUseCase.CreateBankAccountType(bankType)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	bankResponse := response.BankTypeResponse{
		Id:   bankType.Id,
		Name: bankType.Name,
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Create Bank Account Type", bankResponse))
}

func (adminController *AdminController) GetBankAccountTypeById(c echo.Context) error {
	bankAccountTypeId := c.Param("bankAccountTypeId")
	id, _ := strconv.Atoi(bankAccountTypeId)

	var bankType expert.BankAccountType
	bankType.Id = id

	bankType, err := adminController.adminUseCase.GetBankAccountTypeById(bankType)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	bankResponse := response.BankTypeResponse{
		Id:   bankType.Id,
		Name: bankType.Name,
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get Bank Account Type", bankResponse))
}

func (adminController *AdminController) GetAllBankAccountType(c echo.Context) error {
	bankTypes, err := adminController.adminUseCase.GetAllBankAccountType()

	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	bankTypesResponse := make([]response.BankTypeResponse, len(bankTypes))

	for i := 0; i < len(bankTypes); i++ {
		bankTypesResponse[i].Id = bankTypes[i].Id
		bankTypesResponse[i].Name = bankTypes[i].Name
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get All Bank Account Type", bankTypesResponse))
}

func (adminController *AdminController) UpdateBankAccountTypeById(c echo.Context) error {
	var createBankRequest request.BankTypeRequest
	c.Bind(&createBankRequest)

	bankAccountTypeId := c.Param("bankAccountTypeId")
	id, _ := strconv.Atoi(bankAccountTypeId)

	var bankType expert.BankAccountType
	bankType.Id = id
	bankType.Name = createBankRequest.Name

	bankType, err := adminController.adminUseCase.UpdateBankAccountTypeById(bankType)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	bankResponse := response.BankTypeResponse{
		Id:   bankType.Id,
		Name: bankType.Name,
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Update Bank Account Type", bankResponse))
}

func (adminController *AdminController) DeleteBankAccountTypeById(c echo.Context) error {
	bankAccountTypeId := c.Param("bankAccountTypeId")
	id, _ := strconv.Atoi(bankAccountTypeId)

	var bankType expert.BankAccountType
	bankType.Id = id

	err := adminController.adminUseCase.DeleteBankAccountTypeById(bankType)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Delete Bank Account Type", nil))
}