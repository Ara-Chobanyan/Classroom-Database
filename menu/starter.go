package menu

import (
	"database/sql"
	"fmt"

	"github.com/Ara-chobanyan/findaverage-with-db/studentdb"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "ara"
	password = "arayik01"
	dbname   = "class_average"
)

func Start() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", host, port, user, password)
	db, err := sql.Open("postgres", psqlInfo)
	must(err)
	defer db.Close()

	menu, err := StartMenu()
	if err != nil {
		fmt.Println(err)
	}

	switch menu {
	// Create a new class
	case "1":
		err = studentdb.ResetDB(db, dbname)
		must(err)
		db.Close()

		psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
		db, err = sql.Open("postgres", psqlInfo)
		must(err)

		must(studentdb.CreateClassTable(db))
		db.Close()
		Start()
		// Output the entire roster
	case "2":
		defer db.Close()
		psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
		db, err = sql.Open("postgres", psqlInfo)
		must(err)
		allStudents, err := studentdb.AllStudents(db)
		must(err)
		for _, p := range allStudents {
			fmt.Printf("%+v\n", p)
		}
		Start()
		// Find student by there name
	case "3":
		defer db.Close()
		psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
		db, err = sql.Open("postgres", psqlInfo)
		must(err)

		a := GiveAString("Enter Student Name")

		fmt.Println(studentdb.FindStudentByName(db, a))
		Start()
		//Find student by there ID
	case "4":
		defer db.Close()
		psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
		db, err = sql.Open("postgres", psqlInfo)
		must(err)

		a := GiveAInt("Enter the student Id")

		fmt.Println(studentdb.FindStudent(db, a))
		Start()
		//Add a new student to the class
	case "5":
		defer db.Close()
		psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
		db, err = sql.Open("postgres", psqlInfo)
		must(err)

		a := GiveAString("Enter Student name")
		b := CreateAStudentGrade()

		_, err := studentdb.InsertStudent(db, a, b)
		must(err)

		Start()
		// Remove student from the class (ID required)
	case "6":
		defer db.Close()
		psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
		db, err = sql.Open("postgres", psqlInfo)
		must(err)

		a := GiveAInt("Enter student Id")

		err = studentdb.DeleteUser(db, a)
		must(err)
		Start()
		// Update student from the class (ID required)
	case "7":
		defer db.Close()
		psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
		db, err = sql.Open("postgres", psqlInfo)
		must(err)

		a := GiveAInt("Enter Student Id")
		b := CreateAStudentGrade()

		err = studentdb.UpdateUser(db, a, b)
		must(err)
		Start()
		//exit
	case "8":
		db.Close()
		return
	}

}

// must - Small util tool to panic errors
func must(err error) {
	if err != nil {
		panic(err)
	}
}
