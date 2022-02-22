package countingsort

func CountingSort(arr []int) []int {
	arrLen := len(arr)
	if arrLen <= 1 {
		return arr
	}

	max := arr[0]
	for _, val := range arr {
		if val > max {
			max = val
		}
	}

	c := make([]int, max+1)

	for _, val := range arr {
		c[val]++
	}

	for i := 1; i <= max; i++ {
		c[i] = c[i-1] + c[i]
	}

	r := make([]int, arrLen)
	for i := max - 1; i >= 0; i-- {
		index := c[arr[i]] - 1
		r[index] = arr[i]
		c[arr[i]]--
	}

	copy(arr, r)

	return arr
}
