package main

import (
	"cinema/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("views/*")
	r.Static("/static", "./static")

	routes.SetupRoutes(r)

	r.Run(":8000")
}
