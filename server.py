
import asyncio
import functools
from http import HTTPStatus
import logging
import mimetypes
import os
import websockets

from .player import *
from .game import *

logger = logging.getLogger(__package__)

class Server:
    def __init__(self, host='0.0.0.0', port=8080):
        self.host = host
        self.port = port
        self.root = os.path.dirname(os.path.abspath(__file__))

    def run(self):
        start_server = websockets.serve(self.connect, 
            self.host, self.port, 
            process_request=self.process_request)

        logger.info(f'Listening on port {self.host}:{self.port}')

        asyncio.get_event_loop().run_until_complete(start_server)
        asyncio.get_event_loop().run_forever()
    
    async def process_request(server, path, request_headers):
        logger.debug(f'GET {path}')

        if 'Upgrade' in request_headers:
            return

        if path == '/':
            path = '/index.html'
        
        response_headers = [
            ('Server', 'asyncio websocket server'),
            ('Connection', 'close'),
        ]

        full_path = os.path.realpath(os.path.join(server.root, path[1:]))

        if os.path.commonpath((server.root, full_path)) != server.root or \
            not os.path.exists(full_path) or \
            not os.path.isfile(full_path):
            logger.error(f'Could not open file {full_path}')
            return HTTPStatus.NOT_FOUND, [], b'404 File Not Found'
        
        mime_type = mimetypes.guess_type(path)[0]
        response_headers.append(('Content-Type', mime_type))

        body = open(full_path, 'rb').read()
        response_headers.append(('Content-Length', str(len(body))))
        return HTTPStatus.OK, response_headers, body

    async def connect(self, websocket, path):
        logger.debug('Player Connecting')

        p = await Player.new(websocket)
        await p.main_loop()
        await p.remove()

        logger.debug('Player Disconnecting')