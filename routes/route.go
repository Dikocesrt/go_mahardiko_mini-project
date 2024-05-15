package routes

import (
	"habit/constants"
	"habit/controllers/activity"
	"habit/controllers/admin"
	"habit/controllers/chatbot"
	"habit/controllers/expert"
	"habit/controllers/hire"
	"habit/controllers/user"

	myMiddleware "habit/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RouteController struct {
	userController *user.UserController
	activityController *activity.ActivityController
	expertController *expert.ExpertController
	hireController *hire.HireController
	adminController *admin.AdminController
	chatbotController *chatbot.ChatbotController
}

func NewRoute(userController *user.UserController, activityController *activity.ActivityController, expertController *expert.ExpertController, hireController *hire.HireController, adminController *admin.AdminController, chatbotController *chatbot.ChatbotController) *RouteController {
	return &RouteController{
		userController: userController,
		activityController: activityController,
		expertController: expertController,
		hireController: hireController,
		adminController: adminController,
		chatbotController: chatbotController,
	}
}

func (r *RouteController) InitRoute(e *echo.Echo) {
	myMiddleware.LogMiddleware(e)

	//users auth
	e.POST("/users/register", r.userController.Register) //Register User
	e.POST("/users/login", r.userController.Login) //Login User

	//experts auth
	e.POST("/experts/register", r.expertController.Register) //Register Expert
	e.POST("/experts/login", r.expertController.Login) //Login Expert

	//admin auth
	e.POST("/admin/register", r.adminController.Register) //Register Admin
	e.POST("/admin/login", r.adminController.Login) //Login Admin

	userGroup := e.Group("/users")
	userGroup.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	userGroup.Use(myMiddleware.UserOnlyMiddleware)

	userGroup.GET("/:id", r.userController.GetUserById) //Get User By Id
	userGroup.PUT("/:id", r.userController.UpdateProfileById) //Update Profile By Id

	userGroup.POST("/activities", r.activityController.CreateActivity) //Create Activity
	userGroup.GET("/activities/:id", r.activityController.GetActivityById) //Get Activity By Id
	userGroup.PUT("/activities/:id", r.activityController.UpdateActivityById) //Update Activity By Id
	userGroup.DELETE("/activities/:id", r.activityController.DeleteActivityById) //Delete Activity By Id
	userGroup.GET("/activities/user/:id", r.activityController.GetActivityByUserId) //Get Activity By User Id

	userGroup.GET("/experts", r.expertController.GetAllExperts) //Get All Experts
	userGroup.GET("/experts/:id", r.expertController.GetExpertById) //Get Expert By Expert Id

	userGroup.POST("/hires", r.hireController.CreateHire) //Create Hire
	userGroup.GET("/hires/user/:id", r.hireController.GetHiresByUserId) //Get Hires By User Id
	userGroup.GET("/hires/:id", r.hireController.GetHireById) //Get Hire By Id

	userGroup.POST("/chatbot", r.chatbotController.Chat)





	expertGroup := e.Group("/experts")
	expertGroup.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	expertGroup.Use(myMiddleware.ExpertOnlyMiddleware)

	expertGroup.GET("/:id", r.expertController.GetExpertById) //Get Expert By Id
	expertGroup.PUT("/:id", r.expertController.UpdateProfileExpertById) //Update Profile By Expert Id

	expertGroup.GET("/hires/expert/:id", r.hireController.GetHiresByExpertId) //Get Hires By Expert Id
	expertGroup.GET("/hires/:id", r.hireController.GetHireById) //Get Hire By Id
	expertGroup.PUT("/hires/verify/:id", r.hireController.VerifyPayment) //Verify Payment

	expertGroup.GET("/users/:id", r.userController.GetUserById) //Get User By User Id

	expertGroup.GET("/activities/user/:id", r.activityController.GetActivityByUserId) //Get Activity By User Id
	expertGroup.GET("/activities/:id", r.activityController.GetActivityById) //Get Activity By Id




	adminGroup := e.Group("/admin")
	adminGroup.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	adminGroup.Use(myMiddleware.AdminOnlyMiddleware)

	adminGroup.POST("/bank-account-types", r.expertController.CreateBankAccountType) //Create Bank Account Type
	adminGroup.GET("/bank-account-types", r.expertController.GetAllBankAccountType) //Get All Bank Account Type
	adminGroup.GET("/bank-account-types/:id", r.expertController.GetBankAccountTypeById) //Get Bank Account Type By Id
	adminGroup.PUT("/bank-account-types/:id", r.expertController.UpdateBankAccountTypeById) //Update Bank Account Type By Id
	adminGroup.DELETE("/bank-account-types/:id", r.expertController.DeleteBankAccountTypeById) //Delete Bank Account Type By Id

	adminGroup.POST("/expertises", r.expertController.CreateExpertise) //Create Expertise
	adminGroup.GET("/expertises", r.expertController.GetAllExpertise) //Get All Expertise
	adminGroup.GET("/expertise/:id", r.expertController.GetExpertiseById) //Get Expertise By Id
	adminGroup.PUT("/expertise/:id", r.expertController.UpdateExpertiseById) //Update Expertise By Id
	adminGroup.DELETE("/expertise/:id", r.expertController.DeleteExpertiseById) //Delete Expertise By Id

	adminGroup.POST("/activity-types", r.activityController.CreateActivityType) //Create Activity Type
	adminGroup.GET("/activity-types", r.activityController.GetAllActivityType) //Get All Activity Type
	adminGroup.GET("/activity-types/:id", r.activityController.GetActivityTypeById) //Get Activity Type By Id
	adminGroup.PUT("/activity-types/:id", r.activityController.UpdateActivityTypeById) //Update Activity Type By Id
	adminGroup.DELETE("/activity-types/:id", r.activityController.DeleteActivityTypeById) //Delete Activity Type By Id
}