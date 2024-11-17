package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRecipes(c *gin.Context) {
	c.JSON(http.StatusOK, []string{})
}

