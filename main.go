package main

import (
	"authentication/config"
	"authentication/helpers"
	"authentication/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	key := config.GenerateRandomKey()
	helpers.SetJWTKey(key)

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":" + port)
	log.Println("Server is running on port:", port)
}
