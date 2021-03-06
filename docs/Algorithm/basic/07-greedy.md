# 7. Greedy - 贪心

> 掌握基础的数据结构和算法后，我还为你准备四种更加基本的、常用的算法：分治算法、回溯算法、贪心算法和动态规划
>
> 与其说是算法，不如说它们是算法思想，因为它们不是具体的算法，而是常作为我们设计具体算法的指导

- [7. Greedy - 贪心](#7-greedy---贪心)
	- [Go - 贪心算法的通常适用场景](#go---贪心算法的通常适用场景)
	- [Go - 贪心算法的核心思想](#go---贪心算法的核心思想)
	- [Go - 贪心算法与最优解的关系](#go---贪心算法与最优解的关系)
	- [Go - 贪心算法的应用](#go---贪心算法的应用)

贪心算法可以说是生活中最常见的一种思想

## Go - 贪心算法的通常适用场景

贪心算法的适用场景往往具有两个特征：

1. 同时包含限制值和期望值
2. 要求在满足限制值的情况下，期望值最大

## Go - 贪心算法的核心思想

顾名思义，该算法的核心在于如何去贪，实际上很简单：

- 每一步都保证期望值最大

但实际上，保证每一步期望值最大，并不能代表结果的期望值最大。
因此，贪心算法真正困难的地方在于如何平衡与最优解的关系

## Go - 贪心算法与最优解的关系

想要从理论上严格证明贪心算法的正确性是十分复杂的，
但从绝大部分实践的情况来看，贪心算法的正确性又是显而易见

假设我们有若干 10，9，1 面值的纸币，请问最少需要多少张纸币可以凑齐 36 元钱？

- 从贪心的角度：3 张 10 元和 6 张 1 元的纸币是最优解
- 但如果你对数字稍微敏感一点：4 张 9 元的纸币即可

因此，贪心算法并不能总是给出最优解

## Go - 贪心算法的应用

从 0-1 背包问题中理解贪心算法：

场景一：物品是可切割（或分块）的

背包能装 10kg 物品，从商店中选购如下几种散装的糖果：
  1. s1: 10 RMB/kg 存量 2kg
  2. s2: 2 RMB/kg 存量 5kg
  3. s3: 20 RMB/kg 存量 4kg
  4. s4: 5 RMB/kg 存量 8kg

如何购置才能让背包所装糖果的总值最高？

```Golang
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
```
