
import copy
from enum import Enum
import logging
import random

class CardSuit(str, Enum):
    PINE            = "Pine"
    PLUM_BLOSSOM    = "Plum Blossom"
    CHERRY_BLOSSOM  = "Cherry Blossom"
    WISTERIA        = "Wisteria"
    IRIS            = "Iris"
    PEONY           = "Peony"
    CLOVER          = "Clover"
    PAMPAS          = "Pampas"
    CHRYSANTHEMUM   = "Chrysanthemum"
    MAPLE           = "Maple"
    WILLOW          = "Willow"
    PAULOWNIA       = "Paulownia"

class CardCategory(str, Enum):
    PLAIN   = "Plain"
    RIBBON  = "Ribbon"
    ANIMAL  = "Animal"
    BRIGHT  = "Bright"

class CardSpecial(str, Enum):
    NONE            = ""
    BOAR            = "Boar"
    DEER            = "Deer"
    BUTTERFLY       = "Butterfly"
    SAKE_CUP        = "Sake Cup"
    RAIN_MAIN       = "Rain Man"
    POETRY_RIBBON   = "Poetry Ribbon"
    BLUE_RIBBON     = "Blue Ribbon"
    YOKAI           = "Yokai"

class Card:
    def __init__(self, id, suit, category, special=CardSpecial.NONE):
        self.id = id
        self.suit = suit
        self.category = category
        self.special = special

    def toDict(self):
        return {
            'id': self.id,
            'suit': self.suit,
            'category': self.category,
            'special': self.special,
        }

all_cards = [
    Card(0,  CardSuit.PINE, CardCategory.PLAIN),
    Card(1,  CardSuit.PINE, CardCategory.PLAIN),
    Card(2,  CardSuit.PINE, CardCategory.RIBBON, CardSpecial.POETRY_RIBBON),
    Card(3,  CardSuit.PINE, CardCategory.BRIGHT),

    Card(4,  CardSuit.PLUM_BLOSSOM, CardCategory.PLAIN),
    Card(5,  CardSuit.PLUM_BLOSSOM, CardCategory.PLAIN),
    Card(6,  CardSuit.PLUM_BLOSSOM, CardCategory.RIBBON, CardSpecial.POETRY_RIBBON),
    Card(7,  CardSuit.PLUM_BLOSSOM, CardCategory.ANIMAL),

    Card(8,  CardSuit.CHERRY_BLOSSOM, CardCategory.PLAIN),
    Card(9,  CardSuit.CHERRY_BLOSSOM, CardCategory.PLAIN),
    Card(10, CardSuit.CHERRY_BLOSSOM, CardCategory.RIBBON, CardSpecial.POETRY_RIBBON),
    Card(11, CardSuit.CHERRY_BLOSSOM, CardCategory.BRIGHT),

    Card(12, CardSuit.WISTERIA, CardCategory.PLAIN),
    Card(13, CardSuit.WISTERIA, CardCategory.PLAIN),
    Card(14, CardSuit.WISTERIA, CardCategory.RIBBON),
    Card(15, CardSuit.WISTERIA, CardCategory.ANIMAL),
    
    Card(16, CardSuit.IRIS, CardCategory.PLAIN),
    Card(17, CardSuit.IRIS, CardCategory.PLAIN),
    Card(18, CardSuit.IRIS, CardCategory.RIBBON),
    Card(19, CardSuit.IRIS, CardCategory.ANIMAL, CardSpecial.YOKAI),
    
    Card(20, CardSuit.PEONY, CardCategory.PLAIN),
    Card(21, CardSuit.PEONY, CardCategory.PLAIN),
    Card(22, CardSuit.PEONY, CardCategory.RIBBON, CardSpecial.BLUE_RIBBON),
    Card(23, CardSuit.PEONY, CardCategory.ANIMAL, CardSpecial.BUTTERFLY),
    
    Card(24, CardSuit.CLOVER, CardCategory.PLAIN),
    Card(25, CardSuit.CLOVER, CardCategory.PLAIN),
    Card(26, CardSuit.CLOVER, CardCategory.RIBBON),
    Card(27, CardSuit.CLOVER, CardCategory.ANIMAL, CardSpecial.BOAR),
    
    Card(28, CardSuit.PAMPAS, CardCategory.PLAIN),
    Card(29, CardSuit.PAMPAS, CardCategory.PLAIN),
    Card(30, CardSuit.PAMPAS, CardCategory.ANIMAL),
    Card(31, CardSuit.PAMPAS, CardCategory.BRIGHT),
    
    Card(32, CardSuit.CHRYSANTHEMUM, CardCategory.PLAIN),
    Card(33, CardSuit.CHRYSANTHEMUM, CardCategory.PLAIN),
    Card(34, CardSuit.CHRYSANTHEMUM, CardCategory.RIBBON, CardSpecial.BLUE_RIBBON),
    Card(35, CardSuit.CHRYSANTHEMUM, CardCategory.ANIMAL, CardSpecial.SAKE_CUP),
    
    Card(36, CardSuit.MAPLE, CardCategory.PLAIN),
    Card(37, CardSuit.MAPLE, CardCategory.PLAIN),
    Card(38, CardSuit.MAPLE, CardCategory.RIBBON, CardSpecial.BLUE_RIBBON),
    Card(39, CardSuit.MAPLE, CardCategory.ANIMAL, CardSpecial.DEER),
    
    Card(40, CardSuit.WILLOW, CardCategory.PLAIN, CardSpecial.YOKAI),
    Card(41, CardSuit.WILLOW, CardCategory.RIBBON),
    Card(42, CardSuit.WILLOW, CardCategory.ANIMAL),
    Card(43, CardSuit.WILLOW, CardCategory.BRIGHT, CardSpecial.RAIN_MAIN),
    
    Card(44, CardSuit.PAULOWNIA, CardCategory.PLAIN),
    Card(45, CardSuit.PAULOWNIA, CardCategory.PLAIN),
    Card(46, CardSuit.PAULOWNIA, CardCategory.PLAIN, CardSpecial.YOKAI),
    Card(47, CardSuit.PAULOWNIA, CardCategory.BRIGHT),
]

def new_deck():
    deck = copy.deepcopy(all_cards)
    random.shuffle(deck)
    return deck