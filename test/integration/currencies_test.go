package integration

/*
import (
	"github.com/genert/pipedrive-api/pipedrive"
	"github.com/go-test/deep"
	"testing"
)

func TestCurrenciesService_List(t *testing.T) {
	currencies, _, err := client.Currencies.List(nil)

	if err != nil {
		t.Errorf("Could not get currencies: %v", err)
	}

	if currencies.Success != true {
		t.Error("Unsuccessful currencies request")
	}

	expectedCurrency := pipedrive.Currency{
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

func TestCurrenciesService_List2(t *testing.T) {
	currencies, _, err := client.Currencies.List(&pipedrive.CurrenciesListOptions{
		Term: "estonia",
	})

	if err != nil {
		t.Errorf("Could not get currencies: %v", err)
	}

	if currencies.Success != true {
		t.Error("Unsuccessful currencies request")
	}

	expectedCurrency := pipedrive.Currency{
		ID:            42,
		Code:          "EEK",
		Name:          "Estonian Kroon",
		DecimalPoints: 2,
		Symbol:        "",
		ActiveFlag:    true,
		IsCustomFlag:  false,
	}

	if diff := deep.Equal(expectedCurrency, currencies.Data[0]); diff != nil {
		t.Error(diff)
	}
}
*/
