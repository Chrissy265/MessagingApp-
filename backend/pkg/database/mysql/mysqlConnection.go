package mysql

import (
	. "database/sql"
	"fmt"
	"realtime-chat-go-react/pkg/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *DB

/*
 * Initialize database connection
 */
func initializeMySQL() {
	fmt.Println("dbInit")
	dBConnection, err := Open("mysql", config.EnvironmentConfiguration.Connection.MySQL)
	if err != nil {
		fmt.Println("Connection Failed!!")
		return
	}
	err = dBConnection.Ping()
	if err != nil {
		fmt.Println("Ping Failed!!")
		return
	}
	db = dBConnection
	dBConnection.SetMaxOpenConns(10)
	dBConnection.SetMaxIdleConns(5)
	dBConnection.SetConnMaxLifetime(time.Second * 10)
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
