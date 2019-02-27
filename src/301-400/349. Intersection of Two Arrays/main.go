package main

func intersection(nums1 []int, nums2 []int) []int {
	n1, n2 := len(nums1), len(nums2)
	if n1 == 0 || n2 == 0 {
		return nil
	}

	maps := make(map[int]int)

	for _, value := range nums1 {
		maps[value] = 1
	}

	res := make([]int, 0)
	for _, value := range nums2 {
		if value1, ok := maps[value]; ok {
			if value1 == 1 {
				res = append(res, value)
			}
			maps[value] = 2
		}
	}
	return res
}

func main() {

}
