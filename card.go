package main

import (
	"math/rand"
	"time"
)

type CardSuit string
type CardType string
type CardSpecial string

const (
	SuitPine          CardSuit = "Pine"
	SuitPlumBlossom            = "Plum Blossom"
	SuitCherryBlossom          = "Cherry Blossom"
	SuitWisteria               = "Wisteria"
	SuitIris                   = "Iris"
	SuitPeony                  = "Peony"
	SuitClover                 = "Clover"
	SuitPampas                 = "Pampas"
	SuitChrysanthemum          = "Chrysanthemum"
	SuitMaple                  = "Maple"
	SuitWillow                 = "Willow"
	SuitPaulownia              = "Paulownia"

	TypePlain  CardType = "Plain"
	TypeRibbon          = "Ribbon"
	TypeAnimal          = "Animal"
	TypeBright          = "Bright"

	SpecialBoar         CardSpecial = "Boar"
	SpecialDeer                     = "Deer"
	SpecialButterfly                = "Butterfly"
	SpecialSakeCup                  = "Sake Cup"
	SpecialRainMan                  = "Rain Man"
	SpecialPoetryRibbon             = "Poetry Ribbon"
	SpecialBlueRibbon               = "Blue Ribbon"
)

// Card represents a card
type Card struct {
	ID      int         `json:"id"`
	Suit    CardSuit    `json:"suit"`
	Type    CardType    `json:"type"`
	Special CardSpecial `json:"special,omitempty"`
}

// AllCards contains the information of every card
var AllCards = []Card{
	Card{
		ID:   0,
		Suit: SuitPine,
		Type: TypePlain,
	},
	Card{
		ID:   1,
		Suit: SuitPine,
		Type: TypePlain,
	},
	Card{
		ID:      2,
		Suit:    SuitPine,
		Type:    TypeRibbon,
		Special: SpecialPoetryRibbon,
	},
	Card{
		ID:   3,
		Suit: SuitPine,
		Type: TypeBright,
	},
	Card{
		ID:   4,
		Suit: SuitPlumBlossom,
		Type: TypePlain,
	},
	Card{
		ID:   5,
		Suit: SuitPlumBlossom,
		Type: TypePlain,
	},
	Card{
		ID:      6,
		Suit:    SuitPlumBlossom,
		Type:    TypeRibbon,
		Special: SpecialPoetryRibbon,
	},
	Card{
		ID:   7,
		Suit: SuitPlumBlossom,
		Type: TypeAnimal,
	},
	Card{
		ID:   8,
		Suit: SuitCherryBlossom,
		Type: TypePlain,
	},
	Card{
		ID:   9,
		Suit: SuitCherryBlossom,
		Type: TypePlain,
	},
	Card{
		ID:      10,
		Suit:    SuitCherryBlossom,
		Type:    TypeRibbon,
		Special: SpecialPoetryRibbon,
	},
	Card{
		ID:   11,
		Suit: SuitCherryBlossom,
		Type: TypeBright,
	},
	Card{
		ID:   12,
		Suit: SuitWisteria,
		Type: TypePlain,
	},
	Card{
		ID:   13,
		Suit: SuitWisteria,
		Type: TypePlain,
	},
	Card{
		ID:   14,
		Suit: SuitWisteria,
		Type: TypeRibbon,
	},
	Card{
		ID:   15,
		Suit: SuitWisteria,
		Type: TypeAnimal,
	},
	Card{
		ID:   16,
		Suit: SuitIris,
		Type: TypePlain,
	},
	Card{
		ID:   17,
		Suit: SuitIris,
		Type: TypePlain,
	},
	Card{
		ID:   18,
		Suit: SuitIris,
		Type: TypeRibbon,
	},
	Card{
		ID:   19,
		Suit: SuitIris,
		Type: TypeAnimal,
	},
	Card{
		ID:   20,
		Suit: SuitPeony,
		Type: TypePlain,
	},
	Card{
		ID:   21,
		Suit: SuitPeony,
		Type: TypePlain,
	},
	Card{
		ID:      22,
		Suit:    SuitPeony,
		Type:    TypeRibbon,
		Special: SpecialBlueRibbon,
	},
	Card{
		ID:      23,
		Suit:    SuitPeony,
		Type:    TypeAnimal,
		Special: SpecialButterfly,
	},
	Card{
		ID:   24,
		Suit: SuitClover,
		Type: TypePlain,
	},
	Card{
		ID:   25,
		Suit: SuitClover,
		Type: TypePlain,
	},
	Card{
		ID:   26,
		Suit: SuitClover,
		Type: TypeRibbon,
	},
	Card{
		ID:      27,
		Suit:    SuitClover,
		Type:    TypeAnimal,
		Special: SpecialBoar,
	},
	Card{
		ID:   28,
		Suit: SuitPampas,
		Type: TypePlain,
	},
	Card{
		ID:   29,
		Suit: SuitPampas,
		Type: TypePlain,
	},
	Card{
		ID:   30,
		Suit: SuitPampas,
		Type: TypeAnimal,
	},
	Card{
		ID:   31,
		Suit: SuitPampas,
		Type: TypeBright,
	},
	Card{
		ID:   32,
		Suit: SuitChrysanthemum,
		Type: TypePlain,
	},
	Card{
		ID:   33,
		Suit: SuitChrysanthemum,
		Type: TypePlain,
	},
	Card{
		ID:      34,
		Suit:    SuitChrysanthemum,
		Type:    TypeRibbon,
		Special: SpecialBlueRibbon,
	},
	Card{
		ID:      35,
		Suit:    SuitChrysanthemum,
		Type:    TypeAnimal,
		Special: SpecialSakeCup,
	},
	Card{
		ID:   36,
		Suit: SuitMaple,
		Type: TypePlain,
	},
	Card{
		ID:   37,
		Suit: SuitMaple,
		Type: TypePlain,
	},
	Card{
		ID:      38,
		Suit:    SuitMaple,
		Type:    TypeRibbon,
		Special: SpecialBlueRibbon,
	},
	Card{
		ID:      39,
		Suit:    SuitMaple,
		Type:    TypeAnimal,
		Special: SpecialDeer,
	},
	Card{
		ID:   40,
		Suit: SuitWillow,
		Type: TypePlain,
	},
	Card{
		ID:   41,
		Suit: SuitWillow,
		Type: TypeRibbon,
	},
	Card{
		ID:   42,
		Suit: SuitWillow,
		Type: TypeAnimal,
	},
	Card{
		ID:      43,
		Suit:    SuitWillow,
		Type:    TypeBright,
		Special: SpecialRainMan,
	},
	Card{
		ID:   44,
		Suit: SuitPaulownia,
		Type: TypePlain,
	},
	Card{
		ID:   45,
		Suit: SuitPaulownia,
		Type: TypePlain,
	},
	Card{
		ID:   46,
		Suit: SuitPaulownia,
		Type: TypePlain,
	},
	Card{
		ID:   47,
		Suit: SuitPaulownia,
		Type: TypeBright,
	},
}

func NewDeck() []Card {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	deck := make([]Card, len(AllCards))
	perm := r.Perm(len(deck))
	for i, j := range perm {
		deck[i] = AllCards[j]
	}
	return deck
}
