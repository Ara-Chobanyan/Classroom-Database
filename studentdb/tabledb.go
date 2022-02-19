package studentdb

import "database/sql"

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

// InsertStudent - Creates a new student into the class_records table
func InsertStudent(db *sql.DB, name string, grade float64) (int, error) {
	statement := `INSERT INTO class_records(name,average) VALUES($1, $2) RETURNING id`
	var id int
	err := db.QueryRow(statement, name, grade).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

// UpdateUser - Update a student grade by there id
func UpdateUser(db *sql.DB, id int, grade float64) error {
	statement := `UPDATE class_records SET average=$2 WHERE id=$1`
	_, err := db.Exec(statement, id, grade)
	return err
}

func DeleteUser(db *sql.DB, id int) error {
	statement := `DELETE FROM class_records WHERE id=$1`
	_, err := db.Exec(statement, id)
	return err
}
