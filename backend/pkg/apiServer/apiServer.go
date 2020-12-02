package apiServer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"realtime-chat-go-react/pkg/database"
	"strconv"

	"github.com/gorilla/mux"
)

func returnAllUserChats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	i, _ := strconv.ParseInt(vars["id"], 10, 64)
	chats, _ := database.ReturnAllUserChats(i)
	json.NewEncoder(w).Encode(chats)
}

func createNewUserChat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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
	database.CreateNewUserChat(userIds)
}

func deleteUserChat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func getRecentMesages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	i, _ := strconv.ParseInt(vars["id"], 10, 64)
	messages, _ := database.GetRecentMesages(i)
	json.NewEncoder(w).Encode(messages)
}

func getRecentMesagesBefore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	chatID, _ := strconv.ParseInt(vars["chatid"], 10, 64)
	messageID, _ := strconv.ParseInt(vars["messageid"], 10, 64)
	messages, _ := database.GetRecentMesagesBefore(chatID, messageID)
	json.NewEncoder(w).Encode(messages)
}

func createNewMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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
	fmt.Println(chatID)
	fmt.Println(userID)
	fmt.Println(body)
	fmt.Println(message)
	database.CreateNewMessage(chatID, userID, message)
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

func SetRoutes(myRouter *mux.Router) {
	myRouter.HandleFunc("/chats/user/{id}", returnAllUserChats).Methods("GET")
	myRouter.HandleFunc("/chat", createNewUserChat).Methods("POST")
	myRouter.HandleFunc("/chat/{id}/user/{id}", deleteUserChat).Methods("DELETE")

	myRouter.HandleFunc("/chat/{id}/messages", getRecentMesages).Methods("GET")
	myRouter.HandleFunc("/chat/{chatid}/messages/{messageid}", getRecentMesagesBefore).Methods("GET")
	myRouter.HandleFunc("/chat/{chatid}/user/{userid}/message", createNewMessage).Methods("POST")

	myRouter.HandleFunc("/chat/{id}/preferences/user/{id}", getChatPreferences).Methods("GET")
	myRouter.HandleFunc("/chat/{id}/preferences/user/{id}", setChatPreferences).Methods("POST")

	myRouter.HandleFunc("/preferences/user/{id}", getUserPreferences).Methods("GET")
	myRouter.HandleFunc("/preferences/user/{id}", setUserPreferences).Methods("POST")
}
