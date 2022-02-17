package students

import "database/sql"

func createDB(db *sql.DB, name string) error {
	_, err := db.Exec("CREATE DATABASE " + name)
	if err != nil {
		return err
	}
	return nil
}

func InsertStudent(db *sql.DB, name string, grade float32) (int, error) {
	statement := `INSERT INTO class_records(name,average) VALUES($1, $2) RETURNING id`
	var id int
	err := db.QueryRow(statement, name, grade).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func ResetDB(db *sql.DB, name string) error {
	_, err := db.Exec("DROP DATABASE IF EXISTS " + name)
	if err != nil {
		return err
	}
	return createDB(db, name)
}

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
