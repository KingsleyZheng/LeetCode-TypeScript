package main

import "fmt"

func main() {
	b := '9'
	fmt.Println(int(b) - 48)
}

// 中等
// 36. 有效的数独
// 判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。

// 官方解法：一次迭代
// 利用哈希表
func isValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]int, 9)
	columns := make([]map[byte]int, 9)
	boxes := make([]map[byte]int, 9)
	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]int)
		columns[i] = make(map[byte]int)
		boxes[i] = make(map[byte]int)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			num := board[i][j]
			if num != '.' {
				// 关键是这里，枚举子数独
				boxIndex := (i/3)*3 + j/3
				rows[i][num]++
				columns[j][num]++
				boxes[boxIndex][num]++
				if rows[i][num] > 1 || columns[j][num] > 1 || boxes[boxIndex][num] > 1 {
					return false
				}
			}
		}
	}

	return true
}
