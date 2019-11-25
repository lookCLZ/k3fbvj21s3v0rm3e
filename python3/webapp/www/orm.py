#!/usr/local/bin/python3
__author__="Hongrui"

import asyncio,logging,aiomysql 

def log(sql):
    logging.info('SQL:%s' % sql)

async def create_pool(loop, **kw):
    logging.info('create database connection pool...')
    global __pool
    __pool = await aiomysql.create_pool(
        host=kw.get('host','localhsot'),
        port=kw.get('port',3306),
        user=kw.get('user',"root"),
        pwd=kw.get('password',"root"),
        db=kw["db"],
        charset=kw.get("charset","utf8"),
        autocommit=kw.get('autocommit', True),
        maxsize=kw.get('maxsize',10),
        minsize=kw.get("minsize",1),
        loop=loop
    )

async def select(sql,args,size=None):
    log(sql,args)
    global __pool 
    async with __pool.get() as conn:
        async with conn.cursor(aiomysql.DictCursor) as cur:
            await cur.execute(sql.replace("?","%s"),args or ())
            if size:
                rs = await cur.fetchmany(size)
            else:
                rs = await cur.fetchall()
        logging.info('rows returned: %s' % len(rs))
        return rs 

async def execute(sql,args,autocommit=True):
    log(sql)
    async with __pool.get as conn:
        if not autocommit:
            await conn.begin()
        try:
                async with conn.cursor(aiomysql.DictCursor) as cur:
                    await cur.execute(sql.replace('?','%s'),args)
                    affected = cur.rowcount
                if not autocommit:
                    await conn.commit() 
        except BaseException as e:
            if not autocommit:
                await conn.rollback()
            raise
        return affected

def create_args_string(num):
    L = []
    for n in range(num):
        L.append('?')
    return ','.join(L)

class Field(object):
    def __init__(self, name, column_type, primary_key, default):
        self.name = name
        self.column_type = column_type
        self.primary_key = primary_key
        self.default = default
    # 如果要让内部属性不被外部访问，可以把属性的名称前加上两个下划线__
    # 以双下划线开头，并且以双下划线结尾的，是特殊变量，特殊变量是可以直接访问的 
    def __str__(self):
        return '<%s, %s:%s>' % (self.__class__.__name__, self.column_type, self.name)

class StringField(Field):
    def __init__(self, name=None, primary_key=False, default=None, ddl='varchar(100)'):
        super().__init__(name, ddl, primary_key, default)

class BooleanField(Field):
    def __init__(self, name=None, default=False):
        super().__init__(name, 'boolean', False, default)

class IntegerField(Field):
    def __init__(self, name=None, primary_key=False, default=0):
        super().__init__(name, 'bigint', primary_key, default)

class FloatField(Field):
    def __init__(self, name=None, primary_key=False, default=0.0):
        super().__init__(name, 'real', primary_key, default)

class TextField(Field):
    def __init__(self, name=None, default=None):
        super().__init__(name, 'text', False, default)