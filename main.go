package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {

	fsS := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fsS))

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/room/", RoomHandler)
	http.HandleFunc("/favicon.ico", FaviconHandler)
	http.Handle("/ws", websocket.Handler(WebSocketHandler))

	port := 8080
	fmt.Printf("Listening to port: %d", port)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
