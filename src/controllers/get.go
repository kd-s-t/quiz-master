package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/alexeyco/simpletable"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "src/database/quiz_master.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	fmt.Print("Which question do you wanna retrieve using ID: ")
	reader := bufio.NewReader(os.Stdin)
	id, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}
	// remove the delimeter from the string
	id = strings.TrimSuffix(id, "\n")

	row, err := db.Query("SELECT * FROM questions WHERE id = '" + id + "'")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Question"},
			{Align: simpletable.AlignCenter, Text: "Answer"},
		},
	}

	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var question string
		var answer string
		row.Scan(&id, &question, &answer)
		r := []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", id)},
			{Text: question},
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%s", answer)},
		}

		table.Body.Cells = append(table.Body.Cells, r)
	}

	table.SetStyle(simpletable.StyleCompactLite)
	fmt.Println(table.String())
}
