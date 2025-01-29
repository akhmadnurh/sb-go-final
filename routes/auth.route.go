package routes

import (
	"OneTix/controllers"
	"OneTix/middlewares"
	"OneTix/validators"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register/customer", middlewares.Validator[validators.RegisterCustomerBody]("body"), controllers.Register)
	}

}
