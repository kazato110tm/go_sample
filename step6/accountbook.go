package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Item struct {
	category string
	price int
}

type AccountBook struct {
	fileName string
}

func NewAccountBook(fileName string) *AccountBook{
	return &AccountBook{fileName: fileName}
}

func (ab *AccountBook) AddItem(item *Item) error {
	file, err := os.OpenFile(ab.fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	if _, err := fmt.Fprintln(file, item.category, item.price); err != nil{
		return err
	}

	if err := file.Close(); err != nil{
		return err
	}

	return nil
}

func (ab *AccountBook) GetItems(limit int) ([]*Item, error){
	file, err := os.Open(ab.fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var items []*Item

	for scanner.Scan(){
		var item Item

		if err := ab.parseLine(scanner.Text(), &item); err != nil {
			return nil, err
		}
		items = append(items, &item)
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	if len(items) < limit {
		return items, nil
	}

	return items[len(items) - limit : len(items) : len(items)], nil
}

func (ab *AccountBook) parseLine(line string, item *Item) error{
	splited := strings.Split(line, " ")
	if len(splited) != 2 {
		return errors.New("failed to parse line")
	}

	category := splited[0]

	price, err := strconv.Atoi(splited[1])
	if err != nil {
		return err
	}

	item.category = category
	item.price = price

	return nil
}
