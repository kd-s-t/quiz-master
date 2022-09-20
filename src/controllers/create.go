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
	fmt.Print("Enter a question: ")
	reader := bufio.NewReader(os.Stdin)
	question, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}
	// remove the delimeter from the string
	question = strings.TrimSuffix(question, "\n")

	fmt.Print("what is the answer? ")
	ans, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}
	// remove the delimeter from the string
	ans = strings.TrimSuffix(ans, "\n")

	create(question, ans)
}

func create(que string, ans string) {
	db, err := sql.Open("sqlite3", "src/database/quiz_master.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	fmt.Printf("Adding a question.... \n")

	sts := `INSERT INTO questions(question, answer) VALUES('` + que + `?','` + ans + `');`
	_, err = db.Exec(sts)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("success! \n")
}
