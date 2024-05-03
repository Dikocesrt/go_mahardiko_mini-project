package main

import (
	"habit/configs"
	activityController "habit/controllers/activity"
	userController "habit/controllers/user"
	"habit/repositories/mysql"
	activityRepositories "habit/repositories/mysql/activity"
	userRepositories "habit/repositories/mysql/user"
	"habit/routes"
	activityUseCase "habit/usecases/activity"
	userUseCase "habit/usecases/user"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.LoadEnv()
	db := mysql.ConnectDB(configs.InitConfigMySQL())
	
	userRepo := userRepositories.NewUserRepo(db)
	userUC := userUseCase.NewUserUseCase(userRepo)
	userCont := userController.NewUserController(userUC)

	activityRepo := activityRepositories.NewActivityRepo(db)
	activityUC := activityUseCase.NewActivityUseCase(activityRepo)
	activityCont := activityController.NewActivityController(activityUC)

	route := routes.NewRoute(userCont, activityCont)

	e := echo.New()
	route.InitRoute(e)
	e.Logger.Fatal(e.Start(":8080"))
}