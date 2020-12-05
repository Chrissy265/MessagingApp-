package repository

import (
	"database/sql"
	"fmt"
	"realtime-chat-go-react/pkg/database/mysql"
	"realtime-chat-go-react/pkg/model"
)

func GetUserContacts(userId int64) ([]model.User, error) {
	var sqlQuery = "select u.idUser, displayName, google_id " +
		"from mydb.usercontacts uc " +
		"join mydb.user u on u.idUser = uc.IdContact " +
		"join mydb.userpreferences up on up.idUser = uc.IdContact " +
		"where uc.idUser = ?"
	stmt, err := mysql.GetMySQLConnection().Prepare(sqlQuery)
	defer closeStmt(stmt)
	users := []model.User{}

	if err != nil {
		return users, err
	}

	res, err := stmt.Query(userId)
	defer closeRows(res)
	if err != nil {
		return users, err
	}
	return getUsers(res, users)
}

func SearchContact(search string, userID int64) ([]model.User, error) {
	var sqlQuery = "select u.idUser, displayName, google_id " +
		"from mydb.user u " +
		"join mydb.userpreferences up on up.idUser = u.idUser " +
		"where (UPPER(up.displayName) like UPPER(?) || UPPER(u.email) like UPPER(?)) " +
		"AND not u.idUser = ? " +
		"AND (select iduserContacts from usercontacts where idUser = ? and idContact = u.idUser) is null "

	stmt, err := mysql.GetMySQLConnection().Prepare(sqlQuery)
	defer closeStmt(stmt)
	users := []model.User{}

	if err != nil {
		return users, err
	}

	search = "%" + search + "%"
	fmt.Println(search)
	res, err := stmt.Query(search, search, userID, userID)
	defer closeRows(res)
	if err != nil {
		return users, err
	}
	return getUsers(res, users)
}

func getUsers(res *sql.Rows, users []model.User) ([]model.User, error) {
	for res.Next() {
		var userId int64
		var displayName string
		var googleID string

		err := res.Scan(&userId, &displayName, &googleID)
		if err != nil {
			return users, err
		}
		var newUser = model.User{
			UserID:      userId,
			DisplayName: displayName,
			GoogleID:    googleID,
		}
		users = append(users, newUser)

	}
	return users, nil
}

func AddNewContact(userID int64, contactID int64) error {
	var doesContactExist = "select iduserContacts from usercontacts where idUser = ? and idContact = ?"

	stmtQuery, err := mysql.GetMySQLConnection().Prepare(doesContactExist)
	if err != nil {
		return err
	}
	defer closeStmt(stmtQuery)

	res, err := stmtQuery.Query(userID, contactID)
	defer closeRows(res)

	if !res.Next() {
		var sqlQuery = "INSERT into usercontacts (idUser, idContact) VALUES(?,?)"
		tx, err := mysql.GetMySQLConnection().Begin()
		if err != nil {
			return err
		}

		//intert new chat
		stmt, err := tx.Prepare(sqlQuery)
		if err != nil {
			tx.Rollback()
			return err
		}
		defer stmt.Close()
		_, err = stmt.Exec(userID, contactID)
		if err != nil {
			fmt.Println(err)
			tx.Rollback()
			return err
		}

		return tx.Commit()
	}
	return nil
}

func DeleteContact(userID int64, contactID int64) error {
	var sqlQuery = "DELETE from usercontacts where idUser = ? and idContact = ?"

	tx, err := mysql.GetMySQLConnection().Begin()
	if err != nil {
		return err
	}

	//intert new chat
	stmt, err := tx.Prepare(sqlQuery)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(userID, contactID)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
