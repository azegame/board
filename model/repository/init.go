package repository

import (
	"database/sql"
	
	"board/model/dao"
)

var	db *sql.DB

func init() {
	db = dao.GetDB()
}



