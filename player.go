package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type PlayerState string

const (
	PlayerStateNone    PlayerState = "None"
	PlayerStateInLobby             = "InLobby"
	PlayerStateInGame              = "InGame"
)

type Player struct {
	State PlayerState
	Hand  []Card
	Game  *Game
	Conn  *websocket.Conn
	Mutex sync.RWMutex
}

var allPlayers = []*Player{}
var allPlayersMutex sync.RWMutex

func NewPlayer(c *websocket.Conn) *Player {
	p := &Player{
		State: PlayerStateNone,
		Hand:  []Card{},
		Game:  nil,
		Conn:  c,
	}

	allPlayersMutex.Lock()
	allPlayers = append(allPlayers, p)
	allPlayersMutex.Unlock()

	return p
}

func (p *Player) Remove() {
	if p.Game != nil {
		p.Game.RemovePlayer(p)
	}
}

func (p *Player) ChangeState(state PlayerState) error {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	oldState := p.State
	p.State = state

	log.Println("Sending PlayerStateChange", state)
	return p.Conn.WriteJSON(map[string]interface{}{
		"type":     "PlayerStateChange",
		"oldState": oldState,
		"state":    p.State,
	})
}

func (p *Player) SendGameList() error {
	type game struct {
		Name     string `json:"name"`
		Players  int    `json:"players"`
		Capacity int    `json:"capacity"`
	}

	gameList := []game{}
	allGamesMutex.RLock()
	for name, g := range allGames {
		gameList = append(gameList, game{
			Name:     name,
			Players:  len(g.Players),
			Capacity: 2,
		})
	}
	allGamesMutex.RUnlock()

	p.Mutex.RLock()
	defer p.Mutex.RUnlock()

	log.Println("Sending GameListChanged")
	return p.Conn.WriteJSON(map[string]interface{}{
		"type":  "GameListChanged",
		"games": gameList,
	})
}

func (p *Player) SendError(err error) error {
	p.Mutex.RLock()
	defer p.Mutex.RUnlock()

	log.Println("Sending Error", err)
	return p.Conn.WriteJSON(map[string]interface{}{
		"type":    "Error",
		"message": err.Error(),
	})
}

func (p *Player) SendMessage(msg string) error {
	p.Mutex.RLock()
	defer p.Mutex.RUnlock()

	log.Println("Sending Message", msg)
	return p.Conn.WriteJSON(map[string]interface{}{
		"type":    "Message",
		"message": msg,
	})
}

func (p *Player) MainLoop() error {
	p.ChangeState(PlayerStateInLobby)
	p.SendGameList()

	for {
		_, msg, err := p.Conn.ReadMessage()
		if err != nil {
			return err
		}

		packet := map[string]interface{}{}
		err = json.Unmarshal(msg, &packet)
		if err != nil {
			return err
		}

		fmt.Println(packet)

		if t, ok := packet["type"].(string); ok {
			if p.State == "InLobby" {
				switch t {
				case "NewGame":
					if name, ok := packet["name"].(string); ok {
						err = p.CreateNewGame(name)
						if err != nil {
							p.SendError(err)
						}
					}
				case "JoinGame":
					if name, ok := packet["name"].(string); ok {
						err = p.JoinGame(name)
						if err != nil {
							p.SendError(err)
						}
					}
				}
			} else if p.State == "InGame" {
				switch t {
				case "LeaveGame":
					p.Game.RemovePlayer(p)
					p.ChangeState(PlayerStateInLobby)
				}
			}
		}
	}
}

func (p *Player) CreateNewGame(name string) error {
	exists := false
	allGamesMutex.RLock()
	_, exists = allGames[name]
	allGamesMutex.RUnlock()

	if exists {
		return fmt.Errorf("Game already exists")
	}

	p.Game = NewGame(name)
	err := p.Game.AddPlayer(p)
	if err != nil {
		return err
	}

	allPlayersMutex.RLock()
	for _, p := range allPlayers {
		err := p.SendGameList()
		if err != nil {
			return err
		}
	}
	allPlayersMutex.RUnlock()

	return nil
}

func (p *Player) JoinGame(name string) error {
	allGamesMutex.RLock()
	var exists bool
	if p.Game, exists = allGames[name]; exists {
		err := p.Game.AddPlayer(p)
		if err != nil {
			return err
		}
	}
	allGamesMutex.RUnlock()

	return nil
}

func (p *Player) SendHandChanged() error {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	log.Println("Sending HandChanged")
	return p.Conn.WriteJSON(map[string]interface{}{
		"type": "HandChanged",
		"hand": p.Hand,
	})
}
