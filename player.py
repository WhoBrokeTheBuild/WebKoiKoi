
import asyncio
from enum import Enum
import json
import logging
import traceback

from .game import *

logger = logging.getLogger(__package__)

class PlayerState(str, Enum):
    NONE     = 'None'
    IN_LOBBY = 'InLobby'
    IN_GAME  = 'InGame'

class Player:

    all_players = []
    all_players_lock = asyncio.Lock()

    @staticmethod
    async def new(socket):
        p = Player(socket)

        async with Player.all_players_lock:
            Player.all_players.append(p)

        return p

    def __init__(self, socket):
        self.state = PlayerState.NONE
        self.hand = []
        self.game = None
        self.socket = socket
        self.lock = asyncio.Lock()

    async def remove(self):
        async with Player.all_players_lock:
            Player.all_players.remove(self)

        if self.game:
            self.remove_player(self)

    async def main_loop(self):
        await self.change_state(PlayerState.IN_LOBBY)
        await self.send_game_list()

        while True:
            data = await self.socket.recv()
            data = json.loads(data)

            try:
                if self.state == PlayerState.IN_LOBBY:
                    await self.handle_packet_lobby(data)
                elif self.state == PlayerState.IN_GAME:
                    await self.handle_packet_game(data)
            except Exception as e:
                traceback.print_exc() 
                await self.send_error(str(type(e)) + ' ' + str(e))

    async def handle_packet_lobby(self, data):
        if 'type' in data:
            if data['type'] == 'NewGame':
                await self.create_game(data)
            elif data['type'] == 'JoinGame':
                await self.join_game(data)

    async def handle_packet_game(self, data):
        pass

    async def create_game(self, data):
        if 'name' not in data:
            raise Exception("'name' is required to create a new game")

        name = data['name']

        nameTaken = False
        async with Game.all_games_lock:
            if name in Game.all_games.keys():
                nameTaken = True

        if nameTaken:
            raise Exception("Game already exists")

        game = await Game.new(name)
        await game.add_player(self)
    
    async def join_game(self, data):
        if 'name' not in data:
            raise Exception("'name' is required to join a game")

        name = data['name']

        game = None
        async with Game.all_games_lock:
            if data['name'] in Game.all_games.keys():
                game = Game.all_games[name]

        if game:
            await game.add_player(self)
        else:
            raise Exception(f'Game "{name}" does not exist')

    async def change_state(self, state):
        async with self.lock:
            old_state = self.state
            self.state = state
            await self.socket.send(json.dumps({
                'type': 'PlayerStateChange',
                'oldState': old_state,
                'state': self.state,
            }))
    
    async def send_game_list(self):
        game_list = []
        async with Game.all_games_lock:
            for name, g in Game.all_games.items():
                game_list.append({
                    'name': name,
                    'players': len(g.players),
                    'capacity': 2,
                })

        async with self.lock:
            await self.socket.send(json.dumps({
                'type': 'GameListChanged',
                'games': game_list,
            }))

    async def send_hand_changed(self):
        async with self.lock:
            hand = []
            for card in self.hand:
                hand.append(card.toDict())

            await self.socket.send(json.dumps({
                'type': 'HandChanged',
                'hand': hand,
            }))

    async def send_message(self, message):
        async with self.lock:
            await self.socket.send(json.dumps({
                'type': 'Message',
                'message': str(message),
            }))

    async def send_error(self, error):
        async with self.lock:
            await self.socket.send(json.dumps({
                'type': 'Error',
                'message': str(error),
            }))
