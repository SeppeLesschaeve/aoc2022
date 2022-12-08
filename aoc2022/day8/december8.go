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
	firstPartAmount := calculateVisible(treeHeights)
	secondPartAmount := calculateHighestScenic(treeHeights)
	fmt.Println(firstPartAmount)
	fmt.Println(secondPartAmount)
}

func calculateVisible(heights [][]int) int {
	sum := 0
	for i, row := range heights {
		for j, height := range row {
			if isVisibleRow(height, i, j, heights) || isVisibleColumn(height, i, j, heights) {
				sum += 1
			}
		}
	}
	return sum
}

func calculateHighestScenic(heights [][]int) int {
	max := 0
	for i, row := range heights {
		for j, height := range row {
			_, left := leftViewable(height, i, j, heights)
			_, right := rightViewable(height, i, j, heights)
			_, top := topViewable(height, i, j, heights)
			_, bottom := bottomViewable(height, i, j, heights)
			scenic := left * right * top * bottom
			if scenic > max {
				max = scenic
			}
		}
	}
	return max
}

func isVisibleRow(height int, i int, j int, heights [][]int) bool {
	left, _ := leftViewable(height, i, j, heights)
	right, _ := rightViewable(height, i, j, heights)
	return left || right
}

func isVisibleColumn(height int, i int, j int, heights [][]int) bool {
	top, _ := topViewable(height, i, j, heights)
	bottom, _ := bottomViewable(height, i, j, heights)
	return top || bottom
}

func leftViewable(height int, i int, j int, heights [][]int) (bool, int) {
	leftViewable := 0
	isVisibleLeft := true
	for iRow, row := range heights {
		if i == iRow {
			for jCol := j-1; jCol >= 0; jCol-- {
				if row[jCol] < height {
					leftViewable += 1
				} else {
					leftViewable += 1
					isVisibleLeft = false
					break
				}
			}
			break
		}
	}
	return isVisibleLeft, leftViewable
}

func rightViewable(height int, i int, j int, heights [][]int) (bool, int) {
	rightViewable := 0
	isVisibleRight := true
	for iRow, row := range heights {
		if i == iRow {
			for jCol := j+1; jCol <= len(row) - 1; jCol++ {
				if row[jCol] < height {
					rightViewable += 1
				} else {
					rightViewable += 1
					isVisibleRight = false
					break
				}
			}
			break
		}
	}
	return isVisibleRight, rightViewable
}

func topViewable(height int, i int, j int, heights [][]int) (bool, int) {
	topViewable := 0
	isVisibleTop := true
	for iRow := i-1; iRow >= 0; iRow-- {
		if heights[iRow][j] < height {
			topViewable += 1
		} else {
			topViewable += 1
			isVisibleTop = false
			break
		}
	}
	return isVisibleTop, topViewable
}

func bottomViewable(height int, i int, j int, heights [][]int) (bool, int) {
	bottomViewable := 0
	isBottomVisible := true
	for iRow := i+1; iRow <= len(heights) - 1; iRow++ {
		if heights[iRow][j] < height {
			bottomViewable += 1
		} else {
			bottomViewable += 1
			isBottomVisible = false
			break
		}
	}
	return isBottomVisible, bottomViewable
}
