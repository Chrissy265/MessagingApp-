package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"realtime-chat-go-react/pkg/apiServer"
	"realtime-chat-go-react/pkg/config"

	"github.com/gorilla/mux"
)

func setupRoutes() {
	myRouter := mux.NewRouter().StrictSlash(true)
	apiServer.SetRoutes(myRouter)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	fmt.Println("Chat App starting")
	config.InitializeConfiguration(os.Args[1])

	setupRoutes()
}
