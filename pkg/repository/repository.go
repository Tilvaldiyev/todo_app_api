package repository

import (
	"github.com/jmoiron/sqlx"
	"todo_app_api"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
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

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}

// TODO: stmt, err := db.Prepare("INSERT table SET unique_id=? ON DUPLICATE KEY UPDATE id=LAST_INSERT_ID(id)")
//		 res, err := stmt.Exec(unique_id)
//		 lid, err := res.LastInsertId()