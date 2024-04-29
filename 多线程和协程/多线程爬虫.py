import requests
import asyncio
import threading
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


def get(url):
    try:
        r = requests.get(url, allow_redirects=False, timeout=2.0, headers=headers)

    except:
        pass
    else:
        print(r.status_code, r.url)
        print('当前的线程为：{}'.format(threading.currentThread()))

print(u'多线程抓取')
# 创建8个线程来爬取上面的url
threads = []
t1 = time.time()
for url in urls:
    # 共有8个url,每个url都创建一个线程来运行
    # target传入函数对象，args传入函数所需的参数
    get_thread = threading.Thread(target=get,args=[url])
    threads.append(get_thread)
    get_thread.start()

# 等待所有线程结束，join()方法将会阻塞，直到该线程完成

for t in threads:
    t.join()
print(time.time() - t1)
