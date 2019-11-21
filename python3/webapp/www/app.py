#!/usr/local/bin/python3
__author__ = 'Hongrui'
import sys,logging,asyncio,os,json,time
sys.path.append('/usr/local/lib/python3.7/site-packages')

from aiohttp import web
from datetime import datetime


print(__name__)
def index(request):
    return web.Response(body=b'<h1>awefont</h1>',content_type="text/html")

def init(loop):
    app = web.Application(loop=loop)
    app.router.add_route('GET', '/', index)
    srv = loop.create_server(app.make_handler(), '127.0.0.1', 9000)
    logging.info('服务已启动：http://127.0.0.1:9000')
    return srv


# Return an asyncio event loop.
# When called from a coroutine or a callback 
# (e.g. scheduled with call_soon or similar API), 
# this function will always return the running event loop.
# If there is no running event loop set, the function 
# will return the result of get_event_loop_policy().get_event_loop() call.
logging.basicConfig(level=logging.INFO)
loop = asyncio.get_event_loop()
loop.run_until_complete(init(loop))
loop.run_forever()