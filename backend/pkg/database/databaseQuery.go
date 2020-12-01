package database

import (
	"database/sql"
	"fmt"
	"net/http"
	"realtime-chat-go-react/pkg/model"
)

func ReturnAllUserChats(userID int64) (map[int64]*model.Chat, error) {
	var sqlQuery = "SELECT chat.idChat, ucp.idUser, up.displayName, chat.createTime, chat.updateTime FROM userchatpreferences ucp join chat chat on chat.idChat = ucp.idChat join userpreferences up on ucp.idUser = up.idUser where chat.idChat in (Select idChat from userchatpreferences where idUser = ?)"
	stmt, err := getMySQLConnection().Prepare(sqlQuery)
	defer closeStmt(stmt)
	var chats map[int64]*model.Chat
	chats = make(map[int64]*model.Chat)

	if err != nil {
		return chats, err
	}
	res, err := stmt.Query(userID)
	defer closeRows(res)
	if err != nil {
		return chats, err
	}
	for res.Next() {
		var idChat int64
		var idUser int64
		var displayName string
		var createTime string
		var updateTime string

		err = res.Scan(&idChat, &idUser, &displayName, &createTime, &updateTime)
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
				UpdateTime:  updateTime,
				Users:       users}
			chats[idChat] = &newChat
		}

	}
	return chats, err
}

func createNewUserChat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func deleteUserChat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func GetRecentMesages(chatID int64) ([]model.Message, error) {
	fmt.Println(chatID)
	var sqlQuery = "SELECT idMessages0, idSentByUser, message, createdTime FROM messages0 where idChat = ? order by createdTime"
	stmt, err := getMySQLConnection().Prepare(sqlQuery)
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
	for res.Next() {
		var idMessages0 int64
		var idSentByUser int64
		var message string
		var createdTime string

		err = res.Scan(&idMessages0, &idSentByUser, &message, &createdTime)
		if err != nil {
			return messages, err
		}s
		var newMessage = model.Message{
			MessageId:   idMessages0,
			Message:     message,
			UserId:      idSentByUser,
			CreatedTime: createdTime,
		}
		messages = append(messages, newMessage)

	}
	return messages, err
}

func getRecentMesagesBefore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func createNewMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func getChatPreferences(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func setChatPreferences(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func getUserPreferences(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func setUserPreferences(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

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
