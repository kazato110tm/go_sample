package main

import "fmt"


func main(){
	var category string
	var price int

	fmt.Print("name>")
	fmt.Scan(&category)

	fmt.Print("price>")
	fmt.Scan(&price)

	fmt.Println("========")

	fmt.Printf("%s : %d yen\n", category, price)

	fmt.Println("========")
}
