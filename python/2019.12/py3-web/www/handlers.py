#!/usr/bin/env python3
# -*- coding: utf-8 -*-

__author__ = 'Liuhongrui'

' url handlers '

import re
import time
import json
import logging
import hashlib
import base64
import asyncio
import requests

import markdown2

from aiohttp import web

from coroweb import get, post
from apis import Page, APIValueError, APIResourceNotFoundError

from models import User, Comment, Blog, UniquePwd,WxOrder,WxJoiner, next_id
from config import configs

COOKIE_NAME = 'awesession'
_COOKIE_KEY = configs.session.secret


def check_admin(request):
    if request.__user__ is None or not request.__user__.admin:
        raise APIPermissionError()


def get_page_index(page_str):
    p = 1
    try:
        p = int(page_str)
    except ValueError as e:
        pass
    if p < 1:
        p = 1
    return p


def user2cookie(user, max_age):
    '''
    Generate cookie str by user.
    '''
    # build cookie string by: id-expires-sha1
    expires = str(int(time.time() + max_age))
    s = '%s-%s-%s-%s' % (user.id, user.passwd, expires, _COOKIE_KEY)
    L = [user.id, expires, hashlib.sha1(s.encode('utf-8')).hexdigest()]
    return '-'.join(L)


def text2html(text):
    lines = map(lambda s: '<p>%s</p>' % s.replace('&', '&amp;').replace('<',
                                                                        '&lt;').replace('>', '&gt;'), filter(lambda s: s.strip() != '', text.split('\n')))
    return ''.join(lines)


@asyncio.coroutine
def cookie2user(cookie_str):
    '''
    Parse cookie and load user if cookie is valid.
    '''
    if not cookie_str:
        return None
    try:
        L = cookie_str.split('-')
        if len(L) != 3:
            return None
        uid, expires, sha1 = L
        if int(expires) < time.time():
            return None
        user = yield from User.find(uid)
        if user is None:
            return None
        s = '%s-%s-%s-%s' % (uid, user.passwd, expires, _COOKIE_KEY)
        if sha1 != hashlib.sha1(s.encode('utf-8')).hexdigest():
            logging.info('invalid sha1')
            return None
        user.passwd = '******'
        return user
    except Exception as e:
        logging.exception(e)
        return None


@get('/')
def index(*, page='1'):
    page_index = get_page_index(page)
    num = yield from Blog.findNumber('count(id)')
    page = Page(num)
    if num == 0:
        blogs = []
    else:
        blogs = yield from Blog.findAll(orderBy='created_at desc', limit=(page.offset, page.limit))
    return {
        '__template__': 'blogs.html',
        'page': page,
        'blogs': blogs
    }


@get('/blog/{id}')
def get_blog(id):
    blog = yield from Blog.find(id)
    comments = yield from Comment.findAll('blog_id=?', [id], orderBy='created_at desc')
    for c in comments:
        c.html_content = text2html(c.content)
    blog.html_content = markdown2.markdown(blog.content)
    return {
        '__template__': 'blog.html',
        'blog': blog,
        'comments': comments
    }


@get('/register')
def register():
    return {
        '__template__': 'register.html'
    }

# 后台管理
@get('/wx/admin')
def admin():
    return {
        '__template__': 'admin.html'
    }

# 邀请者初始页面
@get('/wx/kanjiahuodong/{pwd_code}')
def register(pwd_code):
    pwds = yield from UniquePwd.findAll('code=?', [pwd_code])
    if len(pwds) == 0:
        raise APIValueError('code', 'code 不存在')
    pwds = pwds[0]
    if pwds.wx_order_id:
        orders = yield from WxOrder.findAll('id=?', [pwds.wx_order_id])
        if len(orders) == 0:
            raise APIValueError('code', '订单不存在')
        order = orders[0]
        if order.people_amount != "":
            joiners = yield from WxJoiner.findAll("order_id=?", order.id)
            if len(joiners) == 0:
                raise APIValueError('code', '查询助力者失败')

        return {
            '__template__': 'kanjiahuodong.html',
            "pwd_code": pwd_code,
            "old_price": order.old_price,
            "sub_amount": order.sub_amount,
            "people_amount": order.people_amount,
            "store_name": order.store_name,
        }

    else:
        return {
            '__template__': 'kanjiahuodong.html',
            "pwd_code": pwd_code
        }

@get('/scanning/qr_code/{pwd_code}')
def scanning(*,pwd_code, code):
    global WxJoiner
    global UniquePwd
    
    pwds = yield from UniquePwd.findAll('code=?', [pwd_code])
    if len(pwds) == 0:
        raise APIValueError('code', 'code 不存在')
    pwds = pwds[0]
    if pwds.wx_order_id !="":
        appid = "wx65b975e308c72245"
        secret = "bf01504ce43d019e757b3183bdef9cad"
        # 获取access_token
        url_for_token = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=" + appid + \
            "&secret="+secret+"&code=" + code+"&grant_type=authorization_code"
        r = requests.get(url_for_token)
        rsp = json.loads(r.content)
        access_token = rsp["access_token"]
        # 获取userinfo
        url_for_user = "https://api.weixin.qq.com/sns/userinfo?access_token="+access_token +  \
            "&openid="+rsp["openid"]+"&lang=zh_CN"

        r_for_user = requests.get(url_for_user)
        r_for_user = json.loads(r_for_user.content)

        joiners = yield from WxJoiner.findAll('order_id=?', [pwds.wx_order_id])

        if len(joiners) > 0:
            for v in joiners:
                if v.user_id == r_for_user["openid"]:
                    raise APIValueError('code', '你已经参与过此砍价活动，请勿重复参加')

        amount = 20
        sum = 0
        if len(joiners) == 0:
            amount = 50
        elif len(joiners) == 1:
            amount = 30
        else:
            for v in joiners:
                sum+=v.help_amount
        if sum >= 200:
            raise APIValueError('code', '此轮砍价已经结束')

        wxJoiner = WxJoiner(
            order_id=pwds.wx_order_id,
            user_id=r_for_user["openid"],
            wx_user_name=r_for_user["nickname"],
            wx_user_image=r_for_user["headimgurl"],
            wx_addr=r_for_user["country"]+"-"+r_for_user["city"],
            wx_sex=r_for_user["sex"],
            create_at=time.time(),
            help_amount= if 
        )
        yield from wxJoiner.save()

    return {
        'help_amount':wxJoiner.help_amount,
        '__template__': 'kanjiahuodong.html',
    }

# 获取微信用户信息
@get('/wx/wechart_user')
def wechart_user(*,pwd_code, code):
    global UniquePwd
    pwds = yield from UniquePwd.findAll('code=?', [pwd_code])
    if len(pwds) == 0:
        raise APIValueError('code', 'code 不存在')
    pwds = pwds[0]
    if pwds.wx_order_id != "":
        appid = "wx65b975e308c72245"
        secret = "bf01504ce43d019e757b3183bdef9cad"
        # 获取access_token
        url_for_token = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=" + appid + \
            "&secret="+secret+"&code=" + code+"&grant_type=authorization_code"
        r = requests.get(url_for_token)
        rsp = json.loads(r.content)
        access_token = rsp["access_token"]
        # 获取userinfo
        url_for_user = "https://api.weixin.qq.com/sns/userinfo?access_token="+access_token +  \
            "&openid="+rsp["openid"]+"&lang=zh_CN"

        r_for_user = requests.get(url_for_user)
        r_for_user = json.loads(r_for_user.content)
        # r_for_user["access_token"] = access_token
        # 获取access_token，用于js-sdk （此access_token跟上面的access_token不一样）
        # url_for_access_token = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + \
        #     appid+"&secret="+secret
        # r_for_access_token = requests.get(url_for_access_token)
        # r_for_access_token = json.loads(r_for_access_token.content)
        # 获取js_sdk
        # url_for_js_sdk = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=" + \
        #     r_for_access_token["access_token"]+"&type=jsapi"
        # r_for_js_sdk = requests.get(url_for_js_sdk)
        # r_for_js_sdk = json.loads(r_for_js_sdk.content)

        # rsp = {}

        wxOrder = WxOrder(
            wx_user_id=r_for_user["openid"],
            wx_user_name=r_for_user["nickname"],
            wx_user_image=r_for_user["headimgurl"],
            wx_addr=r_for_user["country"]+"-"+r_for_user["city"],
            wx_sex=r_for_user["sex"],
            create_at=time.time(),
        )
        yield from wxOrder.save()
        pwds.is_used = 1
        pwds.wx_order_id = wxOrder.id
        yield from pwds.update()
        return r_for_user

@get('/signin')
def signin():
    return {
        '__template__': 'signin.html'
    }


@post('/api/authenticate')
def authenticate(*, email, passwd):
    if not email:
        raise APIValueError('email', 'Invalid email.')
    if not passwd:
        raise APIValueError('passwd', 'Invalid password.')
    users = yield from User.findAll('email=?', [email])
    if len(users) == 0:
        raise APIValueError('email', 'Email not exist.')
    user = users[0]
    # check passwd:
    sha1 = hashlib.sha1()
    sha1.update(user.id.encode('utf-8'))
    sha1.update(b':')
    sha1.update(passwd.encode('utf-8'))
    if user.passwd != sha1.hexdigest():
        raise APIValueError('passwd', 'Invalid password.')
    # authenticate ok, set cookie:
    r = web.Response()
    r.set_cookie(COOKIE_NAME, user2cookie(
        user, 86400), max_age=86400, httponly=True)
    user.passwd = '******'
    r.content_type = 'application/json'
    r.body = json.dumps(user, ensure_ascii=False).encode('utf-8')
    return r


@get('/signout')
def signout(request):
    referer = request.headers.get('Referer')
    r = web.HTTPFound(referer or '/')
    r.set_cookie(COOKIE_NAME, '-deleted-', max_age=0, httponly=True)
    logging.info('user signed out.')
    return r


@get('/manage/')
def manage():
    return 'redirect:/manage/comments'


@get('/manage/comments')
def manage_comments(*, page='1'):
    return {
        '__template__': 'manage_comments.html',
        'page_index': get_page_index(page)
    }


@get('/manage/blogs')
def manage_blogs(*, page='1'):
    return {
        '__template__': 'manage_blogs.html',
        'page_index': get_page_index(page)
    }


@get('/manage/blogs/create')
def manage_create_blog():
    return {
        '__template__': 'manage_blog_edit.html',
        'id': '',
        'action': '/api/blogs'
    }


@get('/manage/blogs/edit')
def manage_edit_blog(*, id):
    return {
        '__template__': 'manage_blog_edit.html',
        'id': id,
        'action': '/api/blogs/%s' % id
    }


@get('/manage/users')
def manage_users(*, page='1'):
    return {
        '__template__': 'manage_users.html',
        'page_index': get_page_index(page)
    }


@get('/api/comments')
def api_comments(*, page='1'):
    page_index = get_page_index(page)
    num = yield from Comment.findNumber('count(id)')
    p = Page(num, page_index)
    if num == 0:
        return dict(page=p, comments=())
    comments = yield from Comment.findAll(orderBy='created_at desc', limit=(p.offset, p.limit))
    return dict(page=p, comments=comments)


@post('/api/blogs/{id}/comments')
def api_create_comment(id, request, *, content):
    user = request.__user__
    if user is None:
        raise APIPermissionError('Please signin first.')
    if not content or not content.strip():
        raise APIValueError('content')
    blog = yield from Blog.find(id)
    if blog is None:
        raise APIResourceNotFoundError('Blog')
    comment = Comment(blog_id=blog.id, user_id=user.id, user_name=user.name,
                      user_image=user.image, content=content.strip())
    yield from comment.save()
    return comment


@post('/api/comments/{id}/delete')
def api_delete_comments(id, request):
    check_admin(request)
    c = yield from Comment.find(id)
    if c is None:
        raise APIResourceNotFoundError('Comment')
    yield from c.remove()
    return dict(id=id)


@get('/wx/unique_pwds')
def list():
    c = yield from UniquePwd.findAll(orderBy='id asc', limit=(0, 100))
    return dict(list=c)

@get('/api/users')
def api_get_users(*, page='1'):
    page_index = get_page_index(page)
    num = yield from User.findNumber('count(id)')
    p = Page(num, page_index)
    if num == 0:
        return dict(page=p, users=())
    users = yield from User.findAll(orderBy='created_at desc', limit=(p.offset, p.limit))
    for u in users:
        u.passwd = '******'
    return dict(page=p, users=users)


_RE_EMAIL = re.compile(
    r'^[a-z0-9\.\-\_]+\@[a-z0-9\-\_]+(\.[a-z0-9\-\_]+){1,4}$')
_RE_SHA1 = re.compile(r'^[0-9a-f]{40}$')


@post('/api/users')
def api_register_user(*, email, name, passwd):
    if not name or not name.strip():
        raise APIValueError('name')
    if not email or not _RE_EMAIL.match(email):
        raise APIValueError('email')
    if not passwd or not _RE_SHA1.match(passwd):
        raise APIValueError('passwd')
    users = yield from User.findAll('email=?', [email])
    if len(users) > 0:
        raise APIError('register:failed', 'email', 'Email is already in use.')
    uid = next_id()
    sha1_passwd = '%s:%s' % (uid, passwd)
    user = User(id=uid, name=name.strip(), email=email, passwd=hashlib.sha1(sha1_passwd.encode('utf-8')).hexdigest(),
                image='http://www.gravatar.com/avatar/%s?d=mm&s=120' % hashlib.md5(email.encode('utf-8')).hexdigest())
    yield from user.save()
    # make session cookie:
    r = web.Response()
    r.set_cookie(COOKIE_NAME, user2cookie(
        user, 86400), max_age=86400, httponly=True)
    user.passwd = '******'
    r.content_type = 'application/json'
    r.body = json.dumps(user, ensure_ascii=False).encode('utf-8')
    return r


@get('/api/blogs')
def api_blogs(*, page='1'):
    page_index = get_page_index(page)
    num = yield from Blog.findNumber('count(id)')
    p = Page(num, page_index)
    if num == 0:
        return dict(page=p, blogs=())
    blogs = yield from Blog.findAll(orderBy='created_at desc', limit=(p.offset, p.limit))
    return dict(page=p, blogs=blogs)


@get('/api/blogs/{id}')
def api_get_blog(*, id):
    blog = yield from Blog.find(id)
    return blog


@post('/api/blogs')
def api_create_blog(request, *, name, summary, content):
    check_admin(request)
    if not name or not name.strip():
        raise APIValueError('name', 'name cannot be empty.')
    if not summary or not summary.strip():
        raise APIValueError('summary', 'summary cannot be empty.')
    if not content or not content.strip():
        raise APIValueError('content', 'content cannot be empty.')
    blog = Blog(user_id=request.__user__.id, user_name=request.__user__.name,
                user_image=request.__user__.image, name=name.strip(), summary=summary.strip(), content=content.strip())
    yield from blog.save()
    return blog


@post('/api/blogs/{id}')
def api_update_blog(id, request, *, name, summary, content):
    check_admin(request)
    blog = yield from Blog.find(id)
    if not name or not name.strip():
        raise APIValueError('name', 'name cannot be empty.')
    if not summary or not summary.strip():
        raise APIValueError('summary', 'summary cannot be empty.')
    if not content or not content.strip():
        raise APIValueError('content', 'content cannot be empty.')
    blog.name = name.strip()
    blog.summary = summary.strip()
    blog.content = content.strip()
    yield from blog.update()
    return blog


@post('/api/blogs/{id}/delete')
def api_delete_blog(request, *, id):
    check_admin(request)
    blog = yield from Blog.find(id)
    yield from blog.remove()
    return dict(id=id)
