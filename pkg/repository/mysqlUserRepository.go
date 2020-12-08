package repository

import (
	"fmt"
	"realtime-chat-go-react/pkg/database/mysql"
	"realtime-chat-go-react/pkg/model"
)

func GetUser(clientID string) (model.User, error) {
	var sqlQuery = "select displayName, u.idUser " +
		"from user u " +
		"join userpreferences up on up.idUser = u.idUser " +
		"where u.google_id = ?"
	stmt, err := mysql.GetMySQLConnection().Prepare(sqlQuery)
	defer closeStmt(stmt)

	if err != nil {
		return model.User{}, err
	}

	res, err := stmt.Query(clientID)
	defer closeRows(res)
	if err != nil {
		return model.User{}, err
	}
	var displayName string
	var userId int64
	if res.Next() {
		err = res.Scan(&displayName, &userId)
		if err != nil {
			return model.User{}, err
		}
		return model.User{
			UserID:      userId,
			DisplayName: displayName,
			GoogleID:    clientID,
		}, nil
	}
	return model.User{}, nil
}

func AddNewUser(clientID string, displayName string, email string) (int64, error) {
	var doesContactExist = "select google_id from user where google_id = ?"

	stmtQuery, err := mysql.GetMySQLConnection().Prepare(doesContactExist)
	if err != nil {
		return -1, err
	}
	defer closeStmt(stmtQuery)

	res1, err := stmtQuery.Query(clientID)
	defer closeRows(res1)

	if !res1.Next() {
		var sqlQuery = "INSERT into user (google_id, email) VALUES(?,?)"
		tx, err := mysql.GetMySQLConnection().Begin()
		if err != nil {
			return -1, err
		}

		//intert new chat
		stmt, err := tx.Prepare(sqlQuery)
		if err != nil {
			tx.Rollback()
			return -1, err
		}
		defer stmt.Close()
		res, err := stmt.Exec(clientID, email)
		id, _ := res.LastInsertId()

		sqlQuery = "INSERT into userpreferences (iduser, displayName) VALUES(?,?)"

		//intert new chat
		stmt2, err := tx.Prepare(sqlQuery)
		if err != nil {
			tx.Rollback()
			return -1, err
		}
		defer stmt2.Close()
		_, err = stmt2.Exec(id, displayName)

		if err != nil {
			fmt.Println(err)
			tx.Rollback()
			return -1, err
		}

		return id, tx.Commit()
	}
	return -1, nil
}
