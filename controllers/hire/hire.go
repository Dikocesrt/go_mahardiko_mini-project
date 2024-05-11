package hire

import (
	"habit/controllers/hire/request"
	"habit/controllers/hire/response"
	hireEntities "habit/entities/hire"
	"habit/utilities/base"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HireController struct {
	hireUseCase hireEntities.UseCaseInterface
}

func NewHireController(hireUseCase hireEntities.UseCaseInterface) *HireController {
	return &HireController{
		hireUseCase: hireUseCase,
	}
}

func (hireController *HireController) CreateHire(c echo.Context) error {
	var hireFromRequest request.HireCreateRequest
	c.Bind(&hireFromRequest)

	file, _ := c.FormFile("payment_image")
	hireEnt := hireEntities.Hire{
		UserId:       hireFromRequest.UserId,
		ExpertId:     hireFromRequest.ExpertId,
		MeetDay:      hireFromRequest.MeetDay,
		MeetTime:     hireFromRequest.MeetTime,
		PaymentImage: hireFromRequest.PaymentImage,
		TotalFee:     hireFromRequest.TotalFee,
	}

	hireEnt, err := hireController.hireUseCase.CreateHire(&hireEnt, file)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	hireResponse := response.HireCreateResponse{
		Id:            hireEnt.Id,
		MeetDay:       hireEnt.MeetDay,
		MeetTime:      hireEnt.MeetTime,
		TotalFee:      hireEnt.TotalFee,
		PaymentStatus: hireEnt.PaymentStatus,
		PaymentImage:  hireEnt.PaymentImage,
		User: response.UserCreateResponse{
			Id:       hireEnt.User.Id,
			Username: hireEnt.User.Username,
			Email:    hireEnt.User.Email,
			FullName: hireEnt.User.FullName,
		},
		Expert: response.ExpertCreateResponse{
			Id:             hireEnt.Expert.Id,
			Username:       hireEnt.Expert.Username,
			Email:          hireEnt.Expert.Email,
			FullName:       hireEnt.Expert.FullName,
		},
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Create Hire", hireResponse))
}

func (hireController *HireController) GetHiresByExpertId(c echo.Context) error {
	expertId := c.Param("expertId")
	id, _ := strconv.Atoi(expertId)

	hiresEnt, err := hireController.hireUseCase.GetHiresByExpertId(id)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	hiresResponse := []response.HireDetailResponse{}
	for _, hireEnt := range hiresEnt {
		hiresResponse = append(hiresResponse, response.HireDetailResponse{
			Id:            hireEnt.Id,
			HireStart:     hireEnt.HireStart,
			HireEnd:       hireEnt.HireEnd,
			TotalFee:      hireEnt.TotalFee,
			PaymentStatus: hireEnt.PaymentStatus,
			PaymentImage:  hireEnt.PaymentImage,
			MeetTime:      hireEnt.MeetTime,
			MeetDay:       hireEnt.MeetDay,
			MeetUrl:       hireEnt.MeetUrl,
			User: response.UserDetailResponse{
				Id:       hireEnt.User.Id,
				Username: hireEnt.User.Username,
				Email:    hireEnt.User.Email,
				FullName: hireEnt.User.FullName,
			},
			Expert: response.ExpertDetailResponse{
				Id:             hireEnt.Expert.Id,
				Username:       hireEnt.Expert.Username,
				Email:          hireEnt.Expert.Email,
				FullName:       hireEnt.Expert.FullName,
			},
		})
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get Hires", hiresResponse))
}

func (hireController *HireController) VerifyPayment(c echo.Context) error {
	hireId := c.Param("hireId")
	id, _ := strconv.Atoi(hireId)

	var hireFromRequest request.HireVerifyRequest
	c.Bind(&hireFromRequest)

	hireEnt := hireEntities.Hire{
		Id:            id,
		PaymentStatus: hireFromRequest.PaymentStatus,
		MeetUrl:       hireFromRequest.MeetUrl,
	}

	hireEnt, err := hireController.hireUseCase.VerifyPayment(&hireEnt)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	hireResponse := response.HireDetailResponse{
		Id:            hireEnt.Id,
		HireStart:     hireEnt.HireStart,
		HireEnd:       hireEnt.HireEnd,
		TotalFee:      hireEnt.TotalFee,
		PaymentStatus: hireEnt.PaymentStatus,
		PaymentImage:  hireEnt.PaymentImage,
		MeetTime:      hireEnt.MeetTime,
		MeetDay:       hireEnt.MeetDay,
		MeetUrl:       hireEnt.MeetUrl,
		User: response.UserDetailResponse{
			Id:       hireEnt.User.Id,
			Username: hireEnt.User.Username,
			Email:    hireEnt.User.Email,
			FullName: hireEnt.User.FullName,
		},
		Expert: response.ExpertDetailResponse{
			Id:             hireEnt.Expert.Id,
			Username:       hireEnt.Expert.Username,
			Email:          hireEnt.Expert.Email,
			FullName:       hireEnt.Expert.FullName,
		},
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Verify Payment", hireResponse))
}