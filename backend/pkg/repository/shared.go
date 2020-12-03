package repository

import "database/sql"

func closeRows(rows *sql.Rows) {
	if rows != nil {
		rows.Close()
	}
}

func closeStmt(stmt *sql.Stmt) {
	if stmt != nil {
		stmt.Close()
	}
}
