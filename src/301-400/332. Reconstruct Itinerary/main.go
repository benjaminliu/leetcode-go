package main

import (
	"fmt"
	"sort"
)

var graph map[string][]string
var itinerary []string

//graph, need more care
func findItinerary(tickets [][]string) []string {
	graph = map[string][]string{}
	itinerary = []string{}

	n := len(tickets)
	for i := 0; i < n; i++ {
		if next, ok := graph[tickets[i][0]]; ok {
			next = append(next, tickets[i][1])
			graph[tickets[i][0]] = next
		} else {
			graph[tickets[i][0]] = []string{tickets[i][1]}
		}
	}

	for _, value := range graph {
		sort.Strings(value)
	}
	find("JFK")
	return reverse(itinerary)
}
func find(cur string) {
	for {
		if next, ok := graph[cur]; ok && len(next) > 0 {
			temp := next[0]
			graph[cur] = next[1:]
			find(temp)
		} else {
			break
		}
	}

	itinerary = append(itinerary, cur)
}
func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func main() {
	fmt.Println(findItinerary([][]string{
		{"JFK", "KUL"},
		{"JFK", "NRT"},
		{"NRT", "JFK"},
	}))
}
