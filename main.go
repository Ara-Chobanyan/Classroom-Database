package main

import (
	"database/sql"
	"fmt"

	"github.com/Ara-chobanyan/findaverage-with-db/students"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "ara"
	password = "arayik01"
	dbname   = "class_average"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", host, port, user, password)
	// db, err := sql.Open("postgres", psqlInfo)
	// must(err)
	// err = students.ResetDB(db, dbname)
	// must(err)
	// db.Close()

	psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
  db, err := sql.Open("postgres", psqlInfo)
	must(err)
	defer db.Close()

	must(students.CreateClassTable(db))
	id, err := students.InsertStudent(db, "Jon", 78.88)
	must(err)
	fmt.Println("id=", id)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

/*
We want to have a new database made
every time the program starts again.
- [x] Create a bare bones talbe
*/
