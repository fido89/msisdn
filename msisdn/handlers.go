package msisdn

import (
	"github.com/go-openapi/runtime/middleware"
	"msisdn/swagger/models"
	"msisdn/swagger/restapi/operations"
)

func ParseMsisdn(params operations.ParseMsisdnParams) middleware.Responder {
	var msisdn string = params.Msisdn
	return operations.NewParseMsisdnOK().WithPayload(
		&models.ParsedMsisdn{
			CountryCode:      msisdn,
			CountryID:        msisdn,
			MNOID:            msisdn,
			SubscriberNumber: msisdn,
		})
}
