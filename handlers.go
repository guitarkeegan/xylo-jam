package main

import (
	"golang.org/x/net/websocket"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles(filepath.Join("templates", "index.html"))

	err = tmpl.Execute(w, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

type Room struct {
	Id   int
	Name string
}

func MakeRoom(name string) *Room {
	return &Room{
		1,
		name,
	}
}

func RoomHandler(w http.ResponseWriter, r *http.Request) {
	slug := strings.TrimPrefix(r.URL.Path, "/room/")

	newRoom := MakeRoom(slug)

	tmpl, err := template.ParseFiles(filepath.Join("templates", "room.html"))

	if err != nil {
		log.Printf("error parsing template: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, newRoom)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func WebSocketHandler(ws *websocket.Conn) {
	defer ws.Close()
	for {
		// Read message from browser
		msg := ""
		err := websocket.Message.Receive(ws, &msg)
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		// Print the message to the console and send it back to the browser
		log.Println("Received:", msg)
		err = websocket.Message.Send(ws, "Received: "+msg)
		if err != nil {
			log.Println("Error sending message:", err)
			break
		}
	}
}

func FaviconHandler(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, filepath.Join("static", "images", "favicon.ico"))

}
