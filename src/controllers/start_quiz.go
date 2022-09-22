package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "src/database/quiz_master.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	row, err := db.Query("SELECT * FROM questions ORDER BY id")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var question string
		var answer string
		row.Scan(&id, &question, &answer)
		fmt.Println(id, ") ", question)
		display(answer)
	}
}

func display(answer string) {

	prompt := promptui.Prompt{
		Label: "Answer",
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	result1 := result != answer
	if result1 {
		fmt.Printf("❌ Wrong answer \n")
		return
	}

	fmt.Printf("✅ Correct!\n")
}
