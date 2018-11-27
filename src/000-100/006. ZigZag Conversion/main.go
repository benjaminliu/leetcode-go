package main

import "fmt"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	length := len(s)
	if numRows >= length {
		return s
	}
	//rowSize := length/numRows + 1

	res := make([]byte, length)

	board := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		board[i] = make([]byte, 0)
	}

	idx := 0
	row := 0
	down := true
	for idx < length {
		board[row] = append(board[row], s[idx])
		idx ++
		if down {
			if row == numRows-1 {
				down = false
				row--
			} else {
				row++
			}
		} else {
			if row == 0 {
				down = true
				row ++
			} else {
				row --
			}
		}
	}

	resIdx := 0
	for i := 0; i < numRows; i++ {
		temp := board[i]
		var tempIdx = len(temp)
		for j := 0; j < tempIdx; j++ {
			res[resIdx] = board[i][j]
			resIdx++
		}
	}

	return string(res)
}

func convert1(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	res := make([]byte, len(s))
	stp := numRows + numRows - 2
	cur := 0
	length := len(s)

	for r := 0; r < numRows; r++ {
		for i := 0; i < length; i += stp {
			p0 := i + r
			if p0 < length {
				res[cur] = s[p0]
				cur++

				if r != 0 && r != numRows-1 {
					p1 := i + stp - r
					if p1 < length {
						res[cur] = s[p1]
						cur++
					}
				}
			}
		}
	}

	return string(res)
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 4))
}
