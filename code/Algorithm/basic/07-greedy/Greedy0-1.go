package greedy

// 找到单价最大的下标
func find_pos_max(prices []int) (pos_max int) {
	for i := 0; i < len(prices); i++ {
		if prices[pos_max] < prices[i] {
			pos_max = i
		}
	}
	return pos_max
}

func greedy(prices, stocks []int, bag int) (ans int) {
	pos_max, left := 0, bag

	find_pos_max(prices)

	for left != 0 {
		if stocks[pos_max] >= left {
			return ans + prices[pos_max]*left
		}
		ans += prices[pos_max] * stocks[pos_max]
		left -= stocks[pos_max]
		// 使用完存量后，只需要置零单价即可
		prices[pos_max] = 0
		pos_max = find_pos_max(prices)
	}

	return -1
}
