package controllers

import (
	"net/http"
	"pesanode/gobackend/utils"
	"github.com/gin-gonic/gin"
)

func SendTrxOtp(c *gin.Context){
	
	// Send Email
	emailData := utils.EmailData{
		URL:       utils.GetEnvVar("") + "/verifyemail/",
		FirstName: "Theed",
		Subject:   "Your account verification code",
		MailTo: "kareo@devs.mobi",
	}

	utils.SendEmail(&emailData)

	message := "We sent an email with a verification code to your email"
	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": message})
}