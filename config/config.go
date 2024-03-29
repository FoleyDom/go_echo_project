package config

import (
	"log"

	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	// github.com/mattn/go-sqlite3
	db, err := gorm.Open(sqlite.Open("./db/todo.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	sqlStmt := `
	create table if not exists todos (id integer not null primary key, text text, checked bool);
	`
	db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return nil
	}

	return db
}

// ConnectDB ...

// func ConnectDB() *gorm.DB {
// 	db, err := gorm.Open(sqlite.Open("./db/todo.db"), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	return db
// }
