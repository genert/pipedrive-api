package integration

import (
	"context"
	"testing"

	"github.com/genert/pipedrive-api/pipedrive"
	"github.com/go-test/deep"
)

func TestCurrenciesService_List(t *testing.T) {
	result, _, err := client.Currencies.List(context.Background(), nil)

	if err != nil {
		t.Errorf("Could not get currencies: %v", err)
	}

	if result.Success != true {
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

	if diff := deep.Equal(expectedCurrency, result.Data[0]); diff != nil {
		t.Error(diff)
	}
}

func TestCurrenciesService_List2(t *testing.T) {
	result, _, err := client.Currencies.List(context.Background(), &pipedrive.CurrenciesListOptions{
		Term: "estonia",
	})

	if err != nil {
		t.Errorf("Could not get currencies: %v", err)
	}

	if result.Success != true {
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

	if diff := deep.Equal(expectedCurrency, result.Data[0]); diff != nil {
		t.Error(diff)
	}
}
