package dynamic

func Dynamic(prices, weights []int, left int) (value int) {
	// lenSnack 表示决策阶段
	lenSnack := len(prices)
	// 初始化 states:
	// states[][] 表示总价值
	states := make([][]int, lenSnack)
	for i := 0; i < lenSnack; i++ {
		states[i] = make([]int, left+1)
	}
	if weights[0] <= left {
		states[0][weights[0]] = prices[0]
	}
	for i := 1; i < lenSnack; i++ {
		// 不选择第 i 个物品
		for j := 0; j <= left; j++ {
			if states[i-1][j] != 0 {
				states[i][j] = states[i-1][j]
			}
		}
		// 选择第 i 个物品
		for j := 0; j <= left-weights[i]; j++ {
			if states[i-1][j] >= 0 && states[i-1][j]+prices[i] > states[i][j+weights[i]] {
				states[i][j+weights[i]] = states[i-1][j] + prices[i]
			}
		}
	}
	// 找到最大的值
	for _, val := range states[lenSnack-1] {
		if val > value {
			value = val
		}
	}
	return value
}
