
arr = input().split(" ")
# arr = '8123456A'
n, arr = arr[0], arr[1:]
result = []
for val in arr:
    if val == 'A':
        result += ['12', '34']
    elif val == 'B':
        result += ['AB', 'CD']
    else:
        result.append(val)
result = [hex(len(result)+1).upper()[2:]] + result
for val in result:
    print(val, end=' ')


