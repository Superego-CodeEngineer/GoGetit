package bubblesort

func bubblesort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}

	for tail := length - 1; tail > 0; tail-- {
		var flag bool
		for i := 0; i < tail; i++ {
			if arr[i] > arr[i+1] {
				flag = true
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
		if !flag {
			break
		}
	}
	return arr
}
