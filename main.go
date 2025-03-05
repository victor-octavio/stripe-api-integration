package main

import (
	"back_end_mirumuh/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.InitRoutes(r.Group(""))

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
