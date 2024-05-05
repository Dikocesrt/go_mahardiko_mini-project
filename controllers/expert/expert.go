package expert

import (
	"habit/controllers/expert/request"
	"habit/controllers/expert/response"
	expertEntities "habit/entities/expert"
	"habit/utilities/base"
	"net/http"

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