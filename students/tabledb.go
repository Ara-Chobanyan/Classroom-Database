package students

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
func InsertStudent(db *sql.DB, name string, grade float32) (int, error) {
	statement := `INSERT INTO class_records(name,average) VALUES($1, $2) RETURNING id`
	var id int
	err := db.QueryRow(statement, name, grade).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}
