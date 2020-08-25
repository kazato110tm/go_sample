package main

import (
	"fmt"
	"os"
)

func main(){
	ab := NewAccountBook("accountbook.txt")

LOOP:
	for {
		var mode int
		fmt.Println("[1]input [2]latest 10 items [3]quit")
		fmt.Print(">")
		fmt.Scan(&mode)

		switch mode{
		case 1:
			var n int
			fmt.Print("How many items? >")
			fmt.Scan(&n)

			for i := 0 ; i < n ; i++{
				if err := ab.AddItem(inputItem()); err != nil {
					fmt.Fprintln(os.Stderr, "error:", err)
					break LOOP
				}
			}
		case 2:
			items, err := ab.GetItems(10)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error:", err)
				break LOOP
			}
			showItems(items)
		case 3:
			fmt.Println("system shutdown")
			return
		}
	}
}

func inputItem() *Item{
	var item Item

	fmt.Print("name >")
	fmt.Scan(&item.category)

	fmt.Print("price >")
	fmt.Scan(&item.price)

	return &item
}

func showItems(items []*Item){
	fmt.Println("==========")
	for _, item := range items {
		fmt.Printf("%s : %d\n", item.category, item.price)
	}
	fmt.Println("==========")
}
