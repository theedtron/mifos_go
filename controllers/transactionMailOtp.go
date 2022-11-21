package controllers

import (
	"net/http"
	"pesanode/gobackend/utils"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type mifosTrx struct {
	savingsId int
	resourceId int
}

func SendTrxOtp(c *gin.Context){
	
	request_id := c.GetString("x-request-id")
	payload := &mifosTrx{}

	// Bind request payload with our model
	if binderr := c.ShouldBindJSON(payload); binderr != nil {

		log.Error().Err(binderr).Str("request_id", request_id).
			Msg("Error occurred while binding request data")

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": binderr.Error(),
		})
		return
	}
	
	// Send Email
	emailData := utils.EmailData{
		URL: utils.GetEnvVar("") + "/verifyemail/",
		FirstName: "Theed",
		Subject: "Your account verification code",
		MailTo: "kareo@devs.mobi",
	}

	utils.SendEmail(&emailData)

	message := "We sent an email with a verification code to your email"
	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": message})
}