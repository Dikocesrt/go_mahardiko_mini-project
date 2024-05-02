package main

import (
	"habit/configs"
	userController "habit/controllers/user"
	"habit/repositories/mysql"
	userRepositories "habit/repositories/mysql/user"
	"habit/routes"
	userUseCase "habit/usecases/user"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.LoadEnv()
	db := mysql.ConnectDB(configs.InitConfigMySQL())
	
	userRepo := userRepositories.NewUserRepo(db)
	userUC := userUseCase.NewUserUseCase(userRepo)
	userCont := userController.NewUserController(userUC)

	route := routes.NewRoute(userCont)

	e := echo.New()
	route.InitRoute(e)
	e.Logger.Fatal(e.Start(":8080"))
}