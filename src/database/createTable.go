package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "src/database/quiz_master.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	fmt.Printf("\nCreating table questions.... \n")

	sts := `
		DROP TABLE IF EXISTS questions;
		CREATE TABLE questions(id INTEGER PRIMARY KEY, question TEXT, answer TEXT);
	`
	_, err = db.Exec(sts)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table questions created! \n")
}
