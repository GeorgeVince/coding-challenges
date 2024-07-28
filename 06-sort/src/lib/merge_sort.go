package lib

func MergeSort(words []string) []string {
	return mergeSort(words, 0, len(words)-1)
}

func mergeSort(arr []string, left int, right int) []string {
	if left < right {
		mid := (left + right) / 2
		mergeSort(arr, left, mid)
		mergeSort(arr, mid+1, right)
		merge(arr, left, right)
	}
	return arr
}

func merge(arr []string, left int, right int) []string {
	mid := (left + right) / 2
	i := left
	j := mid + 1
	k := 0
	temp := make([]string, right-left+1)
	for i <= mid && j <= right {
		if arr[i] < arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			j++
		}
		k++
	}
	for i <= mid {
		temp[k] = arr[i]
		i++
		k++
	}
	for j <= right {
		temp[k] = arr[j]
		j++
		k++
	}
	for i = 0; i < k; i++ {
		arr[left+i] = temp[i]
	}
	return arr
}
