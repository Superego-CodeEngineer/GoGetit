package heapsort

func HeadSort(arr []int) []int {
	arrLen := len(arr)
	if arrLen <= 1 {
		return arr
	}

	buildheap(arr, arrLen-1)

	tail := arrLen - 1
	for tail > 0 {
		arr[0], arr[tail] = arr[tail], arr[0]
		tail--
		heapify(arr, tail, 0)
	}

	return arr
}

func buildheap(arr []int, tail int) {
	for i := tail / 2; i >= 0; i-- {
		heapify(arr, tail, i)
	}
}

func heapify(arr []int, tail, i int) {
	maxPos := i
	for {
		if i*2 <= tail && arr[i] < arr[i*2] {
			maxPos = i * 2
		}
		if i*2+1 <= tail && arr[maxPos] < arr[i*2+1] {
			maxPos = i*2 + 1
		}
		if maxPos == i {
			return
		}
		arr[i], arr[maxPos] = arr[maxPos], arr[i]
		i = maxPos
	}
}
