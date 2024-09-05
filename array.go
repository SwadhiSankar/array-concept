package main

import "fmt"



func main() {
	var productName [4] string
	productName[1] = "FirstString" //assign based on index 
	prices := [5]float64{10.2, 10.3, 10.4, 10.5}
	fmt.Print(prices[3])  //to access based on index
	fmt.Print(productName)

	//slicing 
	featuredPrice := prices[1:3]

	//another way 
	// prices[:3], [1:] 
	fmt.Print(featuredPrice)
	highlightedPrices := prices[:3]
	fmt.Print(highlightedPrices)
    fmt.Print("length : ",len(featuredPrice)) //length //2
	fmt.Print("capacity", cap(featuredPrice)) // capacity //4
}