import asyncio
import aiohttp
import time

HEADERS = {'Accept': 'text/html,application/xhtml+xml,application/xml;q=0.9',
   'Accept-Language': 'zh-CN,zh;q=0.8',
   'Accept-Encoding': 'gzip, deflate',}

URLS = ['http://www.cnblogs.com/moodlxs/p/3248890.html',
        'https://www.zhihu.com/topic/19804387/newest',
        'http://blog.csdn.net/yueguanghaidao/article/details/24281751',
        'https://my.oschina.net/visualgui823/blog/36987',
        'http://blog.chinaunix.net/uid-9162199-id-4738168.html',
        'http://www.tuicool.com/articles/u67Bz26',
        'http://rfyiamcool.blog.51cto.com/1030776/1538367/',
        'http://itindex.net/detail/26512-flask-tornado-gevent']
urls = URLS
headers = HEADERS
headers['user-agent'] = "Mozilla/5.0+(Windows+NT+6.2;+WOW64)+AppleWebKit/537.36+(KHTML,+like+Gecko)+Chrome/45.0.2454.101+Safari/537.36"


async def get(session, url):
    """
    两种方法获得reponse
    :param session:
    :param url:
    :return:
    """
    # 第一种写法
    # async with session.get(url) as response:
    #     return response

    # 第二种写法
    response = await session.get(url)
    return response


async def main():
    t1 = time.time()
    async with aiohttp.ClientSession() as session:
        tasks = [asyncio.create_task(get(session, url)) for url in urls]
        await asyncio.wait(tasks)
        respones = await asyncio.gather(*tasks)
        for r in respones:
            print(r.status, r.url)
    t2 = time.time() - t1
    print(t2)

asyncio.run(main())