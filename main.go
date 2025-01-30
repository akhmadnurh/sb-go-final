package main

import (
	"OneTix/configs"
	"OneTix/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.ConnectDatabase()

	env := configs.LoadEnv()

	// Menampilkan konfigurasi untuk debugging
	log.Printf("Loaded config: %+v\n", env)

	router := gin.Default()

	routes.AuthRoutes(router)

	router.Run(fmt.Sprintf(":%s", env.AppPort))

}
