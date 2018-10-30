package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	if len(board) != 9 || len(board[0]) != 9 {
		return false
	}
	rows := make([]map[byte]int, 9)
	rows[0] = make(map[byte]int)
	rows[1] = make(map[byte]int)
	rows[2] = make(map[byte]int)
	rows[3] = make(map[byte]int)
	rows[4] = make(map[byte]int)
	rows[5] = make(map[byte]int)
	rows[6] = make(map[byte]int)
	rows[7] = make(map[byte]int)
	rows[8] = make(map[byte]int)

	cols := make([]map[byte]byte, 9)
	cols[0] = make(map[byte]byte)
	cols[1] = make(map[byte]byte)
	cols[2] = make(map[byte]byte)
	cols[3] = make(map[byte]byte)
	cols[4] = make(map[byte]byte)
	cols[5] = make(map[byte]byte)
	cols[6] = make(map[byte]byte)
	cols[7] = make(map[byte]byte)
	cols[8] = make(map[byte]byte)

	squre := make([][]map[byte]byte, 3, 3)
	squre[0] = make([]map[byte]byte, 3)
	squre[1] = make([]map[byte]byte, 3)
	squre[2] = make([]map[byte]byte, 3)
	squre[0][0] = make(map[byte]byte)
	squre[0][1] = make(map[byte]byte)
	squre[0][2] = make(map[byte]byte)
	squre[1][0] = make(map[byte]byte)
	squre[1][1] = make(map[byte]byte)
	squre[1][2] = make(map[byte]byte)
	squre[2][0] = make(map[byte]byte)
	squre[2][1] = make(map[byte]byte)
	squre[2][2] = make(map[byte]byte)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			value := board[i][j]
			if value=='.' {
				continue
			}
			if _, ok := rows[i][value]; ok {
				return false
			}
			rows[i][value] = 1
			if _, ok := cols[j][value]; ok {
				return false
			}
			cols[j][value] = 1

			var x byte

			if i < 3 {
				x = 0
			} else if i < 6 {
				x = 1
			} else {
				x = 2
			}

			var y byte

			if j < 3 {
				y = 0
			} else if j < 6 {
				y = 1
			} else {
				y = 2
			}

			if _, ok := squre[x][y][value]; ok {
				return false
			}
			squre[x][y][value] = 1
		}
	}
	return true
}

func main() {
	var data = [][]byte{
		[]byte("53..7...."),
		[]byte("6..195..."),
		[]byte(".98....6."),
		[]byte("8...6...3"),
		[]byte("4..8.3..1"),
		[]byte("7...2...6"),
		[]byte(".6....28."),
		[]byte("...419..5"),
		[]byte("....8..79"),
	}
	fmt.Println(isValidSudoku(data))
}


