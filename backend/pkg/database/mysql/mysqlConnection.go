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
	dBConnection, err := Open("mysql", "root:root@(localhost:3306)/mydb")
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
