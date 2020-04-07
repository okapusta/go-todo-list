package api

import (
	"database/sql"
)

type todo struct {
	ID        int64  `json:"id"`
	Content   string `json:"content"`
	Completed bool   `json:"completed"`
}

func getTodos(db *sql.DB) ([]todo, error) {
	rows, err := db.Query("SELECT id, content, completed FROM todos")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	todos := []todo{}

	for rows.Next() {
		var t todo
		if err := rows.Scan(&t.ID, &t.Content, &t.Completed); err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}

	return todos, nil
}
