package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	length := len(strs)
	if length == 0 {
		return res
	}
	maps := make(map[string][]string)
	arrays := make([]byte, 26)
	for _, str := range strs {
		key := calcKey(str, arrays)
		list := maps[key]
		if list == nil {
			list = make([]string, 0)
		}
		list = append(list, str)
		maps[key] = list
	}
	for _, list := range maps {
		res = append(res, list)
	}
	return res
}

func calcKey(str string, maps []byte) string {
	for i := 0; i < len(str); i++ {
		c := str[i]
		maps[c-'a']++
	}
	res := make([]byte, 0)
	for i := 0; i < 26; i++ {
		for maps[i] > 0 {
			maps[i]--
			res = append(res, byte(i)+'a')
		}
	}
	return string(res)
}

// 26 alphas map to 26 primes,
// each prime means an alpha,
// product of primes is unique
func groupAnagrams1(strs []string) [][]string {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103}

	maps := make(map[int][]string)
	group := make([]int, 0)
	for _, v := range strs {
		product := 1
		for i := 0; i < len(v); i++ {
			product *= primes[int(v[i])-'a']
		}

		if maps[product] == nil {
			maps[product] = []string{v}
			group = append(group, product)
		} else {
			maps[product] = append(maps[product], v)
		}
	}

	res := make([][]string, len(group))
	for i, sa := range group {
		res[i] = maps[sa]
	}

	return res
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
