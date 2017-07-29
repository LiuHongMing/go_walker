package qsort

func quickSort(arr []int, left int, right int) {
	pivot := arr[left]  // 基准元素值
	p := left           // 基准元素坐标
	i, j := left, right // 探测哨兵

	// 当哨兵未相遇的时候
	for i != j {
		// 先从右边开始找小于基准元素的值
		if j >= p && arr[j] >= pivot {
			j--
		}
		if j >= p {
			arr[p] = arr[j]
			p = j
		}

		// 再从左边开始找大于基准元素的值
		if i <= p && arr[i] <= pivot {
			i++
		}
		if i <= p {
			arr[p] = arr[i]
			p = i
		}
	}
	arr[p] = pivot
	// 如果左边的元素大于2个以上才需要进行再划分
	if p-left > 1 {
		quickSort(arr, left, p-1)
	}
	// 如果右边的元素大于2个以上才需要进行再划分
	if right-p > 1 {
		quickSort(arr, p+1, right)
	}
}

func QuickSort(values []int) {
	quickSort(values, 0, len(values)-1)
}
