package main

import (
	"fmt"
	"log"
	"net/http"
	"realtime-chat-go-react/pkg/apiServer"
	"realtime-chat-go-react/pkg/websocket"

	"github.com/gorilla/mux"
)

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket Endpoint Hit")
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	myRouter := mux.NewRouter().StrictSlash(true)
	apiServer.SetRoutes(myRouter)

	myRouter.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	fmt.Println("Distributed Chat App v0.01")
	setupRoutes()
}
