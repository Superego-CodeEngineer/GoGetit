# 3. Sort - 排序

- [3. Sort - 排序](#3-sort---排序)
	- [Go - 排序算法的指标](#go---排序算法的指标)
		- [Go - 排序算法的执行效率](#go---排序算法的执行效率)
		- [Go - 排序算法的内存消耗](#go---排序算法的内存消耗)
		- [Go - 排序算法的稳定性](#go---排序算法的稳定性)
	- [Go - 排序算法的分类](#go---排序算法的分类)
		- [Go - 时间复杂度 O(n<sup>2</sup>) 级](#go---时间复杂度-onsup2sup-级)
		- [Go - 时间复杂度 O(nlogn) 级](#go---时间复杂度-onlogn-级)
		- [Go - 时间复杂度 O(n) 级](#go---时间复杂度-on-级)

排序是生活中最常见的算法，可能现在你还没有接触过排序算法，但当你深入学习时，你发现它有种莫名的熟悉感。

排序算法数不胜数，短短一节肯定无法面面俱到，本节只会对最经典的排序算法进行介绍。

在介绍排序算法之前，我还需要向你简单说明一下排序算法的衡量指标。

## Go - 排序算法的指标

通常情况下，我们需要对排序算法的执行效率、内存消耗和稳定性进行比较：

### Go - 排序算法的执行效率

1. 最好、最坏、平均情况时间复杂度
2. 时间复杂度的系数、常数、低阶
3. 比较次数和移动（交换）次数 

现实中的排序算法应该充分考虑数据的特点，同样的算法，面对不同数据的表现往往是不同的。

### Go - 排序算法的内存消耗

- 原地排序：空间复杂度为 O(1) 的排序算法，即不需要额外的存储空间

需要消耗额外内存的排序算法往往更加稳定、效果更优，但是内存并不是无限的，面对 TB 级的数据量，非原地排序算法就会显得格外奢侈。

### Go - 排序算法的稳定性

- 前提条件：假设原序列中存在值相等的元素。
- 判断依据：值相同的元素在排序算法结束后，原有的先后关系是否发生变化。

面对复杂的业务，当需要使用多种排序算法时，具有稳定性的算法着能继承前一次排序的结果。

## Go - 排序算法的分类

接下来，我会按照排序算法的执行效率来介绍：

### Go - 时间复杂度 O(n<sup>2</sup>) 级

**冒泡排序**

实现步骤：
1. 检查原数组长度：
   - 大于 1 则需要排序，否则不需要排序，直接返回原函数
2. 遍历数组，对相邻的两个数据进行比较：
   - 每次只会对相邻的两个数据进行比较和交换
   - 经过一轮遍历，如果未交换过数据，说明此时数组已有序

```Golang
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
				tmp := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = tmp
			}
		}
		if !flag {
			break
		}
	}
	return arr
}
```

**插入排序**

实现步骤：
1. 检查原数组长度：
   - 大于 1 则需要排序，否则不需要排序，直接返回原函数
2. 将原数组视为两部分，已排序区间和未排序区间：
   - 在已排序区间中依次寻找插入数据的位置：
     - 大于插入数据则进行数据移动
     - 小于插入数据则跳出循环
   - 找到合适的位置插入数据

```Golang
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
```

**选择排序**

1. 检查原数组长度：
   - 大于 1 则需要排序，否则不需要排序，直接返回原函数
2. 将原数组视为两部分，已排序区间和未排序区间：
   - 在未排序区间中寻找值最小的数
   - 交换数据，使其添加到已排序区间的末尾

```Golang
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
		tmp := arr[tail]
		arr[tail] = arr[pos]
		arr[pos] = tmp
	}
	return arr
}
```

冒泡、插入、选择排序均是原地排序，但其中选择排序是不稳定的。

而将冒泡排序和插入排序进行比较，插入排序往往更受欢迎，虽然两者的元素移动次数都是原始数据的逆序度，但在进行数据交换的过程中，冒泡排序需要 3 个赋值操作，而插入排序仅需 1 个。

当然，插入排序还有优化空间，感兴趣的话，可以参考一下 [希尔排序](https://zh.wikipedia.org/wiki/%E5%B8%8C%E5%B0%94%E6%8E%92%E5%BA%8F)

### Go - 时间复杂度 O(nlogn) 级

**归并排序**

实现步骤：

1. 把原数组从中间分成前后两部分
2. 对前后两部分分别排序
3. 将排好序的两部分合并在一起
   - 创建临时数组暂存已排序数据
   - 临时数组覆盖原数组

```Golang
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
```

我们可以很明显的发现，归并排序另外为临时数组开辟了新的空间，因此它并不是原地排序。

我们以归并排序为例，分析一下它的空间复杂度：

1. 通过递推公式求解，整个归并过程需要的复杂度是 O(nlogn)
2. 但实际上，在合并完成后，临时开辟的内存空间会被释放掉，一个函数执行时，只会使用一个内存空间
3. 因此，临时内存空间最大不会超过 n 个数据的大小，因此，空间复杂度应该为 O(n)

**快速排序**

实现步骤：

1. 选取一个数作为 pivot
2. 将数组分为三部分：
   - 小于 pivot 的数放到左边
   - pivot 放在中间
   - 大于 pivot 的数放到右边
3. 分别对左右两个分区做相同的操作

```Golang
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
```

尽管归并和快排都采用了分治的思想，但它们的处理过程是截然相反的：

- 归并自下而上，先排序，再合并
- 快排自上而下，先分区，再排序

快排是一种不稳定的排序，但它巧妙的设计，使其成为原地排序，相比归并排序而言，较高的空间复杂度成为了它致命的缺陷，这也是快排更受欢迎的原因

当然，快排并不是完美的，在分区极度不均匀的情况下，快排的时间复杂度将会退化到 O(n<sup>2</sup>)

**堆排序**

实现步骤：

1. 建堆
   - 自上而下进行 **堆化**
2. 排序 
   - 从后向前排 

```Golang
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
```

同样是原地排序，但相比快排，堆排序的时间复杂度非常稳定 O(nlogn)

但堆排序的性能往往没有快排优秀：

1. 堆排序的数据访问方式对 CPU 缓存很不友好，它不如快排那般是局部顺序访问
2. 堆排序的数据交换次数往往高于快排，并且建堆的过程很容易降低其有序度


### Go - 时间复杂度 O(n) 级

线性排序的复杂度似乎是不可思议的，最主要的原因在于同之前介绍的排序相比，它是非基于比较的排序算法

**桶排序**

> 桶排序和计数排序十分相似，为了方便测试，我将以计数排序为例展示代码

核心思想：

1. 将数据分到 m 个有序的桶里
2. 对每个桶的数据单独进行排序
3. 排序结束后依次取出


**计数排序**

> 这是桶排序的一种特殊情况

```Golang
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

	return r
}
```

**基数排序**

核心思想：

1. 将原数据补齐到同一长度
2. 分割成位进行稳定排序

基数排序对数据的要求很特别：

1. 位之间存在递进关系
2. 每一位的数据范围要能满足线性排序的要求