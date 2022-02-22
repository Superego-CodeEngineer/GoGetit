package insertsort

func insertsort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	for tail := 1; tail < length; tail++ {
		val := arr[tail]
		i := tail - 1
		for ; i >= 0; i-- {
			if arr[i] > val {
				arr[i+1] = arr[i]
			} else {
				break
			}
		}
		arr[i+1] = val
	}
	return arr
}
