package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pesanode/gobackend/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type MifosPayloadPost struct {
	url_extension string
	clientBody []string
}

type MifosPayloadGet struct {
	UrlExt string
}

var url = GetEnvVar("MIFOS_URL")

func mifosPost(clientBody []string, c *gin.Context) {
	cred := GetEnvVar("MIFOS_UN")+GetEnvVar("MIFOS_PASS")
    baseCred := base64.StdEncoding.EncodeToString([]byte(cred))
	formatData, err := json.Marshal(clientBody)
	if err != nil {
		log.Error().Msg("Error occurred while binding request data")
	}
	data := []byte(formatData)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Error().Msg("Error occurred while binding request data")
	}

	// Set headers
	req.Header.Set("Authorization", baseCred)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Fineract-Platform-TenantId", GetEnvVar("MIFOS_TENANT"))

	// Set client timeout
	client := &http.Client{Timeout: time.Second * 10}

	// Validate cookie and headers are attached
	fmt.Println(req.Header)

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		log.Error().Msg("Error reading response.")
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error().Msg("Error reading body. ")
	}

	var apiLog models.ApiLog
	apiLog.FillDefaults()

	db, conErr := GetDatabaseConnection()
	if conErr != nil {
		log.Err(conErr).Msg("Error occurred while getting a DB connection from the connection pool")
		fmt.Printf("Service unavailable")
		return
	}

	// Create a user
	apiLogData := models.ApiLog{
		RequestUrl: url,
		RequestType: "POST",
		RequestBody: string(formatData),
		ResponseBody: string(body),
	}
	result := db.Create(apiLogData)
	if result.Error != nil && result.RowsAffected != 1 {
		log.Err(result.Error).Msg("Error occurred while creating a new user")
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error occurred while creating a new user",
		})
		return
	}

	fmt.Printf("%s\n", body)
}

func MifosGet(payload *MifosPayloadGet) []byte {
	cred := GetEnvVar("MIFOS_UN")+GetEnvVar("MIFOS_PASS")
    baseCred := base64.StdEncoding.EncodeToString([]byte(cred))

	ext := payload.UrlExt
	req, err := http.NewRequest("GET", url+ext, nil)
	if err != nil {
		log.Error().Msg("Error occurred while binding request data")
	}

	// Set headers
	req.Header.Set("Authorization", baseCred)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Fineract-Platform-TenantId", GetEnvVar("MIFOS_TENANT"))

	// Set client timeout
	client := &http.Client{Timeout: time.Second * 10}

	// Validate cookie and headers are attached
	fmt.Println(req.Header)

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		log.Error().Msg("Error reading response.")
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error().Msg("Error reading body. ")
	}

	fmt.Printf("%s\n", body)

	return body
}

