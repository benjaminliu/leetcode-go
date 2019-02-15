package main

import (
	"strconv"
	"strings"
)

func getHint(secret string, guess string) string {
	len1 := len(secret)
	countBull := 0
	maps := make(map[byte]int)
	mapIdx := make(map[int]bool)
	for i := 0; i < len1; i++ {
		if secret[i] == guess[i] {
			countBull++
			mapIdx[i] = true
		} else {
			if value, ok := maps[secret[i]]; ok {
				maps[secret[i]] = value + 1
			} else {
				maps[secret[i]] = 1
			}
		}
	}

	countCow := 0
	for i := 0; i < len1; i++ {
		if _, ok := mapIdx[i]; ok {
			continue
		}

		if value, ok := maps[guess[i]]; ok {
			countCow ++
			temp := value - 1
			if temp == 0 {
				delete(maps, guess[i])
			} else {
				maps[guess[i]] = value - 1
			}
		}
	}

	res := strings.Builder{}
	res.WriteString(strconv.Itoa(countBull))
	res.WriteByte('A')
	res.WriteString(strconv.Itoa(countCow))
	res.WriteByte('B')
	return res.String()
}

func getHint1(secret string, guess string) string {
	bull, cow := 0, 0
	count := make([]int, 10)
	len1 := len(secret)
	for i := 0; i < len1; i++ {
		s := secret[i] - '0'
		g := guess[i] - '0'
		if s == g {
			bull++
		} else {
			if count[s] < 0 {
				cow++
			}

			if count[g] > 0 {
				cow++
			}

			count[s]++
			count[g]--
		}
	}
	return strconv.Itoa(bull) + "A" + strconv.Itoa(cow) + "B"
}

func main() {

}
