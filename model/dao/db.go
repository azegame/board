package dao

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)


var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("sqlite3", "./board.db")
	
	if err != nil {
        fmt.Println("データベースを開けませんでした。")
    }
}


func GetDB() *sql.DB {
	return db
}



