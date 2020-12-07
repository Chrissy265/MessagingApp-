package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"realtime-chat-go-react/pkg/apiServer"
	"realtime-chat-go-react/pkg/config"
	redisConnection "realtime-chat-go-react/pkg/database/redis"
	"realtime-chat-go-react/pkg/repository"
	"realtime-chat-go-react/pkg/websocket"
	"strconv"
	"time"

	redis "github.com/garyburd/redigo/redis"
	socketio "github.com/googollee/go-socket.io"
	"github.com/gorilla/mux"
)

var server *socketio.Server

func createNewMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var message Messagething
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
	chatID, err := strconv.ParseInt(vars["chatid"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userID, err := strconv.ParseInt(vars["userid"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, createdat, err := repository.CreateNewMessage(chatID, userID, message.Message)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	users, err := repository.GetUsersToSendMessageTo(chatID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	websocketMessage := websocket.Message{
		UserID:    strconv.FormatInt(userID, 10),
		ChatID:    strconv.FormatInt(chatID, 10),
		CreatedAt: createdat,
		Content:   message.Message,
	}
	bytes, err := json.Marshal(websocketMessage)
	messageString := string(bytes)

	//rediscon, _ := redisConnection.RedisConn()
	for _, user := range users {
		fmt.Println(user)
		fmt.Println(messageString)
		server.BroadcastToRoom(string(user), "chat message", messageString)
	}
	newMessage := Messagething2{
		MessageId: id,
		CreatedAt: createdat,
	}
	json.NewEncoder(w).Encode(newMessage)

}

func setupRoutes() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/chat/{chatid}/user/{userid}/message", createNewMessage).Methods("POST")
	apiServer.SetRoutes(myRouter)
	server, _ = socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		url := s.URL()
		url.Query().Get("userid")
		s.Join(url.Query().Get("userid"))
		return nil
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	go server.Serve()
	defer server.Close()

	myRouter.Handle("/socket.io/", server)

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	fmt.Println("Chat App starting")
	fmt.Println(os.Args[1])
	config.InitializeConfiguration(os.Args[1])

	fmt.Println("Attempting redis connection")
	var redisConn redis.Conn
	var err error
	for i := 0; i < 10; i++ {
		redisConn, err = redisConnection.RedisConn()
		if err == nil {
			break
		} else {
			time.Sleep(2 * time.Second)
		}
		fmt.Println("Attempting redis connection")
	}

	if err == nil {
		fmt.Println("redis connected")
		defer redisConn.Close()

		redisConnection.PubSubConnection = &redis.PubSubConn{Conn: redisConn}
		defer redisConnection.PubSubConnection.Close()

		go websocket.DeliverMessages()
	} else {
		fmt.Println("Could not start redis. Please make sure to run docker-compose up to start redis in docker container. Websocket functionality will not work")
		fmt.Println(err)
	}

	setupRoutes()
}

type Messagething struct {
	Message string
}

type Messagething2 struct {
	MessageId int64
	CreatedAt string
}
