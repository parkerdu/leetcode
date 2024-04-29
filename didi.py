# n = list(map(int,input().split(" ")))
# data = list(input().split(" "))
n = 6
data = ['3','+','2','+','1','*','-4','*','-5','+','1']
length = len(data)
def didi1(data):

    can_not = []
    for idx,s in enumerate(data):
        if s == '-' or s == '/':
            can_not.append(idx-1)
            can_not.append(idx+1)
        elif s == '*':
            # if int(data[idx-1]) > int(data[idx+1]):
            #     data[idx - 1],data[idx+1] = data[idx+1],data[idx - 1]
            can_not.append(idx - 1)
            can_not.append(idx + 1)
    for i in range(length):
        j = i
        while j < length-2:
            if data[j + 1] == '+' and j + 2 not in can_not and i not in can_not:
                if int(data[j + 2]) < int(data[i]):
                    data[i], data[j + 2] = data[j + 2], data[i]
            elif data[j + 1] == '* ' and j + 2 not in can_not and i not in can_not:
                if int(data[j + 2]) < int(data[i]):
                    data[i], data[j + 2] = data[j + 2], data[i]
            else:
                break
            j += 2
    # for i in range(length-1,-1,-1):
    #     j = i
    #     while j:
    #         if data[j-1] == '+' and j-2 not in can_not and j not in can_not :
    #             if int(data[j-2]) > int(data[j]):
    #                 data[j],data[j-2] = data[j-2], data[j]
    #         else:
    #             break
    #         j -= 2

    for i in data:
        print(i,end=' ')

didi1(data)





