package main

import "fmt"

func main() {
	// map [key datatype]value datatype
	website := map[string]string{

		"Google": "https://google.com",
		"Amazon": "https://aws.com",
	}
	fmt.Println(website)

	//accessing by key 
	fmt.Println(website["Amazon"])

	//setting new value 
	website["linkedin"] = "https://linkedin.com"
	fmt.Println(website)

	//delete a map 
	delete(website,"Google")
	fmt.Print(website)


	for index, value := range website{
		fmt.Println("Index",index)
		fmt.Print("Value",value)
	}
}