package students

import (
	"database/sql"
)

type Student struct {
	id    int
	name  string
	grade float32
}

// createDB - Creates a new PostgreSQL database
func createDB(db *sql.DB, name string) error {
	_, err := db.Exec("CREATE DATABASE " + name)
	if err != nil {
		return err
	}
	return nil
}

// GetStudent - Finds a single student with there name and grade
func GetStudent(db *sql.DB, id int) (*Student, error) {
	var s Student
	row := db.QueryRow("SELECT * FROM class_records WHERE id=$1", id)
	err := row.Scan(&s.id, &s.name, &s.grade)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

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

// InsertStudent - Creates a new student into the class_records table
func InsertStudent(db *sql.DB, name string, grade float32) (int, error) {
	statement := `INSERT INTO class_records(name,average) VALUES($1, $2) RETURNING id`
	var id int
	err := db.QueryRow(statement, name, grade).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

// ResetDB - Drops the database and creates a new one if one already exists
func ResetDB(db *sql.DB, name string) error {
	_, err := db.Exec("DROP DATABASE IF EXISTS " + name)
	if err != nil {
		return err
	}
	return createDB(db, name)
}

// CreateClassTable - Creates a class_records tables to store students name and average grade
func CreateClassTable(db *sql.DB) error {
	statement := `
  CREATE TABLE IF NOT EXISTS class_records(
  id SERIAL,
  name VARCHAR(50),
  average float4
  )
  `
	_, err := db.Exec(statement)
	return err
}
