package routes

import (
	"habit/constants"
	"habit/controllers/activity"
	"habit/controllers/expert"
	"habit/controllers/user"

	myMiddleware "habit/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RouteController struct {
	userController *user.UserController
	activityController *activity.ActivityController
	expertController *expert.ExpertController
}

func NewRoute(userController *user.UserController, activityController *activity.ActivityController, expertController *expert.ExpertController) *RouteController {
	return &RouteController{
		userController: userController,
		activityController: activityController,
		expertController: expertController,
	}
}

func (r *RouteController) InitRoute(e *echo.Echo) {
	myMiddleware.LogMiddleware(e)
	e.POST("/users/register", r.userController.Register)
	e.POST("/users/login", r.userController.Login)

	e.POST("/experts/register", r.expertController.Register)

	eJwt := e.Group("/")
	eJwt.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	eJwt.POST("activities", r.activityController.CreateActivity)
	eJwt.GET("activities/:id", r.activityController.GetActivityById)
	eJwt.PUT("activities/:id", r.activityController.UpdateActivityById)
	eJwt.DELETE("activities/:id", r.activityController.DeleteActivityById)

	eJwt.GET("activities/users/:userId", r.activityController.GetActivityByUserId)

	eJwt.PUT("users/:id", r.userController.UpdateProfileById)
}