package main

func hIndex(citations []int) int {
	len1 := len(citations)
	left, right := 0, len1-1

	for left <= right {
		mid := left + (right-left)/2
		if citations[mid] >= (len1 - mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return len1 - left
}

func main() {

}
