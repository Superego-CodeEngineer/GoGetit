package backtracking

var max = 0

func Backtracking(prices, weights []int, item, left, value int) int {
	// 背包剩余量为零或者已遍历所有物品
	if left == 0 || item == len(prices) {
		if max < value {
			max = value
		}
		return max
	}
	// 跳过当前物品
	max = Backtracking(prices, weights, item+1, left, value)
	// 在容量足够的情况下装下当前物品
	if left-weights[item] >= 0 {
		max = Backtracking(prices, weights, item+1, left-weights[item], value+prices[item])
	}

	return max
}
