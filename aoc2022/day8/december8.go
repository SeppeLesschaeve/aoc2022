package day8

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day8() {
	content, _ := os.ReadFile("day8.txt")
	day8Content := string(content)
	rows := strings.Split(day8Content, "\n")
	treeHeights := make([][]int, len(rows))
	for i, row := range rows {
		for _, tree := range row {
			treeHeight, _ := strconv.Atoi(string(tree))
			treeHeights[i] = append(treeHeights[i], treeHeight)
		}
	}
	firstPartAmount, secondPartAmount := calculateVisibleAndHighestScenic(treeHeights)
	fmt.Println(firstPartAmount)
	fmt.Println(secondPartAmount)
}

func calculateVisibleAndHighestScenic(heights [][]int) (int, int) {
	sum := 0
	max := 0
	for i, row := range heights {
		for j, height := range row {
			var leftDir []int
			for jCol := j - 1; jCol >= 0; jCol-- {
				leftDir = append(leftDir, heights[i][jCol])
			}
			var rightDir []int
			for jCol := j + 1; jCol <= len(row)-1; jCol++ {
				rightDir = append(rightDir, heights[i][jCol])
			}
			var topDir []int
			for iRow := i - 1; iRow >= 0; iRow-- {
				topDir = append(topDir, heights[iRow][j])
			}
			var bottomDir []int
			for iRow := i + 1; iRow <= len(heights)-1; iRow++ {
				bottomDir = append(bottomDir, heights[iRow][j])
			}
			isVisibleLeft, leftViewable := viewableInDirection(height, leftDir)
			isVisibleRight, rightViewable := viewableInDirection(height, rightDir)
			isVisibleTop, topViewable := viewableInDirection(height, topDir)
			isVisibleBottom, bottomViewable := viewableInDirection(height, bottomDir)
			if isVisibleLeft || isVisibleRight || isVisibleTop || isVisibleBottom {
				sum += 1
			}
			scenic := leftViewable * rightViewable * topViewable * bottomViewable
			if scenic > max {
				max = scenic
			}
		}
	}
	return sum, max
}

func viewableInDirection(height int, heights []int) (bool, int) {
	viewable := 0
	isVisible := true
	for _, h := range heights {
		viewable += 1
		if h >= height {
			isVisible = false
			break
		}
	}
	return isVisible, viewable
}
