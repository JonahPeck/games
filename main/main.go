package main

import (
	"fmt"
)

func main() {
	var firstCurrency string
	var secondCurrency string
	// var firstAmount float64
	// var convertedAmount float64
	fmt.Println("Currency Converter")
	fmt.Println(`Please select your first and second currency: `)
	fmt.Println("Euros, Dollars, Pounds, or Yen")
	fmt.Scan(&firstCurrency, &secondCurrency)
	fmt.Println(`You are converting `, firstCurrency, ` to `, secondCurrency, `.`)

}

//idea to create a man united soccer app that displays team data and all the things :)
