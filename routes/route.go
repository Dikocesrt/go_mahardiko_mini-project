package routes

import (
	"habit/constants"
	"habit/controllers/activity"
	"habit/controllers/admin"
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
}

func NewRoute(userController *user.UserController, activityController *activity.ActivityController, expertController *expert.ExpertController, hireController *hire.HireController, adminController *admin.AdminController) *RouteController {
	return &RouteController{
		userController: userController,
		activityController: activityController,
		expertController: expertController,
		hireController: hireController,
		adminController: adminController,
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
	userGroup.POST("/activities", r.activityController.CreateActivity) //Create Activity
	userGroup.GET("/activities/:id", r.activityController.GetActivityById) //Get Activity By Id
	userGroup.PUT("/activities/:id", r.activityController.UpdateActivityById) //Update Activity By Id
	userGroup.DELETE("/activities/:id", r.activityController.DeleteActivityById) //Delete Activity By Id
	userGroup.GET("/activities/user/:userId", r.activityController.GetActivityByUserId) //Get Activity By User Id

	userGroup.GET("/:id", r.userController.GetUserById) //Get User By Id
	userGroup.PUT("/:id", r.userController.UpdateProfileById) //Update Profile By Id

	userGroup.GET("/experts", r.expertController.GetAllExperts) //Get All Experts
	userGroup.GET("/expert/:id", r.expertController.GetExpertById) //Get Expert By Expert Id

	userGroup.POST("/hires", r.hireController.CreateHire) //Create Hire



	expertGroup := e.Group("/experts")
	expertGroup.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	expertGroup.Use(myMiddleware.ExpertOnlyMiddleware)
	expertGroup.PUT("/:id", r.expertController.UpdateProfileExpertById) //Update Profile By Expert Id
	expertGroup.GET("/:id", r.expertController.GetExpertById) //Get Expert By Id

	expertGroup.GET("/user/:id", r.userController.GetUserById) //Get User By User Id

	expertGroup.GET("/activities/user/:userId", r.activityController.GetActivityByUserId) //Get Activity By User Id
	
	expertGroup.GET("/hires/expert/:expertId", r.hireController.GetHiresByExpertId) //Get Hires By Expert Id
	expertGroup.PUT("/hires/verify/:hireId", r.hireController.VerifyPayment) //Verify Payment



	adminGroup := e.Group("/admin")
	adminGroup.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	adminGroup.Use(myMiddleware.AdminOnlyMiddleware)
	adminGroup.POST("/bank-account-types", r.adminController.CreateBankAccountType) //Create Bank Account Type
	adminGroup.GET("/bank-account-types/:bankAccountTypeId", r.adminController.GetBankAccountTypeById) //Get Bank Account Type By Id
	adminGroup.GET("/bank-account-types", r.adminController.GetAllBankAccountType) //Get All Bank Account Type
	adminGroup.PUT("/bank-account-types/:bankAccountTypeId", r.adminController.UpdateBankAccountTypeById) //Update Bank Account Type By Id
	adminGroup.DELETE("/bank-account-types/:bankAccountTypeId", r.adminController.DeleteBankAccountTypeById) //Delete Bank Account Type By Id

	adminGroup.POST("/expertises", r.adminController.CreateExpertise) //Create Expertise
	adminGroup.GET("/expertises", r.adminController.GetAllExpertise) //Get All Expertise
	adminGroup.GET("/expertise/:expertiseId", r.adminController.GetExpertiseById) //Get Expertise By Id
	adminGroup.PUT("/expertise/:expertiseId", r.adminController.UpdateExpertiseById) //Update Expertise By Id
	adminGroup.DELETE("/expertise/:expertiseId", r.adminController.DeleteExpertiseById) //Delete Expertise By Id
}