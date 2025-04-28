package repository

import (
	"database/sql"

	_ "github.com/jackc/pgx/v4"
)
func Open()(*sql.DB,error){
	connStr := "postgres://user:pass@localhost:5437/inventorydb?sslmode=disable"
	db,err:=sql.Open("pgx",connStr)
	if err!=nil{
		return nil,err
	}
	if err:=db.Ping();err!=nil{
		return nil,err
	}
	return db,nil
}