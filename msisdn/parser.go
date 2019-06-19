package msisdn

import (
	"errors"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"msisdn/data"
	"msisdn/swagger/models"
	"msisdn/swagger/restapi/operations"
	"regexp"
	"strings"
)

func ParseMsisdn(params operations.ParseMsisdnParams) middleware.Responder {
	var msisdn = getClearMsisdn(params.Msisdn)

	parsedMsisdn, err := getParsedMsisdn(data.GetInstance(), msisdn)
	if err != nil {
		return operations.NewParseMsisdnNotFound().WithPayload(
			&models.NotFound{
				int64(operations.ParseMsisdnNotFoundCode),
				swag.String(fmt.Sprintf("%s", err)),
			})
	}
	return operations.NewParseMsisdnOK().WithPayload(parsedMsisdn)
}

func getParsedMsisdn(dataStruct data.DataGetter, msisdn string) (*models.ParsedMsisdn, error) {
	var carrier = dataStruct.GetCarrier(msisdn)
	if carrier == nil {
		return nil, errors.New(fmt.Sprintf("MSISDN %s can't be parsed", msisdn))
	}
	var isoCode = dataStruct.GetCountryIsoCode(carrier.CountryCode)
	var parsedMsisdn = &models.ParsedMsisdn{
		CountryCode:      carrier.CountryCode,
		CountryID:        *isoCode,
		MNOID:            carrier.CarrierMNO,
		SubscriberNumber: strings.TrimLeft(strings.TrimLeft(msisdn, carrier.CountryCode), carrier.CarrierMNO),
	}
	return parsedMsisdn, nil
}

func getClearMsisdn(msisdn string) string {
	// Get only numbers in msisdn strins
	re := regexp.MustCompile("[0-9]+")
	var clearMsisdnArray = re.FindAllString(msisdn, -1)
	var clearMsisdn = strings.Join(clearMsisdnArray, "")
	// Clear internation prefix
	clearMsisdn = strings.TrimLeft(clearMsisdn, "0")

	return clearMsisdn
}
