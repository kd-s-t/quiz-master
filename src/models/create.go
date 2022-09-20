package create_model

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

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
