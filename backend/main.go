package main

import (
	"fmt"
	"log"
	"net/http"
	"realtime-chat-go-react/pkg/apiServer"
	redisConnection "realtime-chat-go-react/pkg/database/redis"
	"realtime-chat-go-react/pkg/websocket"

	redis "github.com/garyburd/redigo/redis"

	"github.com/gorilla/mux"
)

func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}
	websocket.UserHub.NewUser(ws, 78)
	for {
		if _, err = redisConnection.RedisConn(); err != nil {
			log.Printf("error on redis conn. %s\n", err)
		}
	}
}

func setupRoutes() {
	myRouter := mux.NewRouter().StrictSlash(true)
	apiServer.SetRoutes(myRouter)

	myRouter.HandleFunc("/ws", serveWs)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	fmt.Println("Distributed Chat App v0.01")
	redisConn, err := redisConnection.RedisConn()
	if err != nil {
		panic(err)
	}
	defer redisConn.Close()

	redisConnection.PubSubConnection = &redis.PubSubConn{Conn: redisConn}
	defer redisConnection.PubSubConnection.Close()

	go websocket.DeliverMessages()
	setupRoutes()
}
