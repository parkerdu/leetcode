"""
具体看 https://blog.csdn.net/daiyu__zz/article/details/81912018
和python文档   https://docs.python.org/zh-cn/3/library/asyncio-task.html
"""
import asyncio
import threading
import time

# 以下为老python版本的单线程并发执行--基于生成器的协程
# @asyncio.coroutine
# def hello():
#     print('Hello world! (%s)' % threading.currentThread())
#     # 异步调用asyncio.sleep(1)-->协程函数:
#     r = yield from asyncio.sleep(10)  #此处为另外一个协程，不是休眠
#     print('Hello again! (%s)' % threading.currentThread())
#
# # 获取EventLoop（事件循环器）:
# loop = asyncio.get_event_loop()
# # 执行coroutine
# tasks = [hello(), hello()]
# loop.run_until_complete(asyncio.wait(tasks))
# loop.close()


# 以下为python3.7中的async 和 await方法
async def hello(times):
    # threading.currentThread()可以告诉你当前的线程是啥，最后的结果可以看到两个协程task1,task2是在一个线程中运行的
    print('Hello world{}! ({})'.format(times, threading.currentThread()))
    # 异步调用asyncio.sleep(1)-->协程函数:
    # 当执行到下面一行行的await时候，程序不会在这里等待3s再去执行下一行，而是跳出该函数执行另一个协程也就是task2,等到
    # 3s之后会回到这里继续执行剩下的代码，所以你会看到先打印Hello world1!
    #                                           紧接着打印Hello world2!
    #                                   过了3s之后打印Hello again1!
    #                                           紧接着打印Hello again2!
    r = await asyncio.sleep(3)  #此处为另外一个协程，不是休眠
    print('Hello again{}! ({})'.format(times, threading.currentThread()))


async def main():
    """
    本来执行两次上面的hello()函数会延迟3*2=6s,但是最终的运行时间仅为3s，说明两个task是并发执行的
    :return:
    """
    print(f"started at {time.strftime('%X')}")
    task1 = asyncio.create_task(hello(1))
    task2 = asyncio.create_task(hello(2))
    await task1
    await task2
    print(f"finished at {time.strftime('%X')}")
asyncio.run(main())

