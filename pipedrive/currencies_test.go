package pipedrive

import (
	"testing"
	"fmt"
)

func TestCurrenciesService_List(t *testing.T) {
	currencies, _, err := client.Currencies.List()

	if err != nil {
		t.Error("Could not get currencies: %v", err)
	}

	if currencies.Success != true {
		t.Error("Unsuccessful currencies request")
	}

	fmt.Println(currencies.Data[0])
}