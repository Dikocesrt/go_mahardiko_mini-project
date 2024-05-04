package routes

import (
	"habit/constants"
	"habit/controllers/activity"
	"habit/controllers/user"

	myMiddleware "habit/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RouteController struct {
	userController *user.UserController
	activityController *activity.ActivityController
}

func NewRoute(userController *user.UserController, activityController *activity.ActivityController) *RouteController {
	return &RouteController{
		userController: userController,
		activityController: activityController,
	}
}

func (r *RouteController) InitRoute(e *echo.Echo) {
	myMiddleware.LogMiddleware(e)
	e.POST("/register", r.userController.Register)
	e.POST("/login", r.userController.Login)

	eJwt := e.Group("/")
	eJwt.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	eJwt.POST("activities", r.activityController.CreateActivity)
	eJwt.GET("activities/:id", r.activityController.GetActivityById)
	eJwt.PUT("activities/:id", r.activityController.UpdateActivityById)
	eJwt.DELETE("activities/:id", r.activityController.DeleteActivityById)

	eJwt.GET("activities/user/:userId", r.activityController.GetActivityByUserId)
}