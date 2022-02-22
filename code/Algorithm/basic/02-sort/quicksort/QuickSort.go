package quicksort

func QuickSort(arr []int) []int {
	arrLen := len(arr)
	if arrLen <= 1 {
		return arr
	}

	quicksort(arr, 0, arrLen-1)
	return arr
}

func quicksort(arr []int, head, tail int) {
	if head >= tail {
		return
	}

	mid := partition(arr, head, tail)
	quicksort(arr, head, mid)
	quicksort(arr, mid+1, tail)
}

func partition(arr []int, head, tail int) int {
	pivot := arr[tail]
	i := head
	for j := head; j < tail; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[tail] = arr[tail], arr[i]
	return head
}
