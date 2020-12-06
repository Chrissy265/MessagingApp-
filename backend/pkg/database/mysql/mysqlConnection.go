package mysql

import (
	. "database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *DB

/*
 * Initialize database connection
 */
func initializeMySQL() {
	dBConnection, err := Open("mysql", "admin:Password1!@(database-2.canbz0ws5cxo.us-east-2.rds.amazonaws.com:3306)/db")
	if err != nil {
		fmt.Println("Connection Failed!!")
	}
	err = dBConnection.Ping()
	if err != nil {
		fmt.Println("Ping Failed!!")
	}
	db = dBConnection
	dBConnection.SetMaxOpenConns(10)
	dBConnection.SetMaxIdleConns(5)
	dBConnection.SetConnMaxLifetime(time.Second * 10)
	fmt.Println("dbInit")
}

/*
 * Get database connection
 */
func GetMySQLConnection() *DB {
	if db == nil {
		initializeMySQL()
	}
	return db
}
