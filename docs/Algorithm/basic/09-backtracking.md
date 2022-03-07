# 9. Backtracking - 回溯

- [9. Backtracking - 回溯](#9-backtracking---回溯)
	- [Go - 经典的八皇后问题](#go---经典的八皇后问题)
	- [Go - 回溯的核心思想](#go---回溯的核心思想)
	- [Go - 回溯算法的应用](#go---回溯算法的应用)

谈到回溯算法，就一定要讨论最著名的八皇后问题，我会在文首给出答案，希望你能在理解回溯算法后回到这里看懂它：

## Go - 经典的八皇后问题

八皇后问题是一个以国际象棋为背景的问题：如何能够在 8×8 的国际象棋棋盘上放置八个皇后，使得任何一个皇后都无法直接吃掉其他的皇后？

为了达到此目的，任两个皇后都不能处于同一条横行、纵行或斜线上

![Figure.1 8queens](/image/Algorithm/basic/09-backtracking/8queens.png)

```Golang
var result [8]int

func cal8queens(row int) {
	if row == 8 {
		printQuees(result)
		return
	}

	for col := 0; col < 8; col++ {
		if isOK(row, col) {
			result[row] = col
			cal8queens(row + 1)
		}
	}
}

func isOK(row, col int) bool {
	leftup, rightup := col-1, col+1
	for i := row - 1; i >= 0; i-- {
		if result[i] == col {
			return false
		}
		if leftup >= 0 && result[i] == leftup {
			return false
		}
		if rightup < 8 && result[i] == rightup {
			return false
		}
		leftup--
		rightup++
	}
	return true
}
```

## Go - 回溯的核心思想

回溯就像时间宝石，能够支撑你：

- 枚举所有的可能满足预期的选择
- 在碰到不符合预期的解时倒退至上一个选择
- 直到选择符合预期的解

## Go - 回溯算法的应用

我们已在 [Go - 贪心算法的应用](07-greedy.md#go---贪心算法的应用) 探讨过 0-1 背包问题，现在，再次从该问题中理解回溯算法：

场景二：物品是不可切割（或分块）的

背包能装 10kg 物品，从商店中选购如下几种零食：
  1. s1: 10 RMB, 2kg
  2. s2: 8 RMB, 5kg
  3. s3: 20 RMB, 9kg
  4. s4: 5 RMB, 3kg

如何购置才能让背包所装糖果的总值最高？

```Golang
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
```
