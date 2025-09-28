package array

// 快速排序
// 快速排序的思想， 找到两个指针左右的去比较并且交互
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	p := arr[0]
	left, right := 1, len(arr)-1

	for left <= right {
		for left <= right && arr[left] < p {
			left++
		}

		for left <= right && arr[right] > p {
			right--
		}

		for left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

	arr[0], arr[right] = arr[right], arr[0]

	quickSort(arr[:right])
	quickSort(arr[left:])

	return arr
}

func QuickSortV2(arr []int, low, right int) {
	if low < right {
		index := partition(arr, low, right)
		QuickSortV2(arr, low, index-1)
		QuickSortV2(arr, index+1, right)
	}
}

func partition(arr []int, low, right int) int {
	index := arr[low]
	i := low + 1

	for j := i; j <= right; j++ {
		if arr[j] < index {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[low], arr[i-1] = arr[i-1], arr[low]
	return i - 1
}
