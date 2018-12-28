package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type StrArr []int

func (arr StrArr) Len() int {
	return len(arr)
}

func (arr StrArr) Less(i, j int) bool {
	leni := maps[arr[i]]
	lenj := maps[arr[j]]

	tempi := arr[i]*pwoerOf10(lenj) + arr[j]
	tempj := arr[j]*pwoerOf10(leni) + arr[i]

	if tempi > tempj {
		return true
	}
	return false
}

func pwoerOf10(i int) int {
	res := 1
	for i > 0 {
		res *= 10
		i--
	}
	return res
}

func (arr StrArr) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

var maps map[int]int = make(map[int]int)

func largestNumber(nums []int) string {
	len1 := len(nums)
	if len1 == 0 {
		return ""
	}

	if len1 == 1 {
		return strconv.Itoa(nums[0])
	}

	allZero := true
	for i := 0; i < len1; i++ {
		if nums[i] != 0 {
			allZero = false
		}
	}
	if allZero == true {
		return "0"
	}

	for _, value := range nums {
		maps[value] = intLen(value)
	}

	sort.Sort(StrArr(nums))

	sb := strings.Builder{}

	for i := 0; i < len1; i++ {
		sb.WriteString(strconv.Itoa(nums[i]))
	}

	return sb.String()
}
func intLen(i int) int {
	if i == 0 {
		return 1
	}
	res := 0

	for i > 0 {
		res++
		i /= 10
	}
	return res
}

func main() {

	fmt.Println(pwoerOf10(3))

	var nums []int

	nums = []int{0, 1}
	fmt.Println(largestNumber(nums))

	nums = []int{10, 2}
	fmt.Println(largestNumber(nums))
	nums = []int{121, 12}
	fmt.Println(largestNumber(nums))

	nums = []int{3, 30, 34, 5, 9}
	fmt.Println(largestNumber(nums))

}
