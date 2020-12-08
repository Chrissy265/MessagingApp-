package repository

import (
	"database/sql"
	"fmt"
	"realtime-chat-go-react/pkg/database/mysql"
	"realtime-chat-go-react/pkg/model"
)

func ReturnAllUserChats(userID int64) ([]model.Chat, error) {
	var sqlQuery = "SELECT chat.idChat, ucp.idUser, up.displayName, chat.createTime " +
		"FROM userchatpreferences ucp " +
		"join chat chat on chat.idChat = ucp.idChat " +
		"join userpreferences up on ucp.idUser = up.idUser " +
		"inner join (SELECT idChat, max(idMessages0) lastMessageSentInChat from messages0 group by idChat) n on n.idChat = chat.idChat " +
		"where chat.idChat in (Select idChat from userchatpreferences where idUser = ?) " +
		"order by lastMessageSentInChat desc"
	stmt, err := mysql.GetMySQLConnection().Prepare(sqlQuery)
	defer closeStmt(stmt)
	var chats []model.Chat

	if err != nil {
		fmt.Println(err)
		return chats, err
	}
	res, err := stmt.Query(userID)
	defer closeRows(res)
	if err != nil {
		return chats, err
	}

	//SQL response is sorted by chat, so we can assume all users in a chat are grouped
	var currentChat model.Chat
	first := true
	for res.Next() {
		var idChat int64
		var idUser int64
		var displayName string
		var createTime string

		err = res.Scan(&idChat, &idUser, &displayName, &createTime)
		if err != nil {
			return chats, err
		}
		var user model.User = model.User{UserID: idUser, DisplayName: displayName}

		if !first && idChat == currentChat.ID {
			currentChat.Users = append(currentChat.Users, user)
		} else {
			if !first {
				chats = append(chats, currentChat)
			}
			first = false
			var users = []model.User{user}
			currentChat = model.Chat{
				CreatedTime: createTime,
				Users:       users,
				ID:          idChat,
			}

		}
	}
	if !first {
		chats = append(chats, currentChat)
	}
	return chats, err
}

func CreateNewUserChat(users []int64) (int64, error) {
	var sqlQueryCreateChat = "INSERT into chat (messageTable) VALUES(0)"
	var sqlQueryCreatePreferences = "INSERT into userchatpreferences (idChat, idUser) VALUES(?,?)"

	tx, err := mysql.GetMySQLConnection().Begin()
	if err != nil {
		return -1, err
	}

	//intert new chat
	stmt, err := tx.Prepare(sqlQueryCreateChat)
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	defer stmt.Close()

	res, err := stmt.Exec()
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	//get new chat id
	id, _ := res.LastInsertId()

	//insert user chat preferences for all users in chat
	stmt2, err := tx.Prepare(sqlQueryCreatePreferences)
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	defer stmt2.Close()

	for _, user := range users {
		if _, err := stmt2.Exec(id, user); err != nil {
			tx.Rollback()
			return -1, err
		}
	}

	return id, tx.Commit()
}

func GetRecentMesages(chatID int64) ([]model.Message, error) {
	var sqlQuery = "SELECT idMessages0, idSentByUser, message, createdTime FROM messages0 where idChat = ? order by createdTime, idMessages0 limit 50"
	stmt, err := mysql.GetMySQLConnection().Prepare(sqlQuery)
	defer closeStmt(stmt)
	messages := []model.Message{}

	if err != nil {
		return messages, err
	}
	res, err := stmt.Query(chatID)
	defer closeRows(res)
	if err != nil {
		return messages, err
	}
	return getRecentMesages(res, messages)
}

func GetRecentMesagesBefore(chatID int64, messageId int64) ([]model.Message, error) {
	var sqlQuery = "SELECT idMessages0, idSentByUser, message, createdTime FROM mydb.messages0 where idChat = ? AND createdTime <= (Select createdTime from mydb.messages0 where idMessages0 = ? limit 1) AND idMessages0 < ? order by createdTime desc, idMessages0 desc limit 50"
	stmt, err := mysql.GetMySQLConnection().Prepare(sqlQuery)
	defer closeStmt(stmt)
	messages := []model.Message{}

	if err != nil {
		return messages, err
	}
	res, err := stmt.Query(chatID, messageId, messageId)
	defer closeRows(res)
	if err != nil {
		return messages, err
	}

	return getRecentMesages(res, messages)
}

func getRecentMesages(res *sql.Rows, messages []model.Message) ([]model.Message, error) {
	for res.Next() {
		var idMessages0 int64
		var idSentByUser int64
		var message string
		var createdTime string

		err := res.Scan(&idMessages0, &idSentByUser, &message, &createdTime)
		if err != nil {
			return messages, err
		}
		var newMessage = model.Message{
			MessageId:   idMessages0,
			Message:     message,
			UserId:      idSentByUser,
			CreatedTime: createdTime,
		}
		messages = append(messages, newMessage)

	}
	return messages, nil
}

func CreateNewMessage(chatId int64, userId int64, message string) (int64, string, error) {
	var sqlQuery = "INSERT into messages0 (idChat, idSentByUser, message) VALUES(?,?,?)"
	tx, err := mysql.GetMySQLConnection().Begin()
	if err != nil {
		return -1, "", err
	}

	//intert new chat
	stmt, err := tx.Prepare(sqlQuery)
	if err != nil {
		tx.Rollback()
		return -1, "", err
	}
	defer stmt.Close()
	res, err := stmt.Exec(chatId, userId, message)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return -1, "", err
	}

	err = tx.Commit()
	if err != nil {
		return -1, "", err
	}
	//get new chat id
	id, _ := res.LastInsertId()

	var sqlQuery1 = "SELECT createdTime FROM messages0 where idMessages0 = ?"
	stmt1, err := mysql.GetMySQLConnection().Prepare(sqlQuery1)
	defer closeStmt(stmt1)

	if err != nil {
		return -1, "", err
	}
	res1, err := stmt1.Query(id)
	defer closeRows(res1)
	if err != nil {
		return -1, "", err
	}

	var createdTime string
	if res1.Next() {

		err := res1.Scan(&createdTime)
		fmt.Println(createdTime)
		if err != nil {
			return -1, createdTime, err
		}
	}

	return id, createdTime, err
}

func GetUsersToSendMessageTo(chatId int64, messageSenderId int64) ([]int64, error) {
	var sqlQuery = "SELECT idUser from userchatpreferences where idChat = ? and not idUser = ?"
	stmt, err := mysql.GetMySQLConnection().Prepare(sqlQuery)
	defer closeStmt(stmt)
	ids := []int64{}

	if err != nil {
		return ids, err
	}

	res, err := stmt.Query(chatId, messageSenderId)
	defer closeRows(res)
	if err != nil {
		return ids, err
	}

	for res.Next() {
		var userId int64
		err := res.Scan(&userId)
		if err != nil {
			return ids, err
		}
		ids = append(ids, userId)
	}
	fmt.Println(ids)
	return ids, nil
}
