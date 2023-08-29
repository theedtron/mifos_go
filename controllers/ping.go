package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Ping(c *gin.Context) {
	t := time.Now()
	c.JSON(http.StatusOK, gin.H{
		"message": mifosFormatDate(t),
	})
}

type MifosUserSearch struct {
	Phone string `json:"phone" binding:"required"`
}

func SearchMifosUser(c *gin.Context){
	request_id := c.GetString("x-request-id")
	hookData := MifosUserSearch{}

	// Bind request payload with our model
	if binderr := c.ShouldBindJSON(&hookData); binderr != nil {

		log.Error().Err(binderr).Str("request_id", request_id).
			Msg("Error occurred while binding request data")

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": binderr.Error(),
		})
		return
	}

	searchPhone := MifosClientSearchString{
		QueryString: hookData.Phone,
	}

	res := fetchClientBySearch(&searchPhone)
	if res.Status != "success" {
		log.Error().Msg("Error occurred")
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data": res,
	})
}
