package main

func intersect(nums1 []int, nums2 []int) []int {
	n1, n2 := len(nums1), len(nums2)
	if n1 == 0 || n2 == 0 {
		return nil
	}

	maps := make(map[int]int)

	for _, value := range nums1 {
		if t, ok := maps[value]; ok {
			maps[value] = t + 1
		} else {
			maps[value] = 1
		}
	}

	res := make([]int, 0)
	for _, value := range nums2 {
		if value1, ok := maps[value]; ok {
			res = append(res, value)

			if value1 > 1 {
				maps[value] = value1 - 1
			} else {
				delete(maps, value)
			}
		}
	}
	return res
}

func main() {

}
