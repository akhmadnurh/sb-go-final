package main

import (
	"OneTix/configs"
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

	router.Run(fmt.Sprintf(":%s", env.AppPort))

}
