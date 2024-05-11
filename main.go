package main

import (
	"habit/configs"
	activityController "habit/controllers/activity"
	expertController "habit/controllers/expert"
	hireController "habit/controllers/hire"
	userController "habit/controllers/user"
	"habit/repositories/mysql"
	activityRepositories "habit/repositories/mysql/activity"
	expertRepositories "habit/repositories/mysql/expert"
	hireRepositories "habit/repositories/mysql/hire"
	userRepositories "habit/repositories/mysql/user"
	"habit/routes"
	activityUseCase "habit/usecases/activity"
	expertUseCase "habit/usecases/expert"
	hireUseCase "habit/usecases/hire"
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

	expertRepo := expertRepositories.NewExpertRepo(db)
	expertUC := expertUseCase.NewExpertUseCase(expertRepo)
	expertCont := expertController.NewExpertController(expertUC)

	hireRepo := hireRepositories.NewUserRepo(db)
	hireUC := hireUseCase.NewHireUseCase(hireRepo)
	hireCont := hireController.NewHireController(hireUC)

	route := routes.NewRoute(userCont, activityCont, expertCont, hireCont)

	e := echo.New()
	route.InitRoute(e)
	e.Logger.Fatal(e.Start(":8080"))
}