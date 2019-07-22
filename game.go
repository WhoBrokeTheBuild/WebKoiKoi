package main

import (
	"log"
	"sync"
)

type GameState string

const (
	GameStateNone              GameState = "None"
	GameStateWaitingForPlayers           = "WaitingForPlayers"
	GameStatePlaying                     = "Playing"
)

type Game struct {
	Name    string
	State   GameState
	Turn    int
	Deck    []Card
	Field   []Card
	Players []*Player
	Mutex   sync.RWMutex
}

var allGames = map[string]*Game{}
var allGamesMutex sync.RWMutex

func NewGame(name string) *Game {
	g := &Game{
		Name:    name,
		State:   GameStateNone,
		Turn:    -1,
		Deck:    NewDeck(),
		Field:   []Card{},
		Players: []*Player{},
	}

	allGamesMutex.Lock()
	allGames[name] = g
	allGamesMutex.Unlock()

	return g
}

func (g *Game) Remove() error {
	for _, p := range g.Players {
		p.Game = nil

		p.SendMessage("Game closed")
		p.ChangeState(PlayerStateInLobby)
	}

	allGamesMutex.Lock()
	delete(allGames, g.Name)
	allGamesMutex.Unlock()

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

func (g *Game) AddPlayer(p *Player) error {
	p.Mutex.Lock()
	p.Game = g
	p.Mutex.Unlock()

	g.Mutex.Lock()
	g.Players = append(g.Players, p)
	g.Mutex.Unlock()

	err := p.ChangeState(PlayerStateInGame)
	if err != nil {
		return nil
	}

	err = g.ChangeState(GameStateWaitingForPlayers)
	if err != nil {
		return nil
	}

	if len(g.Players) == 2 {
		err = g.Start()
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *Game) RemovePlayer(p *Player) error {
	g.Mutex.Lock()
	for i := range g.Players {
		if g.Players[i] == p {
			log.Println("Removing player", i)
			g.Players[i] = g.Players[len(g.Players)-1]
			g.Players = g.Players[:len(g.Players)-1]
			break
		}
	}
	g.Mutex.Unlock()

	log.Println("Removing game because player left")
	err := g.Remove()
	if err != nil {
		return err
	}

	return nil
}

func (g *Game) Start() error {
	err := g.ChangeState(GameStatePlaying)
	if err != nil {
		return err
	}

	g.Mutex.Lock()
	g.Field = g.Deck[len(g.Deck)-8:]
	g.Deck = g.Deck[:len(g.Deck)-8]
	g.Mutex.Unlock()
	g.SendFieldChanged()

	g.Mutex.Lock()
	for _, p := range g.Players {
		p.Mutex.Lock()
		p.Hand = g.Deck[len(g.Deck)-8:]
		g.Deck = g.Deck[:len(g.Deck)-8]
		p.Mutex.Unlock()
		p.SendHandChanged()
	}
	g.Mutex.Unlock()

	err = g.ChangeTurn(0)
	if err != nil {
		return err
	}

	return nil
}

func (g *Game) ChangeState(state GameState) error {
	g.Mutex.Lock()
	defer g.Mutex.Unlock()

	oldState := g.State
	g.State = state

	for _, p := range g.Players {
		p.Mutex.RLock()
		log.Println("Sending GameStateChanged")
		err := p.Conn.WriteJSON(map[string]interface{}{
			"type":     "GameStateChange",
			"oldState": oldState,
			"state":    g.State,
		})
		p.Mutex.RUnlock()
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *Game) ChangeTurn(turn int) error {
	g.Mutex.Lock()
	defer g.Mutex.Unlock()

	g.Turn = turn

	for i, p := range g.Players {
		p.Mutex.RLock()
		log.Println("Sending GameStateChanged")
		err := p.Conn.WriteJSON(map[string]interface{}{
			"type":   "TurnChanged",
			"active": (i == turn),
		})
		p.Mutex.RUnlock()
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *Game) SendFieldChanged() error {
	g.Mutex.Lock()
	defer g.Mutex.Unlock()

	for _, p := range g.Players {
		p.Mutex.RLock()
		log.Println("Sending FieldChanged")
		err := p.Conn.WriteJSON(map[string]interface{}{
			"type":  "FieldChanged",
			"field": g.Field,
		})
		p.Mutex.RUnlock()
		if err != nil {
			return err
		}
	}

	return nil
}
