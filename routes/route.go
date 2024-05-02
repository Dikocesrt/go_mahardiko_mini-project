package routes

import (
	"habit/controllers/user"

	myMiddleware "habit/middlewares"

	"github.com/labstack/echo/v4"
)

type RouteController struct {
	userController *user.UserController
}

func NewRoute(userController *user.UserController) *RouteController {
	return &RouteController{
		userController: userController,
	}
}

func (r *RouteController) InitRoute(e *echo.Echo) {
	myMiddleware.LogMiddleware(e)
	e.POST("/register", r.userController.Register)
	e.POST("/login", r.userController.Login)
}