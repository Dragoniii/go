package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type foo struct {
	Message string
	Age     int
	Name    string
	Surname string
}

type Product struct {
    ProductID        int    `json:"productId"`
    Manufacturer     string `json:"manufacturer"`
    Sku              string `json:"sku"`
    Upc              string `json:"upc"`
    PricePerUnit     string `json:"pricePerUnit"`
    QuantityOnHand   int    `json:"quantityOnHand"`
    ProductName      string `json:"productName"`
}

func marshall() {
	data, _ := json.Marshal(&foo{"Hi", 31, "Joel", "Jo"})
	fmt.Println(string(data))
	//
	product := &Product{
		ProductID:      123,
		Manufacturer:   "Big Box Company",
		PricePerUnit:   "12.99",
		Sku:            "4561qHJK",
		Upc:            "414654444566",
		QuantityOnHand: 28,
		ProductName:    "Gizmo",
	}
	productJson, err := json.Marshal(product)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(productJson))
	//
}