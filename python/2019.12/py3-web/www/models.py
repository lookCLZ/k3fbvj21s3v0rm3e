#!/usr/bin/env python3
# -*- coding: utf-8 -*-

'''
Models for user, blog, comment.
'''

__author__ = 'Michael Liao'

import time
import uuid

from orm import Model, StringField, IntegerField, BooleanField, FloatField, TextField, TimeStampField


def next_id():
    return '%015d%s000' % (int(time.time() * 1000), uuid.uuid4().hex)


class User(Model):
    __table__ = 'users'

    id = StringField(primary_key=True, default=next_id, ddl='varchar(50)')
    email = StringField(ddl='varchar(50)')
    passwd = StringField(ddl='varchar(50)')
    admin = BooleanField()
    name = StringField(ddl='varchar(50)')
    image = StringField(ddl='varchar(500)')
    created_at = FloatField(default=time.time)


class Blog(Model):
    __table__ = 'blogs'

    id = StringField(primary_key=True, default=next_id, ddl='varchar(50)')
    user_id = StringField(ddl='varchar(50)')
    user_name = StringField(ddl='varchar(50)')
    user_image = StringField(ddl='varchar(500)')
    name = StringField(ddl='varchar(50)')
    summary = StringField(ddl='varchar(200)')
    content = TextField()
    created_at = FloatField(default=time.time)


class UniquePwd(Model):
    __table__ = "unique_pwds"

    id = StringField(primary_key=True, default=next_id, ddl='varchar(50)')
    wx_order_id = StringField(ddl='varchar(50)')
    code = StringField(ddl='varchar(50)')
    is_used = BooleanField()
    created = StringField(ddl='varchar(50)')


class WxOrder(Model):
    __table__ = 'orders'

    id = IntegerField(primary_key=True)
    wx_user_id = StringField(ddl='varchar(200)')
    wx_user_name = StringField(ddl='varchar(50)')
    wx_user_image = StringField(ddl='varchar(500)')
    wx_addr = StringField(ddl='varchar(50)')
    wx_sex = BooleanField()
    create_at = StringField(ddl='varchar(50)')
    content = TextField()
    old_price = StringField(ddl='varchar(50)')
    sub_amount = StringField(ddl='varchar(50)')
    people_amount = StringField(ddl='varchar(50)')
    store_name = StringField(ddl='varchar(50)')


class WxJoiner(Model):
    __table__ = 'joiner'

    id = IntegerField(primary_key=True)
    order_id = StringField(ddl='varchar(200)')
    user_id = StringField(ddl='varchar(200)')
    wx_user_name = StringField(ddl='varchar(50)')
    wx_user_image = StringField(ddl='varchar(500)')
    wx_addr = StringField(ddl='varchar(50)')
    wx_sex = BooleanField()
    content = TextField()
    help_amount = StringField(ddl='varchar(50)')
    create_at = StringField(ddl='varchar(50)')


class Comment(Model):
    __table__ = 'comments'

    id = StringField(primary_key=True, default=next_id, ddl='varchar(50)')
    blog_id = StringField(ddl='varchar(50)')
    user_id = StringField(ddl='varchar(50)')
    user_name = StringField(ddl='varchar(50)')
    user_image = StringField(ddl='varchar(500)')
    content = TextField()
    created_at = FloatField(default=time.time)
