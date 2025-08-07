package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var tpl *template.Template
var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", serveWS)
	http.Handle("/res/", http.StripPrefix("/res", http.FileServer(http.Dir("./assets"))))
	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	if err := tpl.ExecuteTemplate(w, "game.gohtml", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type wsMessage struct {
	Action string `json:"action"`
	Amount int    `json:"amount"`
}

type stateMessage struct {
	Type    string `json:"type"`
	Payload any    `json:"payload"`
}

func serveWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	defer conn.Close()

	game := NewGame()
	conn.WriteJSON(stateMessage{Type: "state", Payload: game.State(false, "Hand begins.")})

	for {
		var msg wsMessage
		if err := conn.ReadJSON(&msg); err != nil {
			log.Println("read:", err)
			return
		}
		info := ""
		switch msg.Action {
		case "fold":
			info = game.PlayerFold()
		case "call":
			info = game.PlayerCall()
		case "bet":
			info = game.PlayerBet(msg.Amount)
		default:
			info = "Unknown action"
		}
		if msg.Action != "fold" {
			npc := game.NPCAct()
			info = info + " " + npc
		}
		reveal := game.HandOver()
		if err := conn.WriteJSON(stateMessage{Type: "state", Payload: game.State(reveal, info)}); err != nil {
			log.Println("write:", err)
			return
		}
		if game.HandOver() {
			game = NewGame()
			if err := conn.WriteJSON(stateMessage{Type: "state", Payload: game.State(false, "New hand.")}); err != nil {
				log.Println("write:", err)
				return
			}
		}
	}
}
