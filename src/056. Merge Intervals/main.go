package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	length := len(intervals)
	if length < 2 {
		return intervals
	}

	//easy use instead of sort interface
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	res := make([]Interval, 0)
	temp := intervals[0]
	for i := 1; i < length; i++ {
		//has overflow
		if temp.End >= intervals[i].Start {
			if temp.End < intervals[i].End {
				temp.End = intervals[i].End
			}
		} else {
			res = append(res, temp)
			temp = intervals[i]
		}
	}
	res = append(res, temp)
	return res
}

type IntervalSorter []Interval

func merge1(intervals []Interval) []Interval {
	length := len(intervals)
	if length < 2 {
		return intervals
	}

	sort.Sort(IntervalSorter(intervals))

	res := make([]Interval, 0)
	temp := intervals[0]
	for i := 1; i < length; i++ {
		if temp.End >= intervals[i].Start {
			//has overflow
			if temp.End >= intervals[i].End {
				continue
			} else {
				temp.End = intervals[i].End
			}
		} else {
			res = append(res, temp)
			temp = intervals[i]
		}
	}
	res = append(res, temp)
	return res
}

func (is IntervalSorter) Len() int {
	return len(is)
}

func (is IntervalSorter) Less(i, j int) bool {
	diff := is[i].Start - is[j].Start
	if diff < 0 {
		return true
	} else if diff > 0 {
		return false
	}
	diff = is[i].End - is[j].End
	if diff < 0 {
		return true
	}
	return false
}

func (is IntervalSorter) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

func main() {
	intevers := make([]Interval, 4)
	intevers[0] = Interval{1, 3}
	intevers[1] = Interval{2, 6}
	intevers[2] = Interval{8, 10}
	intevers[3] = Interval{15, 18}

	res := merge(intevers)
	fmt.Println(res)
}
