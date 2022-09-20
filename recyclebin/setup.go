package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func createTable(db *sql.DB) {
	q_table := `CREATE TABLE q (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			"question" TEXT,
			"answer" TEXT,`
	query, err := db.Prepare(q_table)
	if err != nil {
		log.Fatal(err)
	}
	query.Exec()
	fmt.Println("Table created successfully!")
}

func createDB() {
	os.Remove("./quiz_master.db")
	file, err := os.Create("quiz_master.db")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}

func main() {
	createDB()
	database, err := sql.Open("sqlite3", "quiz_master.db")
	if err != nil {
		log.Fatal(err)
	}

	createTable(database)
}
