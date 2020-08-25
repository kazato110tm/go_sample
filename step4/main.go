package main

import "fmt"

type Item struct {
	category string
	price int
}

func main(){
	var n int
	fmt.Print("How many items? > ")
	fmt.Scan(&n)

	items := make([]Item, 0, n)
	fmt.Printf("len(items): %d\n", len(items))
	fmt.Printf("cap(items): %d\n", cap(items))

	for i := 0 ; i < cap(items) ; i++{
		items = inputItem(items)
		fmt.Printf("len(items): %d\n", len(items))
	}

	showItems(items)
}

func inputItem(items []Item) []Item {
	var item Item

	fmt.Print("name > ")
	fmt.Scan(&item.category)

	fmt.Print("price > ")
	fmt.Scan(&item.price)

	items = append(items, item)

	return items
}

func showItems(items []Item) {
	fmt.Println("==========")

	for i := 0 ; i < len(items); i++ {
		fmt.Printf("paied %d yen for %s \n", items[i].price, items[i].category)
	}
	fmt.Println("==========")
	fmt.Printf("len(items): %d\n", len(items))
	fmt.Printf("cap(items): %d\n", cap(items))
}
