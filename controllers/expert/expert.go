package expert

import (
	"habit/controllers/expert/request"
	"habit/controllers/expert/response"
	expertEntities "habit/entities/expert"
	"habit/utilities/base"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ExpertController struct {
	expertUseCase expertEntities.UseCaseInterface
}

func NewExpertController(expertUseCase expertEntities.UseCaseInterface) *ExpertController {
	return &ExpertController{
		expertUseCase: expertUseCase,
	}
}

func (expertController *ExpertController) Register(c echo.Context) error {
	var expertFromRequest request.ExpertRegisterRequest
	c.Bind(&expertFromRequest)

	expertEntities := expertEntities.Expert{
		FullName: expertFromRequest.FullName,
		Username: expertFromRequest.Username,
		Email:    expertFromRequest.Email,
		Password: expertFromRequest.Password,
		Gender:   expertFromRequest.Gender,
		Age:      expertFromRequest.Age,
		Experience: expertFromRequest.Experience,
		Fee:      expertFromRequest.Fee,
		BankAccountTypeId: expertFromRequest.BankAccountTypeId,
		ExpertiseId:   expertFromRequest.ExpertiseId,
	}

	expertEntities.BankAccount.AccountName = expertFromRequest.AccountName
	expertEntities.BankAccount.AccountNumber = expertFromRequest.AccountNumber

	newExpert, err := expertController.expertUseCase.Register(&expertEntities)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	expertResponse := response.ExpertRegisterResponse{
		Id:       newExpert.Id,
		Username: newExpert.Username,
		Email:    newExpert.Email,
		FullName: newExpert.FullName,
		Gender:   newExpert.Gender,
		Age:      newExpert.Age,
		Experience: newExpert.Experience,
		Fee:      newExpert.Fee,
		BankAccount: response.BankAccountResponse{
			AccountName: newExpert.BankAccount.AccountName,
			AccountNumber: newExpert.BankAccount.AccountNumber,
		},
		Expertise: response.ExpertiseResponse{
			Name:        newExpert.Expertise.Name,
			Description: newExpert.Expertise.Description,
		},
	}
	return c.JSON(http.StatusCreated, base.NewSuccessResponse("Success Register", expertResponse))
}

func (expertController *ExpertController) Login(c echo.Context) error {
	var expertFromRequest request.ExpertLoginRequest
	c.Bind(&expertFromRequest)

	expertEntities := expertEntities.Expert{
		Username:    expertFromRequest.Username,
		Password: expertFromRequest.Password,
	}

	newExpert, err := expertController.expertUseCase.Login(&expertEntities)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	expertResponse := response.ExpertLoginResponse{
		Id:       newExpert.Id,
		Username: newExpert.Username,
		Email:    newExpert.Email,
		Token:    newExpert.Token,
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Login", expertResponse))
}

func (expertController *ExpertController) UpdateProfileExpertById(c echo.Context) error {
	var expertFromRequest request.ExpertProfileRequest
	c.Bind(&expertFromRequest)

	file, _ := c.FormFile("profile_picture")

	expertId := c.Param("id")
	id, _ := strconv.Atoi(expertId)

	expertEntities := expertEntities.Expert{
		Id: id,
		FullName: expertFromRequest.FullName,
		Username: expertFromRequest.Username,
		Email:    expertFromRequest.Email,
		Password: expertFromRequest.Password,
		Address:  expertFromRequest.Address,
		Bio:      expertFromRequest.Bio,
		PhoneNumber: expertFromRequest.PhoneNumber,
		Gender:   expertFromRequest.Gender,
		Age:      expertFromRequest.Age,
		ProfilePicture: file.Filename,
		Experience: expertFromRequest.Experience,
		Fee:      expertFromRequest.Fee,
		BankAccountTypeId: expertFromRequest.BankAccountTypeId,
		BankAccount: expertEntities.BankAccount{
			AccountName: expertFromRequest.AccountName,
			AccountNumber: expertFromRequest.AccountNumber,
		},
	}

	newExpert, err := expertController.expertUseCase.UpdateProfileExpertById(&expertEntities, file)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	expertResponse := response.ExpertProfileResponse{
		Id:       newExpert.Id,
		Username: newExpert.Username,
		Email:    newExpert.Email,
		FullName: newExpert.FullName,
		Address:  newExpert.Address,
		Bio:      newExpert.Bio,
		PhoneNumber: newExpert.PhoneNumber,
		Gender:   newExpert.Gender,
		Age:      newExpert.Age,
		ProfilePicture: newExpert.ProfilePicture,
		Experience: newExpert.Experience,
		Fee:      newExpert.Fee,
		BankAccount: response.BankAccountProfileResponse{
			AccountName: newExpert.BankAccount.AccountName,
			AccountNumber: newExpert.BankAccount.AccountNumber,
		},
	}
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Update Profile", expertResponse))
}

func (expertController *ExpertController) GetAllExperts(c echo.Context) error {
	expertEnt, err := expertController.expertUseCase.GetAllExperts()
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	expertResponse := make([]response.ExpertDetailResponse, len(expertEnt))
	for i := 0; i < len(expertEnt); i++ {
		expertResponse[i] = response.ExpertDetailResponse{
			Id:       expertEnt[i].Id,
			Username: expertEnt[i].Username,
			Email:    expertEnt[i].Email,
			FullName: expertEnt[i].FullName,
			Address:  expertEnt[i].Address,
			Bio:      expertEnt[i].Bio,
			PhoneNumber: expertEnt[i].PhoneNumber,
			Gender:   expertEnt[i].Gender,
			Age:      expertEnt[i].Age,
			ProfilePicture: expertEnt[i].ProfilePicture,
			Experience: expertEnt[i].Experience,
			Fee:      expertEnt[i].Fee,
			BankAccount: response.BankAccountProfileResponse{
				AccountName: expertEnt[i].BankAccount.AccountName,
				AccountNumber: expertEnt[i].BankAccount.AccountNumber,
			},
			Expertise: response.ExpertiseDetailResponse{
				Name: expertEnt[i].Expertise.Name,
				Description: expertEnt[i].Expertise.Description,
			},
		}
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get Experts", expertResponse))
}

func (expertController *ExpertController) GetExpertById(c echo.Context) error {
	expertId := c.Param("id")
	id, _ := strconv.Atoi(expertId)

	var expertEnt = expertEntities.Expert{}
	expertEnt.Id = id

	expertEnt, err := expertController.expertUseCase.GetExpertById(&expertEnt)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	expertResponse := response.ExpertDetailResponse{
		Id:       expertEnt.Id,
		Username: expertEnt.Username,
		Email:    expertEnt.Email,
		FullName: expertEnt.FullName,
		Address:  expertEnt.Address,
		Bio:      expertEnt.Bio,
		PhoneNumber: expertEnt.PhoneNumber,
		Gender:   expertEnt.Gender,
		Age:      expertEnt.Age,
		ProfilePicture: expertEnt.ProfilePicture,
		Experience: expertEnt.Experience,
		Fee:      expertEnt.Fee,
		BankAccount: response.BankAccountProfileResponse{
			AccountName: expertEnt.BankAccount.AccountName,
			AccountNumber: expertEnt.BankAccount.AccountNumber,
		},
		Expertise: response.ExpertiseDetailResponse{
			Name: expertEnt.Expertise.Name,
			Description: expertEnt.Expertise.Description,
		},
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get Expert", expertResponse))
}

func (expertController *ExpertController) CreateExpertise(c echo.Context) error {
	var expertiseRequest request.ExpertiseRequest
	c.Bind(&expertiseRequest)

	var expertise expertEntities.Expertise
	expertise.Name = expertiseRequest.Name
	expertise.Description = expertiseRequest.Description

	expertise, err := expertController.expertUseCase.CreateExpertise(expertise)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	expertiseResponse := response.ExpertiseAdminResponse{
		Id:          expertise.Id,
		Name:        expertise.Name,
		Description: expertise.Description,
	}

	return c.JSON(http.StatusCreated, base.NewSuccessResponse("Success Create Expertise", expertiseResponse))
}

func (expertController *ExpertController) GetAllExpertise(c echo.Context) error {
	expertises, err := expertController.expertUseCase.GetAllExpertise()
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	expertiseResponse := make([]response.ExpertiseAdminResponse, len(expertises))

	for i := 0; i < len(expertises); i++ {
		expertiseResponse[i].Id = expertises[i].Id
		expertiseResponse[i].Name = expertises[i].Name
		expertiseResponse[i].Description = expertises[i].Description
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get Expertises", expertiseResponse))
}

func (expertController *ExpertController) GetExpertiseById(c echo.Context) error {
	expertiseId := c.Param("expertiseId")
	id, _ := strconv.Atoi(expertiseId)

	var expertise expertEntities.Expertise
	expertise.Id = id

	expertise, err := expertController.expertUseCase.GetExpertiseById(expertise)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	expertiseResponse := response.ExpertiseAdminResponse{
		Id:          expertise.Id,
		Name:        expertise.Name,
		Description: expertise.Description,
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get Expertise", expertiseResponse))
}

func (expertController *ExpertController) UpdateExpertiseById(c echo.Context) error {
	var expertiseRequest request.ExpertiseRequest
	c.Bind(&expertiseRequest)

	expertiseId := c.Param("expertiseId")
	id, _ := strconv.Atoi(expertiseId)

	var expertise expertEntities.Expertise
	expertise.Id = id
	expertise.Name = expertiseRequest.Name
	expertise.Description = expertiseRequest.Description

	expertise, err := expertController.expertUseCase.UpdateExpertiseById(expertise)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	expertiseResponse := response.ExpertiseAdminResponse{
		Id:          expertise.Id,
		Name:        expertise.Name,
		Description: expertise.Description,
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Update Expertise", expertiseResponse))
}

func (expertController *ExpertController) DeleteExpertiseById(c echo.Context) error {
	expertiseId := c.Param("expertiseId")
	id, _ := strconv.Atoi(expertiseId)

	var expertise expertEntities.Expertise
	expertise.Id = id

	err := expertController.expertUseCase.DeleteExpertiseById(expertise)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Delete Expertise", nil))
}

func (expertController *ExpertController) CreateBankAccountType(c echo.Context) error {
	var createBankRequest request.BankTypeRequest
	c.Bind(&createBankRequest)

	var bankType expertEntities.BankAccountType

	bankType.Name = createBankRequest.Name

	bankType, err := expertController.expertUseCase.CreateBankAccountType(bankType)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	bankResponse := response.BankTypeResponse{
		Id:   bankType.Id,
		Name: bankType.Name,
	}

	return c.JSON(http.StatusCreated, base.NewSuccessResponse("Success Create Bank Account Type", bankResponse))
}

func (expertController *ExpertController) GetBankAccountTypeById(c echo.Context) error {
	bankAccountTypeId := c.Param("bankAccountTypeId")
	id, _ := strconv.Atoi(bankAccountTypeId)

	var bankType expertEntities.BankAccountType
	bankType.Id = id

	bankType, err := expertController.expertUseCase.GetBankAccountTypeById(bankType)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	bankResponse := response.BankTypeResponse{
		Id:   bankType.Id,
		Name: bankType.Name,
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get Bank Account Type", bankResponse))
}

func (expertController *ExpertController) GetAllBankAccountType(c echo.Context) error {
	bankTypes, err := expertController.expertUseCase.GetAllBankAccountType()

	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	bankTypesResponse := make([]response.BankTypeResponse, len(bankTypes))

	for i := 0; i < len(bankTypes); i++ {
		bankTypesResponse[i].Id = bankTypes[i].Id
		bankTypesResponse[i].Name = bankTypes[i].Name
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get Bank Account Types", bankTypesResponse))
}

func (expertController *ExpertController) UpdateBankAccountTypeById(c echo.Context) error {
	var createBankRequest request.BankTypeRequest
	c.Bind(&createBankRequest)

	bankAccountTypeId := c.Param("bankAccountTypeId")
	id, _ := strconv.Atoi(bankAccountTypeId)

	var bankType expertEntities.BankAccountType
	bankType.Id = id
	bankType.Name = createBankRequest.Name

	bankType, err := expertController.expertUseCase.UpdateBankAccountTypeById(bankType)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	bankResponse := response.BankTypeResponse{
		Id:   bankType.Id,
		Name: bankType.Name,
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Update Bank Account Type", bankResponse))
}

func (expertController *ExpertController) DeleteBankAccountTypeById(c echo.Context) error {
	bankAccountTypeId := c.Param("bankAccountTypeId")
	id, _ := strconv.Atoi(bankAccountTypeId)

	var bankType expertEntities.BankAccountType
	bankType.Id = id

	err := expertController.expertUseCase.DeleteBankAccountTypeById(bankType)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Delete Bank Account Type", nil))
}