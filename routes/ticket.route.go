package routes

import (
	"OneTix/constants"
	"OneTix/controllers"
	"OneTix/middlewares"
	"OneTix/validators"

	"github.com/gin-gonic/gin"
)

func TicketRoutes(router *gin.Engine) {
	ticketGroup := router.Group("/ticket")
	{
		ticketGroup.GET("/", middlewares.RoleAuth(constants.CUSTOMER), middlewares.Validator[validators.FindTicketsQuery]("query"), controllers.FindTickets)
		ticketGroup.POST("/", middlewares.RoleAuth(constants.CUSTOMER), middlewares.Validator[validators.CreateTicketBody]("body"), controllers.CreateTicket)
		ticketGroup.POST("/check-in", middlewares.RoleAuth(constants.CUSTOMER), middlewares.Validator[validators.CheckInBody]("body"), controllers.CheckIn)
	}
}
