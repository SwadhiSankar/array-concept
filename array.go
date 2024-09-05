package main

import (
	"fmt"
	
)



func main() {
	var productName [4] string
	productName[1] = "FirstString" //assign based on index 
	prices := []float64{10.2, 10.3, 10.4, 10.5}
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

	//append MEthod 
	prices = append(prices,3.4)
	fmt.Print(prices)

	//... operator to append large array 
	testArray := [] int {1,2,3,4,5,6}
	fmt.Print(testArray)

	testArrayTwo := [] int {7,8,9,10}
	fmt.Print(testArrayTwo)

	testArray = append(testArray, testArrayTwo...)
	fmt.Print(testArray)
}