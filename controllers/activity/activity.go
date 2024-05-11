package activity

import (
	"habit/controllers/activity/request"
	"habit/controllers/activity/response"
	activityEntities "habit/entities/activity"
	"habit/utilities/base"
	"net/http"
	"strconv"

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

	file, _ := c.FormFile("food_images")

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

	activityEnt, err := activityController.activityUseCase.CreateActivity(activityEnt, file)
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

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Create Activity", activityResponse))
}

func (activityController *ActivityController) GetActivityByUserId(c echo.Context) error {
	userId := c.Param("userId")
	id, _ := strconv.Atoi(userId)

	activityEnt, err := activityController.activityUseCase.GetActivityByUserId(id)

	activityResponse := make([]response.Activity, len(activityEnt))
	for i := 0; i < len(activityEnt); i++ {
		activityResponse[i] = response.Activity{
			Id:             activityEnt[i].Id,
			Title:          activityEnt[i].Title,
			ActivityStart:  activityEnt[i].ActivityStart,
			ActivityFinish: activityEnt[i].ActivityFinish,
			UserId:         activityEnt[i].UserId,
			ActivityDetail: response.ActivityDetail{
				HeartRate:      activityEnt[i].ActivityDetail.HeartRate,
				Intensity:      activityEnt[i].ActivityDetail.Intensity,
				CaloriesBurned: activityEnt[i].ActivityDetail.CaloriesBurned,
				FoodDetails:    activityEnt[i].ActivityDetail.FoodDetails,
				ImageUrl:       activityEnt[i].ActivityDetail.ImageUrl,
			},
			ActivityType: response.ActivityType{
				Name:        activityEnt[i].ActivityType.Name,
				Description: activityEnt[i].ActivityType.Description,
			},
		}
	}

	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get Activity", activityResponse))
}

func (activityController *ActivityController) GetActivityById(c echo.Context) error {
	activityId := c.Param("id")
	id, _ := strconv.Atoi(activityId)

	var activityEnt activityEntities.Activity
	activityEnt.Id = id
	activityEnt, err := activityController.activityUseCase.GetActivityById(activityEnt)
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

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get Activity", activityResponse))
}

func (activityController *ActivityController) UpdateActivityById(c echo.Context) error {
	activityId := c.Param("id")
	id, _ := strconv.Atoi(activityId)

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

	activityEnt.Id = id

	activityEnt, err := activityController.activityUseCase.UpdateActivityById(activityEnt)
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

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Update Activity", activityResponse))
}

func (activityController *ActivityController) DeleteActivityById(c echo.Context) error {
	activityId := c.Param("id")
	id, _ := strconv.Atoi(activityId)

	var activityEnt activityEntities.Activity
	activityEnt.Id = id

	err := activityController.activityUseCase.DeleteActivityById(activityEnt)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Delete Activity", nil))
}