package controllers

import (
	"net/http"
	"pesanode/gobackend/utils"
	"strconv"

	// "strconv"
	// "pesanode/gobackend/utils"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type MifosTrx struct {
	SavingsId int `json:"savingsid" binding:"required"`
	ResourceId int `json:"resourceid" binding:"required"`
}

func SendTrxOtp(c *gin.Context){
	
	request_id := c.GetString("x-request-id")
	hookData := MifosTrx{}

	// Bind request payload with our model
	if binderr := c.ShouldBindJSON(&hookData); binderr != nil {

		log.Error().Err(binderr).Str("request_id", request_id).
			Msg("Error occurred while binding request data")

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": binderr.Error(),
		})
		return
	}

	// resource_int, err := strconv.Atoi(payload.resourceId)
	// if err != nil {
	// 	log.Error().Msg("Error occurred while binding request data")
	// }

	trx_id := MifosAcTransferResourceId{
		ResourceId: hookData.ResourceId,
	}

	mifos_trx_detail := fetchAcTransferById(&trx_id)
	if mifos_trx_detail.Status != "success" {
		log.Error().Msg("Error occurred while formating mifos response")
	}

	senderClientId := mifos_trx_detail.Data["fromClient"].(map[string]interface{})["id"]
	receiverClientId := mifos_trx_detail.Data["toClient"].(map[string]interface{})["id"]
	
	prepareMail(senderClientId.(float64))
	prepareMail(receiverClientId.(float64))
	

	message := "We sent an email with a verification code to your email"
	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": message})
}

func prepareMail(senderId float64) {
	senderMifosId := MifosClientSearchId{
		ClientId: strconv.FormatFloat(senderId,'f',0,64),
	}

	mifosClientDetail := fetchClientById(&senderMifosId)
	if mifosClientDetail.Status != "success" {
		log.Error().Msg("Error occurred while formating mifos response")
	}

	email := mifosClientDetail.Data["externalId"]
	
	// Send Email
	senderEmailData := utils.EmailData{
		URL: utils.GetEnvVar("") + "/verifyemail/",
		FirstName: "Theed",
		Subject: "Your account verification code",
		MailTo: email.(string),
	}

	utils.SendEmail(&senderEmailData)
}