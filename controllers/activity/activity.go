package activity

import (
	"habit/controllers/activity/request"
	"habit/controllers/activity/response"
	activityEntities "habit/entities/activity"
	"habit/utilities/base"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ActivityController struct {
	activityUseCase activityEntities.UseCaseInterface
}

func NewActivityController(activityUseCase activityEntities.UseCaseInterface) *ActivityController {
	return &ActivityController{
		activityUseCase: activityUseCase,
	}
}

func (activityController *ActivityController) CreateActivity(c echo.Context) error {
	var activityReq request.ActivityCreateRequest
	c.Bind(&activityReq)

	activityEnt := activityEntities.Activity{
		Title:          activityReq.Title,
		ActivityStart:  activityReq.ActivityStart,
		ActivityFinish: activityReq.ActivityFinish,
		UserId:         activityReq.UserId,
		ActivityTypeId: activityReq.ActivityTypeId,
		ActivityDetail: activityEntities.ActivityDetail{
			HeartRate:      activityReq.ActivityDetail.HeartRate,
			Intensity:      activityReq.ActivityDetail.Intensity,
			CaloriesBurned: activityReq.ActivityDetail.CaloriesBurned,
			FoodDetails:    activityReq.ActivityDetail.FoodDetails,
			ImageUrl:       activityReq.ActivityDetail.ImageUrl,
		},
	}

	activityEnt, err := activityController.activityUseCase.CreateActivity(activityEnt)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	activityResponse := response.Activity{
		Id:             activityEnt.Id,
		Title:          activityEnt.Title,
		ActivityStart:  activityEnt.ActivityStart,
		ActivityFinish: activityEnt.ActivityFinish,
		UserId:         activityEnt.UserId,
		ActivityDetail: response.ActivityDetail{
			HeartRate:      activityEnt.ActivityDetail.HeartRate,
			Intensity:      activityEnt.ActivityDetail.Intensity,
			CaloriesBurned: activityEnt.ActivityDetail.CaloriesBurned,
			FoodDetails:    activityEnt.ActivityDetail.FoodDetails,
			ImageUrl:       activityEnt.ActivityDetail.ImageUrl,
		},
		ActivityType: response.ActivityType{
			Name:        activityEnt.ActivityType.Name,
			Description: activityEnt.ActivityType.Description,
		},
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Register", activityResponse))
}