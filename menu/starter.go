package menu

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Ara-chobanyan/findaverage-with-db/studentdb"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "name"
	password = "password"
	dbname   = "class_average"
)

// Start - Starts the application
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
		psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
		db, err = sql.Open("postgres", psqlInfo)
		must(err)

		allStudents, err := studentdb.AllStudents(db)
		must(err)
		for _, p := range allStudents {
			fmt.Printf("%+v\n", p)
		}
		db.Close()
		Start()
		// Find student by there name
	case "3":
		psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
		db, err = sql.Open("postgres", psqlInfo)
		must(err)

		a := GiveAString("Enter Student Name", os.Stdin)

		fmt.Println(studentdb.FindStudentByName(db, a))
		db.Close()
		Start()
		//Find student by there ID
	case "4":
		psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
		db, err = sql.Open("postgres", psqlInfo)
		must(err)

		a := GiveAInt("Enter the student Id", os.Stdin)

		fmt.Println(studentdb.FindStudentById(db, a))
		db.Close()
		Start()
		//Add a new student to the class
	case "5":
		psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
		db, err = sql.Open("postgres", psqlInfo)
		must(err)

		a := GiveAString("Enter Student name", os.Stdin)
		b := CreateAStudentGrade(os.Stdin)

		_, err := studentdb.InsertStudent(db, a, b)
		must(err)

		db.Close()
		Start()
		// Remove student from the class (ID required)
	case "6":
		psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
		db, err = sql.Open("postgres", psqlInfo)
		must(err)

		a := GiveAInt("Enter student Id", os.Stdin)

		err = studentdb.DeleteUser(db, a)
		must(err)
		db.Close()
		Start()
		// Update student from the class (ID required)
	case "7":
		psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
		db, err = sql.Open("postgres", psqlInfo)
		must(err)

		a := GiveAInt("Enter Student Id", os.Stdin)
		b := CreateAStudentGrade(os.Stdin)

		err = studentdb.UpdateUser(db, a, b)
		must(err)
		db.Close()
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
