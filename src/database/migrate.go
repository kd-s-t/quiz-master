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

	fmt.Printf("Adding 2 questions.... \n")

	sts := `
		INSERT INTO questions(question, answer) VALUES('1+1?','2');
		INSERT INTO questions(question, answer) VALUES('Mercedes is a?','car');
	`
	_, err = db.Exec(sts)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table questions populated done! \n")
}
