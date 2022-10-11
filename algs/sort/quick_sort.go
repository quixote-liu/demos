package sort

func QuickSort(array []int) {
	quickSort(array, 0, len(array)-1)
}

func quickSort(array []int, begin, end int) {
	if begin < end {
		// 进行切分
		loc := partition(array, begin, end)

		// 对左部分进行快排
		quickSort(array, begin, loc-1)

		// 对右部分进行快排
		quickSort(array, loc+1, end)
	}
}

// 切分函数，并返回切分元素的下标
func partition(array []int, begin, end int) int {
	n := begin + 1 // 将array[begin]作为基准数，因此从array[begin+1]开始
	m := end       // array[end]是数组的最后一位

	// 没有重合之前
	for n < m {
		// 如果第一个数大于基准书
		if array[n] > array[begin] {
			// 替换第一个数和第二个数，并将第二数左移
			array[n], array[m] = array[m], array[n]
			m--
		} else {
			// 第一个数右移
			n++
		}
	}

	// 重合之后，数组被切分成两部分：
	// 前面一部分都小于等于基准数，后面一部分都大于基准数

	// 获取切分的那个下标点
	// 这里必须取>=，否则如果数组中的元素全部相同时会出现问题
	if array[n] >= array[begin] {
		n--
	}
	array[n], array[begin] = array[begin], array[n]
	return n
}
