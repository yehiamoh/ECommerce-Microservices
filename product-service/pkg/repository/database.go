package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func Open() (*sql.DB, error){
	db,err:= sql.Open("pgx","host=localhost user=user password=pass dbname=productdb port=5435 sslmode=disable")
	if err!=nil{
		return nil,fmt.Errorf("error in connecting To the database : %v",err)
	}
	if err:=db.Ping();err!=nil{
		fmt.Println("Cannot connect to the database",err)
		return nil,err
	}
	return db,nil
}