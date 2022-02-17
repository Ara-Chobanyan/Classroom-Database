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
	db, err := sql.Open("postgres", psqlInfo)
	must(err)
	err = students.ResetDB(db, dbname)
	must(err)
	db.Close()

	psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	must(err)
	defer db.Close()

	must(students.CreateClassTable(db))

	id, err := students.InsertStudent(db, "Ara", 88.99)
	must(err)
	_, err = students.InsertStudent(db, "Jon", 99.81)
	must(err)
	_, err = students.InsertStudent(db, "Steve", 76.58)
	must(err)

	student, err := students.GetStudent(db, id)
	must(err)
	fmt.Println("Student grade:", student)

	allStudents, err := students.AllStudents(db)
	must(err)
	for _, p := range allStudents {
		fmt.Printf("%+v\n", p)
	}
}

// must - Small util tool to panic errors
func must(err error) {
	if err != nil {
		panic(err)
	}
}
