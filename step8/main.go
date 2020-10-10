package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/tenntenn/sqlite"
)

func main() {
	db, err := sql.Open(sqlite.DriverName, "accountbook.db")
	if err != nil {
		fmt.Fprintln(os.Stderr, "error : ", err)
		os.Exit(1)
	}

	ab := NewAccountBook(db)

	if err := ab.CreateTable(); err != nil {
		fmt.Fprintln(os.Stderr, "error : ", err)
		os.Exit(1)
	}

LOOP:
	for {
		var mode int
		fmt.Println("[1]input [2]lastest 10 objects [3]totalization [4]quit")
		fmt.Print(">")
		fmt.Scan(&mode)

		switch mode {
		case 1:
			var n int
			fmt.Print("How many inputs? >")
			fmt.Scan(&n)
			for i := 0; i < n; i++ {
				if err := ab.AddItem(inputItem()); err != nil {
					fmt.Fprintln(os.Stderr, "error : ", err)
					break LOOP
				}
			}
		case 2:
			items, err := ab.GetItems(10)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error : ", err)
				break LOOP
			}
			showItems(items)
		case 3:
			summaries, err := ab.GetSummaries()
			if err != nil {
				fmt.Fprintln(os.Stderr, "エラー：", err)
				break LOOP
			}
			showSummary(summaries)
		case 4:
			fmt.Println("shutdown")
			return
		}
	}
}

func inputItem() *Item {
	var item Item
	fmt.Print("Item name > ")
	fmt.Scan(&item.Category)

	fmt.Print("Price > ")
	fmt.Scan(&item.Price)

	return &item
}

func showItems(items []*Item) {
	fmt.Println("=================")
	for _, item := range items {
		fmt.Printf("[%04d] %s:%d yen\n", item.ID, item.Category, item.Price)
	}
	fmt.Println("=================")
}

func showSummary(summaries []*Summary) {
	fmt.Println("=================")
	fmt.Printf("品目\t個数\t合計\t平均\n")
	for _, s := range summaries {
		fmt.Printf("%s\t%d\t%d円\t%.2f円\n", s.Category, s.Count, s.Sum, s.Avg())
	}
	fmt.Println("=================")
}
