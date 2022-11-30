package controllers

import (
	"encoding/json"
	"pesanode/gobackend/mifos"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
)

type SuccessResponse struct {
	Status string
	Data map[string]interface{}
}

type MifosClientSearchString struct {
	QueryString string
}

type MifosClientSearchId struct {
	ClientId string
}

type MifosAcTransferResourceId struct {
	ResourceId int
}

func mifosFormatDate (f_date time.Time) string {
	const (
		layoutISO = "2006-01-02"
		layoutUS  = "2 January 2006"
	)

	return f_date.Format(layoutUS)
}


func mifosResponse (resp map[string]interface{}) SuccessResponse {

	if  value, isMapContainsKey := resp["developerMessage"]; isMapContainsKey {

		log.Error().Msg("Mifos error check API logs for further info"+ value.(string))
		res := SuccessResponse{
			Status: "error",
			Data: resp,
		}

		return res
	}else if  value, isMapContainsKey := resp["timestamp"]; isMapContainsKey {

		log.Error().Msg("Mifos error check API logs for further info"+ value.(string))
		res := SuccessResponse{
			Status: "error",
			Data: resp,
		}

		return res
	}else{
		res := SuccessResponse{
			Status: "success",
			Data: resp,
		}

		return res
	}

}

func isElementExist(s []string, str string) bool {
    for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func fetchClientBySearch(queryString *MifosClientSearchString) SuccessResponse {
	payload := mifos.MifosPayloadGet{
		UrlExt: "/search?query=" + queryString.QueryString + "&resource=clients&exactMatch=true",
	}

	result := mifos.MifosGet(&payload)
	var res_format map[string]interface{}
	err := json.Unmarshal(result, res_format)
	if err == nil {
		log.Error().Err(err).
			Msg("Error occurred processing mifos response")
	}

	return mifosResponse(res_format)
	
}

func fetchClientById(id *MifosClientSearchId) SuccessResponse {
	payload := mifos.MifosPayloadGet{
		UrlExt: "/clients/" + id.ClientId,
	}

	client_result := mifos.MifosGet(&payload)
	var client_format map[string]interface{}
	err := json.Unmarshal(client_result, &client_format)
	if err == nil {
		log.Error().Err(err).
			Msg("Error occurred processing mifos client response")
	}

	return mifosResponse(client_format)
}

func fetchAcTransferById(id *MifosAcTransferResourceId) SuccessResponse {
	payload := mifos.MifosPayloadGet{
		UrlExt: "/accounttransfers/" + strconv.Itoa(id.ResourceId),
	}

	result := mifos.MifosGet(&payload)
	var res_format map[string]interface{}
	err := json.Unmarshal(result, &res_format)
	if err == nil {
		log.Error().Err(err).
			Msg("Error occurred processing mifos response")
	}

	return mifosResponse(res_format)
	
}