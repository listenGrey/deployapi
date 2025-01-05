package controller

import (
	"deployapi/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Handler(c *gin.Context) {
	var req model.Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Process the input string
	response := model.Response{
		Output: strings.ToUpper(req.Input) + " Already processed.",
	}

	// Send JSON response
	c.JSON(http.StatusOK, response)
}
