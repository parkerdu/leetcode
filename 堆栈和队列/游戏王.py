import sys
input('请输入信息：')
def input_data():

    list = []
    try:
        while True:
            line = sys.stdin.readline().strip()
            if line == '':
                break
            lines = line.split()
            if len(lines) == 1:
                list.append([int(lines[0])])
            else:
                list.append([int(lines[0]), int(lines[1])])
    except:
        pass
    print(list)
    return list


def max_value():
    nums = input_data()
    num_case = nums[0][0]
    first = 1
    results = []
    for i in range(num_case):
        last = first + nums[first][0] + 1
        time = nums[first][1]
        result = dp(nums[first:last], time)
        first = last
        print(first)
        results.append(result)
    return results

def dp(data,max_time):
    n = len(data)
    dp = [0] * (n+1)
    for i in range(n+1):
        dp[i] = [0] * (max_time + 1)
    for i in range(1,n+1):
        for j in range(1,max_time+1):
            if data[i-1][1] > j:
                dp[i][j] = dp[i-1][j]
            else:
                dp[i][j] = max(data[i-1][0] + dp[i-1][j-data[i-1][1]], dp[i-1][j])
    return dp[-1][-1]



a = max_value()
for i in a:
    print(i)

