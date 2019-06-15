package msisdn

import (
	"errors"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"msisdn/swagger/models"
	"msisdn/swagger/restapi/operations"
)

func ParseMsisdn(params operations.ParseMsisdnParams) middleware.Responder {
	var msisdn string = params.Msisdn

	parsedMsisdn, err := getParsedMsisdn(msisdn)
	if err != nil {
		fmt.Sprintf("%s", err)
		return operations.NewParseMsisdnNotFound().WithPayload(
			&models.NotFound{
				int64(operations.ParseMsisdnNotFoundCode),
				swag.String(fmt.Sprintf("%s", err)),
			})
	}
	return operations.NewParseMsisdnOK().WithPayload(parsedMsisdn)
}

func getParsedMsisdn(msisdn string) (*models.ParsedMsisdn, error) {
	if msisdn == "error" {
		return nil, errors.New(fmt.Sprintf("MSISDN %s can't be parsed", msisdn))
	}

	var test = &models.ParsedMsisdn{
		CountryCode:      msisdn,
		CountryID:        msisdn,
		MNOID:            msisdn,
		SubscriberNumber: msisdn,
	}
	return test, nil
}
