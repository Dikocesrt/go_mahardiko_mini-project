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
	e.POST("/users/register", r.userController.Register) //Register User
	e.POST("/users/login", r.userController.Login) //Login User

	e.POST("/experts/register", r.expertController.Register) //Register Expert
	e.POST("/experts/login", r.expertController.Login) //Login Expert

	userGroup := e.Group("/")
	userGroup.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	userGroup.Use(myMiddleware.UserOnlyMiddleware)
	userGroup.POST("activities", r.activityController.CreateActivity) //Create Activity
	userGroup.GET("activities/:id", r.activityController.GetActivityById) //Get Activity By Id
	userGroup.PUT("activities/:id", r.activityController.UpdateActivityById) //Update Activity By Id
	userGroup.DELETE("activities/:id", r.activityController.DeleteActivityById) //Delete Activity By Id
	userGroup.GET("activities/users/:userId", r.activityController.GetActivityByUserId) //Get Activity By User Id
	userGroup.PUT("users/:id", r.userController.UpdateProfileById) //Update Profile By Id

	expertGroup := e.Group("/")
	expertGroup.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	expertGroup.Use(myMiddleware.ExpertOnlyMiddleware)
	expertGroup.PUT("experts/:id", r.expertController.UpdateProfileExpertById) //Update Profile By Expert Id
}