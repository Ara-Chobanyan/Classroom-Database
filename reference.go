// package main
//
// import (
// 	"database/sql"
// 	"fmt"
//
// 	"github.com/Ara-chobanyan/findaverage-with-db/studentdb"
// 	_ "github.com/lib/pq"
// )
//
// const (
// 	host     = "localhost"
// 	port     = "5432"
// 	user     = "ara"
// 	password = "arayik01"
// 	dbname   = "class_average"
// )
//
// func main() {
// 	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", host, port, user, password)
// 	db, err := sql.Open("postgres", psqlInfo)
// 	must(err)
// 	err = studentdb.ResetDB(db, dbname)
// 	must(err)
// 	db.Close()
//
// 	psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
// 	db, err = sql.Open("postgres", psqlInfo)
// 	must(err)
// 	defer db.Close()
//
// 	must(studentdb.CreateClassTable(db))
//
// 	id, err := studentdb.InsertStudent(db, "Ara", 88.99)
// 	must(err)
// 	_, err = studentdb.InsertStudent(db, "Jon", 99.81)
// 	must(err)
// 	_, err = studentdb.InsertStudent(db, "Steve", 76.58)
// 	must(err)
//
// 	student, err := studentdb.FindStudent(db, id)
// 	must(err)
// 	fmt.Println("Student grade:", student)
//
// 	allStudents, err := studentdb.AllStudents(db)
// 	must(err)
// 	for _, p := range allStudents {
// 		fmt.Printf("%+v\n", p)
// 	}
// 	fmt.Println(studentdb.FindStudentByName(db, "Ara"))
// 	err = studentdb.UpdateUser(db, 2, 100)
// 	must(err)
// 	fmt.Println(studentdb.FindStudentByName(db, "Jon"))
// 	err = studentdb.DeleteUser(db, 2)
// 	must(err)
// }
//
// // must - Small util tool to panic errors
// func must(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
