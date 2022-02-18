package students

import "database/sql"

type student struct {
	id    int
	name  string
	grade float32
}

// GetStudent - Finds a single student with there name and grade
func GetStudent(db *sql.DB, id int) (*student, error) {
	var s student
	row := db.QueryRow("SELECT * FROM class_records WHERE id=$1", id)
	err := row.Scan(&s.id, &s.name, &s.grade)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func AllStudents(db *sql.DB) ([]student, error) {
	rows, err := db.Query("SELECT * FROM class_records")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret []student
	for rows.Next() {
		var s student

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
