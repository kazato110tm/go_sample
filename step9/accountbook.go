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

func (ab *AccountBook) GetSummaries() ([]*Summary, error) {
	const sqlStr = `
	SELECT
		category,
		COUNT(1) as count,
		SUM(price) as sum
	FROM
		items
	GROUP BY
		category
	`
	rows, err := ab.db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var summaries []*Summary
	for rows.Next() {
		var s Summary
		err := rows.Scan(&s.Category, &s.Count, &s.Sum)
		if err != nil {
			return nil, err
		}
		summaries = append(summaries, &s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return summaries, nil
}

type Summary struct {
	Category string
	Count    int
	Sum      int
}

func (s *Summary) Avg() float64 {
	if s.Count == 0 {
		return 0
	}
	return float64(s.Sum) / float64(s.Count)
}
