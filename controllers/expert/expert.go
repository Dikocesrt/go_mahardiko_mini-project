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
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Register", expertResponse))
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
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Update", expertResponse))
}