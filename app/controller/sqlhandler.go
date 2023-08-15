package controller

import "database/sql"

type SQLHandler struct {
	Db *sql.DB
}

func NewSqlHandler() *SQLHandler {
	sqlHandler := SQLHandler{}
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/opc-site?parseTime=true")
	if err != nil {
		panic(err)
	}

	sqlHandler.Db = db

	return &sqlHandler
}
