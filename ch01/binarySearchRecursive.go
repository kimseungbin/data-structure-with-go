package ch01

func BinarySearchRecursive(data []int, low, high, value int) int {
	mid := low + (high-low)/2
	if data[mid] == value {
		return mid
	} else if data[mid] < value {
		return BinarySearchRecursive(data, mid+1, high, value)
	} else {
		return BinarySearchRecursive(data, low, mid-1, value)
	}
}
