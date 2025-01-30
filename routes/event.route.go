package routes

import (
	"OneTix/constants"
	"OneTix/controllers"
	"OneTix/middlewares"
	"OneTix/validators"

	"github.com/gin-gonic/gin"
)

func EventRoutes(router *gin.Engine) {
	eventGroup := router.Group("/event")
	{
		eventGroup.GET("/", middlewares.RoleAuth(constants.ORGANIZER, constants.CUSTOMER), middlewares.Validator[validators.FindEventsQuery]("query"), controllers.FindEvents)
		eventGroup.POST("/", middlewares.RoleAuth(constants.ORGANIZER), middlewares.Validator[validators.CreateEventBody]("body"), controllers.CreateEvent)
		eventGroup.DELETE("/", middlewares.RoleAuth(constants.ORGANIZER), middlewares.Validator[validators.DeleteEventQuery]("query"), controllers.DeleteEvent)
		eventGroup.PUT("/", middlewares.RoleAuth(constants.ORGANIZER), middlewares.Validator[validators.UpdateEventQuery]("query"), middlewares.Validator[validators.UpdateEventBody]("body"), controllers.UpdateEvent)
	}
}
