package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func wsConnect(c *websocket.Conn) {
	defer c.Close()

	log.Println("Player Connecting")

	p := NewPlayer(c)

	err := p.MainLoop()
	if err != nil {
		fmt.Println(err)
	}

	p.Remove()

	log.Println("Player Disconnecting")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println("Could not open file.", err)
	}
	err = template.ExecuteTemplate(w, "index.html", r)
	if err != nil {
		fmt.Println("Could not parse template.", err)
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	wsConnect(ws)
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ws", wsHandler)

	fmt.Println("http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
