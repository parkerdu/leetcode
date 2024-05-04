"""
控制台输入数据，详见https://exercise.acmcoder.com/online/online_judge_ques?ques_id=9579&konwledgeId=137&opencustomeinput=true
3 1
2 3 1
5 4
1 2 1
3 4 0
2 5 1
3 2 1
下面的代码可以把所有行的数据存入一个二维列表
如果想打印输出的话比如打印 3 1
print(str(3)+' '+str(1))
"""
input('输入:')
def get_data():
    data = []
    try:
        while True:
            n = list(map(int, input().split(" ")))
            data.append(n)
            for i in range(n[0]):
                abc = list(map(int,input().split(" ")))
                data.append(abc)
    except:
        pass
    return data
print(data)