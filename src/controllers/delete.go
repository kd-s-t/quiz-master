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
	fmt.Print("Enter an ID to delete a question: ")
	reader := bufio.NewReader(os.Stdin)
	id, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}
	// remove the delimeter from the string
	id = strings.TrimSuffix(id, "\n")

	delete(id)
}

func delete(id string) {
	db, err := sql.Open("sqlite3", "src/database/quiz_master.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	fmt.Printf("Deleting a question.... \n")

	sts := `DELETE FROM questions WHERE id =` + id + `;`
	_, err = db.Exec(sts)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("success! \n")
}
