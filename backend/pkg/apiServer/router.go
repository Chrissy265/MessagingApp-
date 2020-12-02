package apiServer

import "github.com/gorilla/mux"

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
