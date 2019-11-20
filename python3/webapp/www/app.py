#!/usr/local/bin/python3
__author__ = 'Hongrui'
import sys
sys.path.append('/usr/local/lib/python3.7/site-packages')

import logging;logging.basicConfig(level=logging.INFO)

import asyncio,os,json,time
from aiohttp import web


from datetime import datetime



def index(request):
    return web.Response(body=b'<h1>Awesome</h1>')

app = web.Application()
app.add_routes([web.get('/', index)])

web.run_app(app)
