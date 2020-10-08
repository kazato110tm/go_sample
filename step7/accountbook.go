package main

import "database/sql"

type Item struct {
	ID       int
	Category string
	Price    int
}

type AccountBook struct {
	db *sql.DB
}

func NewAccountBook(db *sql.DB) *AccountBook {
	return &AccountBook{db: db}
}

func (ab *AccountBook) CreateTable() error {
	const sqlStr = `CREATE TABLE IF NOT EXISTS items(
			id INTEGER PRIMARY KEY,
			category TEXT NOT NULL,
			price INTEGER NOT NULL
		);`

	_, err := ab.db.Exec(sqlStr)
	if err != nil {
		return err
	}

	return nil
}

func (ab *AccountBook) AddItem(item *Item) error {
	const sqlStr = `INSERT INTO items(category, price) VALUES (?,?);`
	_, err := ab.db.Exec(sqlStr, item.Category, item.Price)
	if err != nil {
		return err
	}
	return nil
}

func (ab *AccountBook) GetItems(limit int) ([]*Item, error) {
	const sqlStr = `SELECT * FROM items ORDER BY id DESC LIMIT ?`
	rows, err := ab.db.Query(sqlStr, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*Item
	for rows.Next() {
		var item Item
		err := rows.Scan(&item.ID, &item.Category, &item.Price)
		if err != nil {
			return nil, err
		}
		items = append(items, &item)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}
