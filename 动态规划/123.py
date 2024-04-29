data = list(map(int,input().split()))

def xiaomi1(prices):
    max_total_profit = 0
    first_profits = [0] * len(prices)
    min_price = float('inf')

    for i in range(len(prices)):
        min_price = min(min_price, prices[i])
        profit = prices[i] - min_price
        max_total_profit = max(max_total_profit, profit)
        first_profits[i] = max_total_profit

    max_price = float('-inf')

    for j in range(len(prices) - 1, 0, -1):
        max_price = max(max_price, prices[j])
        profit = max_price - prices[j]
        max_total_profit = max(max_total_profit, first_profits[j - 1] + profit)

    print(max_total_profit)

# a = [2,1 ,5, 0, 2, 3, 1, 4]
xiaomi1(data)