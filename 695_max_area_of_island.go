package main

func maxAreaOfIsland(grid [][]int) int {
	h, w := len(grid), len(grid[0])
	vis := make([][]bool, h)
	for i := 0; i < h; i++ {
		vis[i] = make([]bool, w)
	}
	type point struct {
		y, x int
	}
	dy := []int{0, 1, 0, -1}
	dx := []int{1, 0, -1, 0}
	ans := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if vis[i][j] || grid[i][j] == 0 {
				continue
			}
			que := []point{}
			que = append(que, point{y: i, x: j})
			vis[i][j] = true
			sum := 0
			for len(que) > 0 {
				now := que[0]
				que = que[1:]
				sum++
				for k := 0; k < 4; k++ {
					ny, nx := now.y+dy[k], now.x+dx[k]
					if ny < 0 || ny >= h || nx < 0 || nx >= w || vis[ny][nx] || grid[ny][nx] == 0 {
						continue
					}
					que = append(que, point{y: ny, x: nx})
					vis[ny][nx] = true
				}
			}
			if sum > ans {
				ans = sum
			}
		}
	}
	return ans
}

// You are given an m x n binary matrix grid. An island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.) You may assume all four edges of the grid are surrounded by water.

// The area of an island is the number of cells with a value 1 in the island.

// Return the maximum area of an island in grid. If there is no island, return 0.

// Input: grid = [[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,0,1,0,0,0,0,0,0,0,0],[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,0,0,0,0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]]
// Output: 6
// Explanation: The answer is not 11, because the island must be connected 4-directionally.
// Example 2:

// Input: grid = [[0,0,0,0,0,0,0,0]]
// Output: 0

// Constraints:

// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 50
// grid[i][j] is either 0 or 1.
