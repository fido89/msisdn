package msisdn

import (
	"msisdn/data"
	"msisdn/swagger/models"
	"reflect"
	"testing"
)

func TestGetClearMsisdn(t *testing.T) {
	expected := "38640607903"
	actual := getClearMsisdn("+(386)40607903")
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

	actual = getClearMsisdn("00386 40 607-903")
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

type FakeData struct {
	carrier        data.Carrier
	countryIsoCode string
}

func (d *FakeData) GetCountryIsoCode(string) *string {
	return &d.countryIsoCode
}

func (d *FakeData) GetCarrier(msisdn string) *data.Carrier {
	return &d.carrier
}

func TestGetParsedMsisdn(t *testing.T) {
	f := FakeData{
		carrier:        data.Carrier{"386", "40"},
		countryIsoCode: "SI",
	}
	expected := &models.ParsedMsisdn{
		CountryCode: "386", CountryID: "SI", MNOID: "40", SubscriberNumber: "607903",
	}
	actual, returnedError := getParsedMsisdn(&f, "38640607903")
	if returnedError != nil {
		t.Error("Error isn't expected")
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestGetParsedMsisdnWithError(t *testing.T) {
	f := FakeData{}
	_, returnedError := getParsedMsisdn(&f, "38640607903")
	if returnedError != nil {
		t.Error("Error is expected")
	}

}
