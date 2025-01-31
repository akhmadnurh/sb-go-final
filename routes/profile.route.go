package routes

import (
	"OneTix/constants"
	"OneTix/controllers"
	"OneTix/middlewares"
	"OneTix/validators"

	"github.com/gin-gonic/gin"
)

func ProfileRoutes(router *gin.Engine) {
	profileGroup := router.Group("/profile")
	{
		profileGroup.GET("/customer", middlewares.RoleAuth(constants.CUSTOMER), controllers.GetProfile)
		profileGroup.GET("/organizer", middlewares.RoleAuth(constants.ORGANIZER), controllers.GetOrganizerProfile)
		profileGroup.PUT("/customer", middlewares.RoleAuth(constants.CUSTOMER), middlewares.Validator[validators.UpdateProfileBody]("body"), controllers.UpdateProfile)
		profileGroup.PUT("/organizer", middlewares.RoleAuth(constants.ORGANIZER), middlewares.Validator[validators.UpdateOrganizerProfileBody]("body"), controllers.UpdateOrganizerProfile)
		profileGroup.PATCH("/password", middlewares.RoleAuth(constants.CUSTOMER, constants.ORGANIZER), middlewares.Validator[validators.UpdatePasswordBody]("body"), controllers.UpdatePassword)
	}
}
