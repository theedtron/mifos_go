package controllers

import (
	"pesanode/gobackend/utils"
	"time"

	"github.com/rs/zerolog/log"
)

type SuccessResponse struct {
	Status string
	Data []string
}

type MifosClientSearchString struct {
	QueryString string
}

func mifosFormatDate (f_date time.Time) string {
	const (
		layoutISO = "2006-01-02"
		layoutUS  = "2 January 2006"
	)

	return f_date.Format(layoutUS)
}


func mifosResponse (resp[] string) SuccessResponse {
	if !isElementExist(resp,"developerMessage") && !isElementExist(resp,"timestamp") {

		res := SuccessResponse{
			Status: "success",
			Data: resp,
		}

		return res
	}else{
		log.Error().Msg("Mifos error check API logs for further info")
		res := SuccessResponse{
			Status: "error",
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
	payload := utils.MifosPayloadGet{
		UrlExt: "search?query=" + queryString.QueryString + "&resource=clients&exactMatch=true",
	}

	result := utils.MifosGet(&payload)
	return mifosResponse(string(result))

	
}