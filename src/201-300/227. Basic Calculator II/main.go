package main

import (
	"container/list"
	"fmt"
)

func calculate(s string) int {
	len1 := len(s)
	if len1 == 0 {
		return 0
	}
	oprandStack := list.New()
	operatorStack := make([]byte, 0)

	i := 0
	for s[i] < '0' || s[i] > '9' {
		i++
	}

	num := int(s[i] - '0')
	i++
	lastOperator := byte(0)
	for ; i < len1; i++ {
		if s[i] == '+' {
			calc(oprandStack, &lastOperator, &num)
			operatorStack = append(operatorStack, '+')
		} else if s[i] == '-' {
			calc(oprandStack, &lastOperator, &num)
			operatorStack = append(operatorStack, '-')
		} else if s[i] == '*' {
			calc(oprandStack, &lastOperator, &num)
			lastOperator = '*'
		} else if s[i] == '/' {
			calc(oprandStack, &lastOperator, &num)
			lastOperator = '/'
		} else if s[i] >= '0' && s[i] <= '9' {
			num = num*10 + int(s[i]-'0')
		}
	}
	calc(oprandStack, &lastOperator, &num)

	item := oprandStack.Front()
	oprandStack.Remove(item)
	left := item.Value.(int)
	idx := 0
	for oprandStack.Len() > 0 {
		rightItem := oprandStack.Front()
		oprandStack.Remove(rightItem)
		right := rightItem.Value.(int)

		if operatorStack[idx] == '+' {
			left += right
		} else {
			left -= right
		}
		idx++
	}
	return left
}

func calc(oprandStack *list.List, lastOperator *byte, num *int) {
	temp := *num
	*num = 0

	if *lastOperator == '*' {
		item := oprandStack.Back()
		oprandStack.Remove(item)

		left := item.Value.(int)
		temp = left * temp

		*lastOperator = 0
	} else if *lastOperator == '/' {
		item := oprandStack.Back()
		oprandStack.Remove(item)

		left := item.Value.(int)
		temp = left / temp
		*lastOperator = 0
	}

	oprandStack.PushBack(temp)
}

func main() {
	fmt.Println(calculate("3+2*2"))
	fmt.Println(calculate(" 3/2 "))
	fmt.Println(calculate("3+5/2"))
}
