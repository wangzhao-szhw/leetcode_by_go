package main

import "fmt"

//深度优先搜索
var (
	dx = []int{0, 0, -1, 1}
	dy = []int{-1, 1, 0, 0}
)

func dfs(image [][]int, x int, y int, color int, new int) {
	image[x][y] = new
	for i := 0; i < 4; i++ {
		xm, ym := x+dx[i], y+dy[i]
		if xm >= 0 && xm < len(image) && ym >= 0 && ym < len(image[0]) && image[xm][ym] == color {
			dfs(image, xm, ym, color, new)
		}
	}
}
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	curColor := image[sr][sc]
	if image[sr][sc] != newColor {
		dfs(image, sr, sc, curColor, newColor)
	}
	return image
}

//广度优先搜索
func floodFill2(image [][]int, sr int, sc int, newColor int) [][]int {
	n := len(image)
	m := len(image[0])
	que := [][]int{}
	curColor := image[sr][sc]
	if curColor == newColor {
		return image
	}
	que = append(que, []int{sr, sc})
	image[sr][sc] = newColor
	for i := 0; i < len(que); i++ {
		tmp := que[i]
		for j := 0; j < 4; j++ {
			xm, ym := tmp[0]+dx[j], tmp[1]+dy[j]
			if xm >= 0 && ym >= 0 && xm < n && ym < m && image[xm][ym] == curColor {
				que = append(que, []int{xm, ym})
				image[xm][ym] = newColor
			}
		}
	}
	return image

}
func main() {
	x := [][]int{}
	x[0] = []int{1, 1, 1}
	x[1] = []int{1, 1, 0}
	x[2] = []int{1, 0, 1}
	sr := 1
	sc := 1
	newColor := 2
	res1 := floodFill(x, sr, sc, newColor)
	res2 := floodFill2(x, sr, sc, newColor)
	fmt.Println(res1)
	fmt.Println(res2)

}
