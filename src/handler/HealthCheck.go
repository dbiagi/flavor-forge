package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheckResponse struct {
	Status string `json:"status"`
}

func HealthCheck(c *gin.Context) {
	health := HealthCheckResponse{Status: "UP"}
	c.JSON(http.StatusOK, health)
}

func HealthCheckComplete(c *gin.Context) {
	health := HealthCheckResponse{Status: "UP"}
	c.JSON(http.StatusOK, health)
}
