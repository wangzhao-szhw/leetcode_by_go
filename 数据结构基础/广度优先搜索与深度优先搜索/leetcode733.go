package main

import "fmt"

//深度优先搜索
var (
	dx = []int{0, 0, -1, 1}
	dy = []int{-1, 1, 0, 0}
)

//广度优先搜索
func floodFill2(image [][]int, sr int, sc int, newColor int) [][]int {
	n := len(image)
	m := len(image[0])
	var que [][]int
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
	var x [][]int

	x = [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}

	sr := 1
	sc := 1
	newColor := 2
	res2 := floodFill2(x, sr, sc, newColor)

	fmt.Println(res2)

}
