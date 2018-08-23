package main

import (
		"fmt"
)

/*
给定一个包含了一些 0 和 1的非空二维数组 grid , 一个 岛屿 是由四个方向 (水平或垂直) 的 1 (代表土地) 构成的组合。你可以假设二维矩阵的四个边缘都被水包围着。
找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为0。)
示例 1:
[[0,0,1,0,0,0,0,1,0,0,0,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,1,1,0,1,0,0,0,0,0,0,0,0],
 [0,1,0,0,1,1,0,0,1,0,1,0,0],
 [0,1,0,0,1,1,0,0,1,1,1,0,0],
 [0,0,0,0,0,0,0,0,0,0,1,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,0,0,0,0,0,0,1,1,0,0,0,0]]
对于上面这个给定矩阵应返回 6。注意答案不应该是11，因为岛屿只能包含水平或垂直的四个方向的‘1’。
示例 2:
[[0,0,0,0,0,0,0,0]]
对于上面这个给定的矩阵, 返回 0。
注意: 给定的矩阵grid 的长度和宽度都不超过 50。
 */


func maxAreaOfIsland(grid [][]int) int {
	mx:=0
	n:=len(grid)
	m:=len(grid[0])
	for i:=0;i<n;i++ {
		for j:=0;j<m;j++{
			if grid[i][j]==1 {
				num:=dfs(grid,i,j)
				mx=Max(mx,num)
			}
		}
	}
	return int(mx)
}

func Max(a,b int) int{
	if a>b {
		return a
	}else {
		return b
	}

}
func dfs(grid [][]int,x0,y0 int) int {
	s:=1
	n:=len(grid)
	m:=len(grid[0])

	grid[x0][y0]=0

	dire:=[][]int{{0,1},{0,-1},{1,0},{-1,0}}
	for i:=0;i<4 ;i++  {
		x:=x0+dire[i][0]
		y:=y0+dire[i][1]
		if x>=0 && y>=0{
			if x<n && y<m {
				if grid[x][y]==1 {
					s=s+dfs(grid,x,y)
					fmt.Println(s)
				}
			}
		}
	}
	return s
}

func main()  {

	grid:=[][]int{{1,1}}

	a:=maxAreaOfIsland(grid)
	fmt.Println(a)
}