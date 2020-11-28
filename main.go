package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TutorialEdge/realtime-chat-go-react/pkg/websocket"
)


//define our WebSocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket Endpoint Hit") 

	conn, err := websocket.Upgrade(w,r) 
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})

	//map our '/ws' endpoint to the 'serveWs' function
	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println(" Distributed Chat App v0.01")
	setupRoutes()
	http.ListenAndServe(":8080", nil)

}
