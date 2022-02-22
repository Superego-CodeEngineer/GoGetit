package selectsort

func selectsort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	for tail := 0; tail < length; tail++ {
		pos := tail
		for i := tail + 1; i < length; i++ {
			if arr[i] < arr[pos] {
				pos = i
			}
		}
		arr[tail], arr[pos] = arr[pos], arr[tail]
	}
	return arr
}
