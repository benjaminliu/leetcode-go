package main

import "fmt"

var maps = map[string]int{
	"I":  1,
	"IV": 4,
	"V":  5,
	"IX": 9,
	"X":  10,
	"XL": 40,
	"L":  50,
	"XC": 90,
	"C":  100,
	"CD": 400,
	"D":  500,
	"CM": 900,
	"M":  1000,
}

func romanToInt(s string) int {
	length := len(s)
	max := 0
	sum := 0
	for i := length - 1; i >= 0; i-- {
		key := s[i : i+1]
		value := maps[key]

		if value > max {
			max = value
		}

		if value < max {
			sum -= value
		} else {
			sum += value
		}
	}
	return sum
}
func romanToInt1(s string) int {
	length := len(s)
	idx := 0
	res := 0
	for idx < length {
		if idx+1 < length {
			key := s[idx : idx+2]
			if value, ok := maps[key]; ok {
				res += value
				idx += 2
				continue
			}
		}
		key := s[idx : idx+1]
		res += maps[key]
		idx ++
	}
	return res
}

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("MCMXCIV"))
}
