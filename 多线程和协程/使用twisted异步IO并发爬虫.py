"""
https://www.cnblogs.com/dion-90/articles/8451842.html
这篇博客有详细的各种异步爬取的实例代码
"""
from twisted.web.client import Agent, defer
from twisted.internet import reactor
from twisted.web.http_headers import Headers
HEADERS = {'Accept': ['text/html,application/xhtml+xml,application/xml;q=0.9'],
   'Accept-Language': ['zh-CN,zh;q=0.8'],
   'Accept-Encoding': ['gzip, deflate'],}

URLS = ['http://www.cnblogs.com/moodlxs/p/3248890.html',
        'http://blog.csdn.net/yueguanghaidao/article/details/24281751',
        'https://my.oschina.net/visualgui823/blog/36987',
        'http://blog.chinaunix.net/uid-9162199-id-4738168.html',
        'http://www.tuicool.com/articles/u67Bz26',
        'http://rfyiamcool.blog.51cto.com/1030776/1538367/',
        'http://itindex.net/detail/26512-flask-tornado-gevent',]
urls = URLS
headers = HEADERS
headers['user-agent'] = ["Mozilla/5.0+(Windows+NT+6.2;+WOW64)+AppleWebKit/537.36+(KHTML,+like+Gecko)+Chrome/45.0.2454.101+Safari/537.36"]



def all_done(arg):
    reactor.stop()


def callback(response):
    # print('Response version:', response.version)
    print('Response code:', response.code)
    # print('Response phrase:', response.phrase)


deferred_list = []
agent = Agent(reactor)
url_list = ['http://www.bing.com', 'http://www.baidu.com', ]
for url in url_list:

    deferred = agent.request(b'GET',bytes(url,encoding='utf-8'),Headers(headers))
    deferred.addCallback(callback)
    deferred_list.append(deferred)

dlist = defer.DeferredList(deferred_list)
dlist.addBoth(all_done)

reactor.run()