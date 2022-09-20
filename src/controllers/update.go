package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Print("Enter an ID to update a question and answer: ")
	reader := bufio.NewReader(os.Stdin)
	id, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}
	// remove the delimeter from the string
	id = strings.TrimSuffix(id, "\n")
	get(id)

	fmt.Print("Enter an updated question: ")
	question, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}
	// remove the delimeter from the string
	question = strings.TrimSuffix(question, "\n")

	fmt.Print("update the answer: ")
	ans, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}
	// remove the delimeter from the string
	ans = strings.TrimSuffix(ans, "\n")

	update(id, question, ans)
}

func update(id string, new_que string, new_ans string) {
	db, err := sql.Open("sqlite3", "src/database/quiz_master.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	fmt.Printf("Updating a question.... \n")

	sts := `UPDATE questions
	SET question = '` + new_que + `', answer= '` + new_ans + `'
	WHERE id = '` + id + `';
	`
	_, err = db.Exec(sts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("success! \n")
}

func get(id string) {
	db, err := sql.Open("sqlite3", "src/database/quiz_master.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	row, err := db.Query(`SELECT * FROM questions WHERE id = '` + id + `';`)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var question string
		var answer string
		row.Scan(&id, &question, &answer)
		fmt.Println(id, ") ", question, " answer is:", answer)
	}
}
