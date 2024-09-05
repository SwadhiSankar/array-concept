package main

import "fmt"



func main() {
	var productName [4] string
	productName[1] = "FirstString" //assign based on index 
	prices := [5]float64{10.2, 10.3, 10.4, 10.5}
	fmt.Print(prices[3])  //to access based on index
	fmt.Print(productName)
}