package controller

import (
	"deployapi/dao"
	"deployapi/model"
	"github.com/gin-gonic/gin"
	"log"
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
	response := model.Response{}

	output := strings.ToUpper(req.Input) + " Already processed."
	response.Output = output

	err := dao.Send(output)
	if err != nil {
		response.Output = "grpc wrong"
		log.Fatalf("grpc error: %v", err)
	}

	// Send JSON response
	c.JSON(http.StatusOK, response)
}

func Hello(c *gin.Context) {
	req := c.Param("msg")

	// Process the input string
	response := model.Response{}

	output := strings.ToUpper(req) + " Already processed."
	response.Output = output

	err := dao.Send(output)
	if err != nil {
		response.Output = "grpc wrong"
		log.Fatalf("grpc error: %v", err)
	}

	// Send JSON response
	c.JSON(http.StatusOK, response)
}
