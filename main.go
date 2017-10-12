package main

import (
	"github.com/genert/api-client/pipedrive"
	"fmt"
)

const apiKey = "bc5b30cb07ac9572597b427c1767ab650eef03ef"

func main() {
	client := pipedrive.New(apiKey)

	record, _, _ := client.Deals.List()

	fmt.Println("Success = ", record.Success)
	fmt.Println("Deals = ", len(record.Data))

	for i := 0; i < len(record.Data); i++ {
		fmt.Println(record.Data[i].Title)
	}
}
