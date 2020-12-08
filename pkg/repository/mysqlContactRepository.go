package repository

import (
	"database/sql"
	"fmt"
	"realtime-chat-go-react/pkg/database/mysql"
	"realtime-chat-go-react/pkg/model"
)

func GetUserContacts(userId int64) ([]model.Contact, error) {
	var sqlQuery = "select u.idUser, displayName, google_id " +
		"from usercontacts uc " +
		"join user u on u.idUser = uc.IdContact " +
		"join userpreferences up on up.idUser = uc.IdContact " +
		"where uc.idUser = ?"
	stmt, err := mysql.GetMySQLConnection().Prepare(sqlQuery)
	defer closeStmt(stmt)
	users := []model.User{}
	usersContacts := []model.Contact{}
	if err != nil {
		return usersContacts, err
	}

	res, err := stmt.Query(userId)
	defer closeRows(res)
	if err != nil {
		return usersContacts, err
	}
	userarray, _ := getUsers(res, users)
	chats, _ := ReturnAllUserChats(userId)

	for _, u := range userarray {
		contact := model.Contact{
			ChatId:  -1,
			Contact: u,
		}
		usersContacts = append(usersContacts, contact)
	}

	for _, chat := range chats {
		for _, chatUsers := range chat.Users {
			for i, user := range usersContacts {
				if chatUsers.UserID == user.Contact.UserID {
					usersContacts[i].ChatId = chat.ID
					break
				}
			}
		}
	}

	return usersContacts, err
}

func SearchContact(search string, userID int64) ([]model.User, error) {
	var sqlQuery = "select u.idUser, displayName, google_id " +
		"from user u " +
		"join userpreferences up on up.idUser = u.idUser " +
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
