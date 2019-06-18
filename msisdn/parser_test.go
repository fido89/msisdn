package msisdn

import "testing"

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
