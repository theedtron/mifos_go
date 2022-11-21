package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	t := time.Now()
	c.JSON(http.StatusOK, gin.H{
		"message": mifosFormatDate(t),
	})
}
