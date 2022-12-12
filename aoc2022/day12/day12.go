package day12

import (
	"fmt"
	"os"
	"strings"
)

type Position struct {
	x int
	y int
}

func Day12() {
	content, _ := os.ReadFile("day12.txt")
	day12Content := string(content)
	heightMap, start, end := getHeightMap(day12Content)

	dist := map[Position]int{end: 0}
	queue := []Position{end}
	var shortest *Position

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if heightMap[curr.x][curr.y] == 97 && shortest == nil {
			shortest = &curr
		}

		for i := curr.x - 1; i <= curr.x+1; i++ {
			for j := curr.y - 1; j <= curr.y+1; j++ {
				next := Position{i, j}
				if getHamming(curr, next) == 1 && next.x >= 0 && next.x < len(heightMap) && next.y >= 0 && next.y < len(heightMap[0]) {
					_, seen := dist[next]
					if !seen && heightMap[curr.x][curr.y] <= heightMap[next.x][next.y]+1 {
						dist[next] = dist[curr] + 1
						queue = append(queue, next)
					}
				}
			}
		}
	}

	fmt.Println(dist[start])
	fmt.Println(dist[*shortest])
}

func abs(i int) int {
	if i >= 0 {
		return i
	} else {
		return -i
	}
}

func getHamming(start Position, end Position) int {
	return abs(end.x-start.x) + abs(end.y-start.y)
}

func getHeightMap(content string) ([][]int32, Position, Position) {
	lines := strings.Split(content, "\n")
	var start Position
	var end Position
	var heightMap [][]int32
	for iRow, row := range lines {
		var heightRow []int32
		for iCol, col := range row {
			if col == 83 {
				start = Position{iRow, iCol}
				heightRow = append(heightRow, 97)
			} else if col == 69 {
				end = Position{iRow, iCol}
				heightRow = append(heightRow, 122)
			} else {
				heightRow = append(heightRow, col)
			}
		}
		heightMap = append(heightMap, heightRow)
	}
	return heightMap, start, end
}
