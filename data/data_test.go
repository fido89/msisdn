package data

import (
	"reflect"
	"testing"
)

func TestGetCarrierAndCountryCodeLineShouldBeIgnored(t *testing.T) {
	carrierData, countryCode := getCarrierAndCountryCode("# this is a test", "file.txt")
	if carrierData != nil || countryCode != nil {
		t.Error("Line should be ignored")
	}
}

func TestGetCarrierAndCountryCodeLineShouldBeProcessed(t *testing.T) {
	carrierData, countryCode := getCarrierAndCountryCode("38631|Telekom Slovenije", "386.txt")
	if carrierData == nil || countryCode == nil {
		t.Error("Line should be processed")
	}
	expected := "38631"
	if *carrierData != expected {
		t.Errorf("Expected %s, got %s", expected, *carrierData)
	}
}

func TestGetCarrier(t *testing.T) {
	carriersData := make(map[string]string)
	carriersData["38631"] = "386"
	data := Data{carriers: carriersData, carrierMaxLenght: 5}

	expected := &Carrier{
		CarrierMNO: "31", CountryCode: "386",
	}
	actual := data.GetCarrier("38631607903")

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestGetCarrierNilExpected(t *testing.T) {
	data := Data{carriers: make(map[string]string), carrierMaxLenght: 0}
	actual := data.GetCarrier("38631607903")
	if actual != nil {
		t.Errorf("Expected to be nil")
	}
}

func TestGetCountryIsoCode(t *testing.T) {
	countryIsoCodes := make(map[string]string)
	countryIsoCodes["386"] = "SI"
	data := Data{
		countryIsoCodes: countryIsoCodes,
	}

	expected := "SI"
	actual := data.GetCountryIsoCode("386")
	if actual == nil {
		t.Errorf("Expected %s, got nil", expected)
	}

	if *actual != expected {
		t.Errorf("Expected %s, got %s", expected, *actual)
	}
}

func TestGetCountryIsoCodeNilExpected(t *testing.T) {
	data := Data{
		countryIsoCodes: make(map[string]string),
	}
	actual := data.GetCountryIsoCode("386")
	if actual != nil {
		t.Errorf("Expected to be nil")
	}
}
