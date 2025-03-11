package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeController(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", nil)
}
