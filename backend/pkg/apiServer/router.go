package apiServer

import (
	"net/http"
	"realtime-chat-go-react/pkg/userAuthentication"

	"github.com/gorilla/mux"
)

func SetRoutes(myRouter *mux.Router) {

	myRouter.HandleFunc("/user/{clientid}", getUser).Methods("Get")

	myRouter.HandleFunc("/chats/user/{id}", returnAllUserChats).Methods("GET")
	myRouter.HandleFunc("/chat", createNewUserChat).Methods("POST")
	myRouter.HandleFunc("/chat/{id}/user/{id}", deleteUserChat).Methods("DELETE")

	myRouter.HandleFunc("/chat/{id}/messages", getRecentMesages).Methods("GET")
	myRouter.HandleFunc("/chat/{chatid}/messages/{messageid}", getRecentMesagesBefore).Methods("GET")
	//myRouter.HandleFunc("/chat/{chatid}/user/{userid}/message", createNewMessage).Methods("POST")

	myRouter.HandleFunc("/user/{id}/contacts", getUserContacts).Methods("GET")
	myRouter.HandleFunc("/contacts/search/user/{id}", searchContact).Methods("GET")
	myRouter.HandleFunc("/user/{userid}/contact/{contactid}", deleteContact).Methods("DELETE")
	myRouter.HandleFunc("/user/{userid}/contact/{contactid}", addNewContact).Methods("POST")
	myRouter.HandleFunc("/chat/{chatid}/user/{userid}/message", createNewMessage).Methods("POST")

	myRouter.HandleFunc("/user", addNewUser).Methods("POST")

	myRouter.HandleFunc("/auth/google/login", userAuthentication.OauthGoogleLogin).Methods("GET")
	myRouter.HandleFunc("/auth/google/callback", userAuthentication.OauthGoogleCallback).Methods("GET")
	myRouter.Handle("/", http.FileServer(http.Dir("./login.html")))
}
