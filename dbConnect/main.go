package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("")

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/crud_15_11")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	fmt.Println("Connected to db")

	_id := 1
	username := "Ramesh"
	password := "Venom"
	email := "ramesh@google.com"
	fave_supe := "Green Goblin"

	insert, err := db.Query("INSERT into fans VALUES(?,?,?,?,?)", _id, username, password, email, fave_supe)
	if err != nil {
		fmt.Println("Trouble while inserting.")
		panic(err.Error())
	}
	insert.Close()

	res, err := db.Query("SELECT id, username from fans")
	if err != nil {
		fmt.Println("Error while querying:")
		panic(err.Error())
	}

	for res.Next() {
		var row Tag

		err = res.Scan(&row.id, &row.name)
		if err != nil {
			fmt.Println("Trouble reading the result variable.")
			panic(err.Error())
		}
		log.Printf("%d\t%s", row.id, row.name)
	}

	fmt.Println(res)
}

type Tag struct {
	id   int    `json:"_id"`
	name string `json:"name"`
}
