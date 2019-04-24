class Solution(object):
    def __init__(self):
        pass

    def allCellsDistOrder(self, R, C, r0, c0):
        """思想：
            距离d取值为d=0,1,2,3,...,max_distance
            所以把相同距离的坐标都存入同一个列表
            然后再一个一个取出来就是按距离排好序的，因为对于相同距离的坐标不分先后"""
        cell_dict = []
        max_distance = max((r0+c0),(R-1-r0+C-c0),(r0+C-1-c0),(c0+R-1-r0))
        for i in range(max_distance+1):
            cell_dict.append([])
        for r in range(R):
            for c in range(C):
                d = abs(r-r0)+abs(c-c0)
                # 这里如果有相同的距离，键相同字典就不存了
                cell_dict[d].append([r,c])
        print(cell_dict)
        sort_cell = []
        for i in cell_dict:
            for j in i:
                sort_cell.append(j)
        # 对字典按值排序
        # sorted(cell_dict.keys())
        # 取出排序后所有的值，并转为列表
        # sort_cell = list(cell_dict.values())
        return sort_cell

    def allCellsDistOrder1(self, R, C, r0, c0):
        """法二：
            step1:先遍历所有的坐标对，并存入一个列表中
            step2:对该列表按cell距离作为权值进行排序
            """
        arr = []
        for i in range(R):
            for j in range(C):
                arr.append([i, j])
        print(arr)
        def key_value(item):
            # 按距离作为权值进行排序
            return abs(item[0]-r0)+abs(item[1]-c0)
        return sorted(arr, key=key_value)
        # 也可省略上述的key_value函数，换成下面的一行形式，其中p为arr列表中一个元素也即[r,c]
        # return sorted(arr, key=lambda p: abs(p[0] - r0) + abs(p[1] - c0))
if __name__ == "__main__":
    su = Solution()
    cell = su.allCellsDistOrder1(2,2,0,1)
    print(cell)
    list1 = [7, -8, 5, 4, 0, -2, -5]
    print(sorted(list1, key=lambda x: (x < 0)))
