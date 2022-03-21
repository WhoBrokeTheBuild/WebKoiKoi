
import asyncio
from enum import Enum
import json
import logging

from .card import *

logger = logging.getLogger(__package__)

class GameState(str, Enum):
    NONE                = 'None'
    WAITING_FOR_PLAYERS = 'WaitingForPlayers'
    PLAYING             = 'Playing'

class Game:

    all_games = {}
    all_games_lock = asyncio.Lock()

    @staticmethod
    async def new(name):
        g = Game(name)

        async with Game.all_games_lock:
            Game.all_games[name] = g

        return g

    def __init__(self, name):
        self.name = name
        self.state = GameState.NONE
        self.turn = -1
        self.deck = new_deck()
        self.field = []
        self.players = []
        self.lock = asyncio.Lock()

    async def remove(self):
        from .player import PlayerState

        async with Game.all_games_lock:
            del Game.all_games[self.name]

        async with self.lock:
            for p in self.players:
                await p.send_message('Game has ended')
                await p.change_state(PlayerState.IN_LOBBY)
                await p.send_game_list()

    async def change_state(self, state):
        from .player import Player

        async with self.lock:
            old_state = self.state
            self.state = state

            for p in self.players:
                await p.socket.send(json.dumps({
                    'type': 'GameStateChange',
                    'oldState': old_state,
                    'state': self.state,
                }))
    
    async def add_player(self, player):
        from .player import PlayerState

        if self.state == GameState.PLAYING:
            raise Exception('Game is already started')

        async with player.lock:
            player.game = self

        async with self.lock:
            self.players.append(player)

        await player.change_state(PlayerState.IN_GAME)

        if len(self.players) == 2:
            await self.start()
        else:
            await self.change_state(GameState.WAITING_FOR_PLAYERS)

    async def remove_player(self, player):
        async with self.lock:
            self.players.remove(player)

        await self.remove()

    async def start(self):
        from .player import Player

        await self.change_state(GameState.PLAYING)

        async with self.lock:
            self.field = self.deck[-8:]
            self.deck = self.deck[:-8]

        await self.send_field_changed()

        async with self.lock:
            for p in Player.all_players:
                p.hand = self.deck[-8:]
                self.deck = self.deck[:-8]
                await p.send_hand_changed()
    
    async def send_field_changed(self):
        async with self.lock:
            field = []
            for card in self.field:
                field.append(card.toDict())

            for p in self.players:
                await p.socket.send(json.dumps({
                    'type': 'FieldChanged',
                    'field': field,
                }))
