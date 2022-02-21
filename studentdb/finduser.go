package studentdb

import (
	"database/sql"
	"fmt"
)

type Student struct {
	id    int
	name  string
	grade float32
}

// FindStudent - Finds a single Student by id that returns the students name and grade
func FindStudentById(db *sql.DB, id int) (*Student, error) {
	var s Student
	row := db.QueryRow("SELECT * FROM class_records WHERE id=$1", id)
	err := row.Scan(&s.id, &s.name, &s.grade)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

// FindStudentByName - Finds the Student via there name
func FindStudentByName(db *sql.DB, name string) *Student {
	var s Student
	row := db.QueryRow("SELECT * FROM class_records WHERE name=$1", name)
	err := row.Scan(&s.id, &s.name, &s.grade)
	if err != nil {
		fmt.Println("Could not find Student")
	}

	return &s
}

// AllStudents - Querys all Students in the class table
func AllStudents(db *sql.DB) ([]Student, error) {
	rows, err := db.Query("SELECT * FROM class_records")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret []Student
	for rows.Next() {
		var s Student

		if err := rows.Scan(&s.id, &s.name, &s.grade); err != nil {
			return nil, err
		}
		ret = append(ret, s)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return ret, nil
}
