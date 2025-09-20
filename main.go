package main

import (
	"fmt"
	"net/http"
	"vinted-sniper/client"
)

func main() {
	httpClient := &http.Client{}
	items, err := client.FetchVintedItems(httpClient)
	if err != nil {
		fmt.Printf("Error fetching items: %v\n", err)
		return
	}

	client.PrintItems(items)

}
