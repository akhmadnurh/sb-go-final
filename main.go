package main

import (
	"OneTix/configs"
	"OneTix/routes"
	"OneTix/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.ConnectDatabase()

	env := configs.LoadEnv()

	// Menampilkan konfigurasi untuk debugging
	log.Printf("Loaded config: %+v\n", env)

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, utils.GenerateResponse(true, "Welcome to OneTix API"))
	})

	router.GET("docs/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, env.DocsURL)
	})

	routes.AuthRoutes(router)
	routes.EventRoutes(router)
	routes.ProfileRoutes(router)
	routes.TicketRoutes(router)

	router.Run(fmt.Sprintf(":%s", env.AppPort))

}
