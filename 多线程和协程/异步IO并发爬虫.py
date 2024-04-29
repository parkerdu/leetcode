"""
利用python中的asyncio来实现并发爬取
"""
import asyncio
import requests
import time

HEADERS = {'Accept': 'text/html,application/xhtml+xml,application/xml;q=0.9',
   'Accept-Language': 'zh-CN,zh;q=0.8',
   'Accept-Encoding': 'gzip, deflate',}

URLS = ['http://www.cnblogs.com/moodlxs/p/3248890.html',

        'http://blog.csdn.net/yueguanghaidao/article/details/24281751',
        'https://my.oschina.net/visualgui823/blog/36987',
        'http://blog.chinaunix.net/uid-9162199-id-4738168.html',
        'http://www.tuicool.com/articles/u67Bz26',
        'http://rfyiamcool.blog.51cto.com/1030776/1538367/',
        'http://itindex.net/detail/26512-flask-tornado-gevent']
urls = URLS
headers = HEADERS
headers['user-agent'] = "Mozilla/5.0+(Windows+NT+6.2;+WOW64)+AppleWebKit/537.36+(KHTML,+like+Gecko)+Chrome/45.0.2454.101+Safari/537.36"


async def get(url):
    try:
        t1 = time.time()
        # 本来request.get不是一个可等待对象，所以下面这种直接在前面加await的方法不行
        # r = await requests.get(url, allow_redirects=False, timeout=2.0, headers=headers)

        # 但是可以通过下面的方法把用request请求的方法变成可等待对象
        loop = asyncio.get_event_loop()
        r = await loop.run_in_executor(None, requests.get, url)
        t2 = time.time() - t1
        print(t2)
    except:
        pass
    else:
        # try通过就会来这里运行
        print(r.status_code, r.url)


async def main():
    t1 = time.time()
    tasks = [asyncio.create_task(get(url)) for url in urls]
    await asyncio.wait(tasks)
    t2 = time.time() - t1
    print(t2)

asyncio.run(main())