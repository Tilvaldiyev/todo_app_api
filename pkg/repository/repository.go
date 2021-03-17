package repository

import "database/sql"

type Authorization interface {

}

type TodoList interface {

}

type TodoItem interface {

}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{}
}

// TODO: stmt, err := db.Prepare("INSERT table SET unique_id=? ON DUPLICATE KEY UPDATE id=LAST_INSERT_ID(id)")
//		 res, err := stmt.Exec(unique_id)
//		 lid, err := res.LastInsertId()