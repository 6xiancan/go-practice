package main

func _getIsland(grid [][]int, i int, j int) bool {
	tag := true
	m, n := len(grid), len(grid[0])
	grid[i][j] = 1
	if i == 0 || i+1 == m || j == 0 || j+1 == n {
		tag = false
	}
	if i+1 < m && grid[i+1][j] == 0 {
		if !_getIsland(grid, i+1, j) {
			tag = false
		}
	}
	if i-1 >= 0 && grid[i-1][j] == 0 {
		if !_getIsland(grid, i-1, j) {
			tag = false
		}
	}
	if j+1 < n && grid[i][j+1] == 0 {
		if !_getIsland(grid, i, j+1) {
			tag = false
		}
	}
	if j-1 >= 0 && grid[i][j-1] == 0 {
		if !_getIsland(grid, i, j-1) {
			tag = false
		}
	}
	return tag
}

func closedIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				continue
			}
			if _getIsland(grid, i, j) {
				res++
			}
		}
	}
	return res
}
