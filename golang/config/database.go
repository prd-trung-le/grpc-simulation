package config

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

// ConnectDB to get all needed db connections for application
func ConnectDB() *sqlx.DB {
	DB = getDBConnection()

	return DB
}

func getDBConnection() *sqlx.DB {
	db, err := sqlx.Open("mysql", "root:1345@tcp(127.0.0.1:3306)/grpc_auth")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	//TODO: experiment with correct values
	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(5)

	fmt.Println("Connected to DB")
	return db
}
