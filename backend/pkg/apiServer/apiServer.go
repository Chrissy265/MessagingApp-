package apiServer

import (
	"encoding/json"
	"fmt"
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
}

func deleteUserChat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func getRecentMesages(w http.ResponseWriter, r *http.Request) {
	fmt.Println("asdfasdf")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	i, _ := strconv.ParseInt(vars["id"], 10, 64)
	messages, _ := database.GetRecentMesages(i)
	json.NewEncoder(w).Encode(messages)
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

func SetRoutes(myRouter *mux.Router) {
	myRouter.HandleFunc("/chats/user/{id}", returnAllUserChats).Methods("GET")
	myRouter.HandleFunc("/chat/user/{id}", createNewUserChat).Methods("POST")
	myRouter.HandleFunc("/chat/{id}/user/{id}", deleteUserChat).Methods("DELETE")

	myRouter.HandleFunc("/chat/{id}/messages", getRecentMesages).Methods("GET")
	myRouter.HandleFunc("/chat/{id}/messages/{id}", getRecentMesagesBefore).Methods("GET")
	myRouter.HandleFunc("/user/{id}/chat/{id}/message", createNewMessage).Methods("POST")

	myRouter.HandleFunc("/chat/{id}/preferences/user/{id}", getChatPreferences).Methods("GET")
	myRouter.HandleFunc("/chat/{id}/preferences/user/{id}", setChatPreferences).Methods("POST")

	myRouter.HandleFunc("/preferences/user/{id}", getUserPreferences).Methods("GET")
	myRouter.HandleFunc("/preferences/user/{id}", setUserPreferences).Methods("POST")
}
