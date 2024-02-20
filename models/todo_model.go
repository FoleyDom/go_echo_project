package models

import (
	"gorm.io/gorm"

	"github.com/foleydom/go_echo_project/config"
)

type Todo struct {
	// ID      string `gorm:"primaryKey"`
	Text    string
	Checked bool
}

func ConnectDB() *gorm.DB {
	db := config.InitDB()
	return db
}

func CreateTodo(db *gorm.DB, todo *Todo) *Todo {
	db.Create(&todo)
	err := db.Error
	if err != nil {
		panic(err)
	}
	return todo
}

func GetTodos(db *gorm.DB) []Todo {
	var todos []Todo
	db.Find(&todos)
	return todos
}

func GetTodosByChecked(db *gorm.DB, checked bool) []Todo {
	var todos []Todo
	db.Where("checked = ?", checked).Find(&todos)
	return todos
}
