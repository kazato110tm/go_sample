package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Item struct {
	category string
	price int
}

func main(){
	file, err := os.Create("accountbook.txt")
	if err != nil {
		log.Fatal(err)
	}

	var n int
	fmt.Print("How many items? > ")
	fmt.Scan(&n)

	for i := 0 ; i < n ; i++{
		if err := inputItem(file); err != nil{
			log.Fatal(err)
		}
	}

	if err := showItems(); err != nil{
		log.Fatal(err)
	}
}

func inputItem(file *os.File) error {
	var item Item

	fmt.Print("name > ")
	fmt.Scan(&item.category)

	fmt.Print("price > ")
	fmt.Scan(&item.price)

	line := fmt.Sprintf("%s %d\n", item.category, item.price)
	if _, err := file.WriteString(line); err != nil {
		return err
	}

	return nil
}

func showItems() error {
	file, err := os.Open("accountbook.txt")

	if err != nil {
		return err
	}

	fmt.Println("==========")

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := scanner.Text()

		splited := strings.Split(line, " ")
		if len(splited) != 2 {
			return errors.New("failed to parse line")
		}

		category := splited[0]
		price, err := strconv.Atoi(splited[1])
		if err != nil{
			return err
		}

		fmt.Printf("%s : %d yen\n", category, price)
	}
	if err := scanner.Err(); err != nil{
		return err
	}

	fmt.Println("==========")
	return nil
}
