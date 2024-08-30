package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func unmarshall() {
	productJSON := `{
		"productId":456,
		"manufacturer":"Small Box Company",
		"sku":"4hslj90JKL",
		"upc":"424654445566",
		"pricePerUnit":"$9.99",
		"quantityOnHand":18,
		"productName":"Sprocket"
	}`

	product := Product{}
	err := json.Unmarshal([]byte(productJSON), &product)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("json unmarshal product: %s", product.ProductName)
	fmt.Println("")
}