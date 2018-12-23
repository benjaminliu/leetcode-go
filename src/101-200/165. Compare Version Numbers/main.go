package main

import "fmt"

func compareVersion(version1 string, version2 string) int {
	len1 := len(version1)
	len2 := len(version2)
	if len1 == 0 {
		if len2 == 0 {
			return 0
		}
		return -1
	}

	if len2 == 0 {
		return 1
	}
	idx1 := 0
	idx2 := 0

	for idx1 < len1 && idx2 < len2 {
		temp1 := parseOneVersion(version1, &idx1, len1)

		temp2 := parseOneVersion(version2, &idx2, len2)

		if temp1 > temp2 {
			return 1
		} else if temp1 < temp2 {
			return -1
		}
		temp1 = 0
		temp2 = 0
	}

	for idx1 < len1 {
		temp := parseOneVersion(version1, &idx1, len1)
		if temp > 0 {
			return 1
		}
	}
	for idx2 < len2 {
		temp := parseOneVersion(version2, &idx2, len2)
		if temp > 0 {
			return -1
		}
	}
	return 0
}

func parseOneVersion(version string, idx *int, len1 int) int {
	temp := 0
	for *idx < len1 {
		char1 := version[*idx]
		*idx++
		if char1 == '.' {
			break
		} else {
			temp = temp*10 + int(char1-'0')
		}
	}
	return temp
}

func main() {
	fmt.Println(compareVersion("1.0.1", "1.0.1"))
	fmt.Println(compareVersion("7.5.2.4", "7.5.3"))
	fmt.Println(compareVersion("0.0.1", "1.0.1"))
	fmt.Println(compareVersion("1.0", "1"))
}
