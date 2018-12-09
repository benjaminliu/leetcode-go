package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	len1 := len(gas)
	if len1 == 0 {
		return 0
	}
	sum := 0
	diff := make([]int, len1)
	for i := 0; i < len1; i++ {
		diff [i] = gas[i] - cost[i]
		sum += diff[i]
	}
	if sum < 0 {
		return -1
	}

	found := false
	for i := 0; i < len1; i++ {
		if diff[i] < 0 {
			continue
		}
		found = true
		sum = 0
		for j := i; j < len1+i; j++ {
			k := j
			if j >= len1 {
				k = j - len1
			}
			sum += diff[k]
			if sum < 0 {
				found = false
				break
			}
		}
		if found == true {
			return i
		}
	}
	return -1
}

func canCompleteCircuit1(gas []int, cost []int) int {
	len1 := len(gas)
	if len1 == 0 {
		return 0
	}
	sum, gasSum, costSum := 0, 0, 0

	start := 0
	for i := 0; i < len1; i++ {
		gasSum += gas[i]
		costSum += cost[i]
		sum += gas[i] - cost[i]
		if sum < 0 {
			sum = 0
			start = i + 1
		}
	}

	if gasSum >= costSum {
		return start
	}
	return -1
}

func main() {
	//gas := []int{1, 2, 3, 4, 5}
	//cost := []int{3, 4, 5, 1, 2}

	//gas := []int{2,3,4}
	//cost := []int{3,4,3}

	gas := []int{2}
	cost := []int{2}

	fmt.Println(canCompleteCircuit(gas, cost))
}
