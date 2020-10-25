#!/usr/bin/env python3

import logging
from .server import Server

if __name__ == "__main__":
    logger = logging.getLogger(__package__)
    logger.setLevel(logging.DEBUG)
    
    stream = logging.StreamHandler()
    stream.setLevel(logging.DEBUG)

    formatter = logging.Formatter("%(levelname)s:%(name)s - %(message)s")
    stream.setFormatter(formatter)

    logger.addHandler(stream)
    
    s = Server()
    s.run()
