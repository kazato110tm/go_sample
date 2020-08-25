package main

import "fmt"

type Item struct {
	category string
	price int
}

func main(){
	item := inputItem()

	fmt.Println("==========")
	fmt.Printf("paied %d yen for %s \n", item.price, item.category)
	fmt.Println("==========")
}

func inputItem() Item {
	var item Item

	fmt.Print("name > ")
	fmt.Scan(&item.category)

	fmt.Print("price > ")
	fmt.Scan(&item.price)

	return item
}
