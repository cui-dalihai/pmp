package main

import "fmt"

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	mem := make([][]int, m)
	for l := 0; l < m; l++ {
		mem[l] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			var p1, p2, p int
			if j-1 >= 0 {
				p1 = mem[i][j-1] + grid[i][j]
			} else {
				p1 = 0
			}

			if i-1 >= 0 {
				p2 = mem[i-1][j] + grid[i][j]
			} else {
				p2 = 0
			}

			if p1 == 0 && p2 == 0 {
				p = grid[i][j]
			} else if p1 == 0 {
				p = p2
			} else if p2 == 0 {
				p = p1
			} else {
				if p1 < p2 {
					p = p1
				} else {
					p = p2
				}
			}
			mem[i][j] = p
		}
	}
	return mem[m-1][n-1]
}

func main() {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	grid = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(minPathSum(grid))
}
