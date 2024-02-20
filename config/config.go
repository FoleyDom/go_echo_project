package config

import (
	"log"

	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	// github.com/mattn/go-sqlite3
	db, err := gorm.Open(sqlite.Open("./db/todo.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	sqlStmt := `
	create table if not exists todo (id string not null primary key, text text, checked bool);
	`
	_ = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return nil
	}

	// create table if not exists
	return db
}
