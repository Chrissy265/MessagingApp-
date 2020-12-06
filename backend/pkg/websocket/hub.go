package websocket

import (
	"log"
	redisConnection "realtime-chat-go-react/pkg/database/redis"
	"strconv"
	"sync"

	"github.com/garyburd/redigo/redis"

	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

type Hub struct {
	Users map[int64][]User
	sync.RWMutex
}

var UserHub *Hub

func init() {
	UserHub = &Hub{
		Users: make(map[int64][]User),
	}
}

func (h *Hub) NewUser(conn *websocket.Conn, userID int64) *User {
	u := &User{
		ID:           userID,
		conn:         conn,
		ConnectionId: uuid.NewV4().String(),
	}
	if err := redisConnection.PubSubConnection.Subscribe(strconv.FormatInt(u.ID, 10)); err != nil {
		panic(err)
	}
	h.Lock()
	defer h.Unlock()
	h.Users[userID] = append(h.Users[userID])
	return u
}

func (h *Hub) DeleteUser(user User) {
	newUserConnectionIds := []User{}
	connectionIds := h.Users[user.ID]
	for _, u := range connectionIds {
		if u.ConnectionId != user.ConnectionId {
			newUserConnectionIds = append(newUserConnectionIds, u)
		}
	}
	if len(newUserConnectionIds) == 0 {
		if err := redisConnection.PubSubConnection.Unsubscribe(strconv.FormatInt(user.ID, 10)); err != nil {
			panic(err)
		}
		if _, ok := h.Users[user.ID]; ok {
			h.Lock()
			defer h.Unlock()
			delete(h.Users, user.ID)
		}
	} else {
		h.Lock()
		defer h.Unlock()
		h.Users[user.ID] = newUserConnectionIds

	}
}

func DeliverMessages() {
	for {
		switch v := redisConnection.PubSubConnection.Receive().(type) {
		case redis.Message:
			userID, _ := strconv.ParseInt(v.Channel, 10, 64)
			UserHub.findAndDeliver(userID, string(v.Data))

		case redis.Subscription:
			log.Printf("subscription message: %s: %s %d\n", v.Channel, v.Kind, v.Count)

		case error:
			log.Println("error pub/sub, delivery has stopped")
			return
		}
	}
}

func (h *Hub) findAndDeliver(userID int64, content string) {
	m := Message{
		Content: content,
	}

	for _, u := range h.Users[userID] {
		if err := u.conn.WriteJSON(m); err != nil {
			log.Printf("error on message delivery e: %s\n", err)
		} else {
			log.Printf("user %s found, message sent\n", userID)
		}
	}
	log.Printf("user %s not found at our hub\n", userID)
}
