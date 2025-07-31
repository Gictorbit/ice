package mysql

import "database/sql"

type TodoMysql struct {
	db *sql.DB
}

func NewTodoMySQL(db *sql.DB) *TodoMysql {
	return &TodoMysql{db: db}
}
