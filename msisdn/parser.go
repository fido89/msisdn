package msisdn

import (
	"errors"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"msisdn/data"
	"msisdn/swagger/models"
	"msisdn/swagger/restapi/operations"
	"strings"
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
	var carrier = data.GetCarrier(msisdn)
	if carrier == nil {
		return nil, errors.New(fmt.Sprintf("MSISDN %s can't be parsed", msisdn))
	}
	var isoCode = data.GetCountryIsoCode(carrier.CountryCode)
	var parsedMsisdn = &models.ParsedMsisdn{
		CountryCode:      carrier.CountryCode,
		CountryID:        *isoCode,
		MNOID:            carrier.CarrierMNO,
		SubscriberNumber: strings.TrimLeft(strings.TrimLeft(msisdn, carrier.CountryCode), carrier.CarrierMNO),
	}
	return parsedMsisdn, nil
}
