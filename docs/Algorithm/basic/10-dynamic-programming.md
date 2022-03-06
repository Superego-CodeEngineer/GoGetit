# 10. Dynamic Programming - 动态规划

- [10. Dynamic Programming - 动态规划](#10-dynamic-programming---动态规划)
  - [Go - 动态规划的通用模型](#go---动态规划的通用模型)
  - [Go - 动态规划的三个特征](#go---动态规划的三个特征)
  - [Go - 动态规划的两种方法](#go---动态规划的两种方法)
  - [Go - 动态规划的应用](#go---动态规划的应用)

动态规划一般用来解决最优问题

## Go - 动态规划的通用模型

解决问题的过程中需要经历多个决策阶段，且每个决策阶段都对应着一组状态

## Go - 动态规划的三个特征

1. 最优子结构
2. 无后效性
3. 重复子问题

## Go - 动态规划的两种方法

1. 状态转移表法
2. 状态转移方程法

## Go - 动态规划的应用

认真思考过 [Go - 回溯算法的应用](09-backtracking.md#go---回溯算法的应用) 后，相信你会发现有很多子问题是重复的。
现在，我们来看看动态规划是如何优化 0-1 问题的：

```Golang

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
```

当然，回溯也可以使用在 [递归](02-recursion.md) 的优化方法——备忘表，以达到同样的目的
