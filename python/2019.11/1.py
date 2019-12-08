#!/usr/local/bin/python3
# -*- coding: utf-8 -*-

# 在计算机内存中，统一使用Unicode编码，
# 当需要保存到硬盘或者需要传输的时候，就转换为UTF-8编码。

# 用记事本编辑的时候，从文件读取的UTF-8字符被转换为Unicode
# 字符到内存里，编辑完成后，保存的时候再把Unicode转换为UTF-8保存到文件：

# 浏览网页的时候，服务器会把动态生成的
# Unicode内容转换为UTF-8再传输到浏览器：
print('hello,world',"nihao","dage")
a = 123
print(a)
a = 10/3
print(a)
a = 9/3
print(a)
a = 10//3
print(a)
print("你好")
print('你好')
print(ord('A'))
print(ord('中'))
print(chr(65))
print('\u4e2d\u6587')
# 由于Python的字符串类型是str，在内存中以Unicode表示，一个字符对应若干个字节。
# 如果要在网络上传输，或者保存到磁盘上，就需要把str变为以字节为单位的bytes
a=b'ABC'
print(a)
print('ABd'.encode('ascii'))
print('ABE你好'.encode('utf-8'))
print(b'ABCd'.decode('ascii'))
print(b'ABC'.decode('utf-8'))
print(b'\xe4\xb8\xad\xe6\x96\x87'.decode('utf-8'))
print(b'\xe4\xb8\xad\xff'.decode('utf-8',errors='ignore'))
print(len("你好".encode('utf-8')))
print('Hello, %s' % 'world')
print('你好，%s%s'%('日本','韩国'))
print('%1d-%3d' % (3,1))
print('age:%s gender:%s %%'%(25,True))
print('hello,{0},成绩提升了{1:.1f}%'.format('小明',17.125))
abs(7)

def myfunc(x):
    if x>=0:
        print("x大于0")
    else:
        print("x不大于0")
# myfunc("b")

def absx(x):
    if not isinstance(x,(int,float)):
        raise TypeError('bad operand type')
    if x >= 0:
        print("llll")
    return 1,1
x,y=absx(9)
print(x,y)

def power(x,n=2):
    s=1
    while n>0:
        s=s*x
        n=n-1
    return s

print(power(n=3,x=2))

def calc(*numbers):
    print(numbers)

calc(1,3,4,5,6)