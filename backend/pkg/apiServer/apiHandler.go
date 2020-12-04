package apiServer

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"realtime-chat-go-react/pkg/repository"
	"strconv"

	"github.com/gorilla/mux"
)

func returnAllUserChats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	vars := mux.Vars(r)
	i, _ := strconv.ParseInt(vars["id"], 10, 64)
	chats, _ := repository.ReturnAllUserChats(i)
	json.NewEncoder(w).Encode(chats)
}

func createNewUserChat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var userIds []int64
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &userIds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	repository.CreateNewUserChat(userIds)
}

func deleteUserChat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func getRecentMesages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	vars := mux.Vars(r)
	i, _ := strconv.ParseInt(vars["id"], 10, 64)
	messages, _ := repository.GetRecentMesages(i)
	json.NewEncoder(w).Encode(messages)
}

func getRecentMesagesBefore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	vars := mux.Vars(r)
	chatID, _ := strconv.ParseInt(vars["chatid"], 10, 64)
	messageID, _ := strconv.ParseInt(vars["messageid"], 10, 64)
	messages, _ := repository.GetRecentMesagesBefore(chatID, messageID)
	json.NewEncoder(w).Encode(messages)
}

func createNewMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var message string
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	chatID, _ := strconv.ParseInt(vars["chatid"], 10, 64)
	userID, _ := strconv.ParseInt(vars["userid"], 10, 64)

	repository.CreateNewMessage(chatID, userID, message)
	_, err = repository.GetUsersToSendMessageTo(chatID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

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
