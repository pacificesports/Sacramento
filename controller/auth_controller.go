package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sacramento/service"
)

func LoginWithCustomID(c *gin.Context) {
	result := service.CreateCustomToken(c.Param("uid"))
	if result == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to login user with given id: " + c.Param("uid")})
	} else {
		c.JSON(http.StatusOK, gin.H{"token": result})
	}
}
