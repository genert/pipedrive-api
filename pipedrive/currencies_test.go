package pipedrive

import (
	"github.com/go-test/deep"
	"testing"
)

func TestCurrenciesService_List(t *testing.T) {
	currencies, _, err := client.Currencies.List()

	if err != nil {
		t.Error("Could not get currencies: %v", err)
	}

	if currencies.Success != true {
		t.Error("Unsuccessful currencies request")
	}

	expectedCurrency := Currency{
		ID:            2,
		Code:          "AFN",
		Name:          "Afghanistan Afghani",
		DecimalPoints: 2,
		Symbol:        "AFN",
		ActiveFlag:    true,
		IsCustomFlag:  false,
	}

	if diff := deep.Equal(expectedCurrency, currencies.Data[0]); diff != nil {
		t.Error(diff)
	}
}
