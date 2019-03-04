package main

func guessNumber(n int) int {
	left, right := 0, n
	for left <= right {
		mid := left + (right-left)/2
		temp := guess(mid)
		if temp == 0 {
			return mid
		}
		if temp == 1 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func guess(num int) int {
	return 1
}

func main() {

}
