package mergesort

func MergeSort(arr []int) []int {
	arrLen := len(arr)
	if arrLen <= 1 {
		return arr
	}

	mergesort(arr, 0, arrLen-1)
	return arr
}

func mergesort(arr []int, head, tail int) {
	if head >= tail {
		return
	}

	mid := (head + tail) / 2
	mergesort(arr, head, mid)
	mergesort(arr, mid+1, tail)
	merge(arr, head, mid, tail)
}

func merge(arr []int, head, mid, tail int) {
	tmpArr := make([]int, tail-head+1)

	i, j, k := head, mid+1, 0
	for ; i <= mid && j <= tail; k++ {
		if arr[i] <= arr[j] {
			tmpArr[k] = arr[i]
			i++
		} else {
			tmpArr[k] = arr[j]
			j++
		}
	}

	for ; i <= mid; i++ {
		tmpArr[k] = arr[i]
		k++
	}
	for ; j <= tail; j++ {
		tmpArr[k] = arr[j]
		k++
	}
	copy(arr[head:tail+1], tmpArr)
}
