package repository

import (
	"database/sql"
	"fmt"
	"net/http"
	"realtime-chat-go-react/pkg/database/mysql"
	"realtime-chat-go-react/pkg/model"
)

func ReturnAllUserChats(userID int64) (map[int64]*model.Chat, error) {
	var sqlQuery = "SELECT chat.idChat, ucp.idUser, up.displayName, chat.createTime " +
		"FROM userchatpreferences ucp " +
		"join chat chat on chat.idChat = ucp.idChat " +
		"join userpreferences up on ucp.idUser = up.idUser " +
		"inner join (SELECT idChat, max(idMessages0) lastMessageSentInChat from messages0 group by idChat) n on n.idChat = chat.idChat " +
		"where chat.idChat in (Select idChat from userchatpreferences where idUser = ?) " +
		"order by lastMessageSentInChat desc"
	fmt.Println("asdfa")
	stmt, err := mysql.GetMySQLConnection().Prepare(sqlQuery)
	defer closeStmt(stmt)
	var chats map[int64]*model.Chat
	chats = make(map[int64]*model.Chat)

	if err != nil {
		fmt.Println(err)
		return chats, err
	}
	res, err := stmt.Query(userID)
	defer closeRows(res)
	if err != nil {
		return chats, err
	}
	fmt.Println("asdfa")
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

		if chat, ok := chats[idChat]; ok {
			chat.Users = append(chat.Users, user)
		} else {
			var users = []model.User{user}
			var newChat = model.Chat{
				CreatedTime: createTime,
				Users:       users}
			chats[idChat] = &newChat
		}

	}
	return chats, err
}

func CreateNewUserChat(users []int64) error {
	var sqlQueryCreateChat = "INSERT into chat (messageTable) VALUES(0)"
	var sqlQueryCreatePreferences = "INSERT into userchatpreferences (idChat, idUser) VALUES(?,?)"

	tx, err := mysql.GetMySQLConnection().Begin()
	if err != nil {
		return err
	}

	//intert new chat
	stmt, err := tx.Prepare(sqlQueryCreateChat)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec()
	if err != nil {
		tx.Rollback()
		return err
	}
	//get new chat id
	id, _ := res.LastInsertId()
	fmt.Println(id)

	//insert user chat preferences for all users in chat
	stmt2, err := tx.Prepare(sqlQueryCreatePreferences)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt2.Close()

	for _, user := range users {
		if _, err := stmt2.Exec(id, user); err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

func deleteUserChat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func GetRecentMesages(chatID int64) ([]model.Message, error) {
	var sqlQuery = "SELECT idMessages0, idSentByUser, message, createdTime FROM messages0 where idChat = ? order by createdTime desc, idMessages0 desc limit 50"
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

func CreateNewMessage(chatId int64, userId int64, message string) (int64, error) {
	var sqlQuery = "INSERT into messages0 (idChat, idSentByUser, message) VALUES(?,?,?)"

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
	res, err := stmt.Exec(chatId, userId, message)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return -1, err
	}
	//get new chat id
	id, _ := res.LastInsertId()

	return id, tx.Commit()
}

func GetUsersToSendMessageTo(chatId int64, messageSenderId int64) ([]int64, error) {
	var sqlQuery = "SELECT idUser where idChat = ? and not idUser = ?"
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
	return ids, nil
}
