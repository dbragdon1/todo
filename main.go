package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Item struct {
	Id          int
	Description string
	Time        string
}

func getItems(db *sql.DB) ([]Item, error) {

	rows, err := db.Query("SELECT rowid, item, time FROM todo;")

	if err != nil {
		panic(err)

	}

	var items []Item

	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.Id, &item.Description, &item.Time); err != nil {
			return items, err
		}

		items = append(items, item)

	}

	if err = rows.Err(); err != nil {
		return items, err
	}

	return items, nil

}

func addItem(db *sql.DB, item string) {

	now := time.Now()
	_, err := db.Exec("INSERT INTO todo (item, time) VALUES (?, ?)", item, now)

	if err != nil {
		panic(err)
	}

}

func printItems(items []Item) {
	for _, item := range items {
		fmt.Println(item.Id, "-", item.Description, "-", item.Time)
	}
}

func popItem(db *sql.DB, itemId int) {

	_, err := db.Exec("DELETE FROM todo WHERE rowid == ?", itemId)

	fmt.Printf("Popping item %d\n", itemId)
	if err != nil {
		panic(err)
	}

}

func main() {

	db, err := sql.Open("sqlite3", "./todo.db")

	db.SetMaxOpenConns(1)

	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS todo (item text, time datetime);")

	if err != nil {
		panic(err)
	}

	args := os.Args[1:]

	if len(args) == 0 {

		items, err := getItems(db)

		if err != nil {
			panic(err)
		}

		printItems(items)

	} else if len(args) == 1 {

		newItem := args[0]

		flag.Parse()

		fmt.Println("Added new todo item: ", newItem)

		addItem(db, newItem)

	} else if len(args) == 2 {
		if args[0] == "pop" {
			popIndex, err := strconv.Atoi(args[1])
			if err != nil {
				panic(err)
			}

			popItem(db, popIndex)
		}
	}

}
