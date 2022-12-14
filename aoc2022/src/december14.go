package src

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Day14() {
	content, _ := os.ReadFile("input/day14.txt")
	day14Content := string(content)
	blockPositions1, maxDepth1 := getBlockPositions(day14Content)
	blockPositions2, maxDepth2 := getBlockPositions(day14Content)
	left, right := getLeftAndRight(blockPositions2, maxDepth2)
	fmt.Println(simulateSand1(blockPositions1, maxDepth1))
	fmt.Println(simulateSand2(blockPositions2, maxDepth2, left, right))
}

func getLeftAndRight(positions map[Position]bool, depth int) (int, int) {
	left := math.MaxInt
	right := math.MinInt
	for position, b := range positions {
		if b {
			if position.x == depth {
				if position.y < left {
					left = position.y
				}
				if position.y > right {
					right = position.y
				}
			}
		}
	}
	return left, right
}

func getBlockPositions(content string) (map[Position]bool, int) {
	blockPositions := make(map[Position]bool)
	maxDepth := 0
	paths := strings.Split(content, "\n")
	for _, path := range paths {
		var pathPositions []Position
		coords := strings.Split(path, " -> ")
		for _, coord := range coords {
			pos := strings.Split(coord, ",")
			x, _ := strconv.Atoi(pos[1])
			y, _ := strconv.Atoi(pos[0])
			pathPositions = append(pathPositions, Position{x, y})
		}
		blockPositions, maxDepth = getBlockedPathPositions(pathPositions, blockPositions, maxDepth)
	}
	return blockPositions, maxDepth
}

func getBlockedPathPositions(positions []Position, blockPositions map[Position]bool, maxDepth int) (map[Position]bool, int) {
	for i := 0; i < len(positions)-1; i++ {
		pos1 := positions[i]
		pos2 := positions[i+1]
		var from, to Position
		if pos1.x == pos2.x {
			from = Position{pos1.x, min(pos1.y, pos2.y)}
			to = Position{pos1.x, max(pos1.y, pos2.y)}
			for i := from.y; i <= to.y; i++ {
				blockPositions[Position{pos1.x, i}] = true
			}
			if pos1.x > maxDepth {
				maxDepth = pos1.x
			}
		} else if pos1.y == pos2.y {
			from = Position{min(pos1.x, pos2.x), pos1.y}
			to = Position{max(pos1.x, pos2.x), pos1.y}
			for i := from.x; i <= to.x; i++ {
				blockPositions[Position{i, pos1.y}] = true
				if i > maxDepth {
					maxDepth = i
				}
			}
		}
	}
	return blockPositions, maxDepth
}

func simulateSand1(positions map[Position]bool, depth int) int {
	sand := 0
	var sandPos Position
	for sandPos.x != depth+1 {
		sandPos = Position{0, 500}
		for !positions[sandPos] {
			bottom := Position{sandPos.x + 1, sandPos.y}
			bottomLeft := Position{sandPos.x + 1, sandPos.y - 1}
			bottomRight := Position{sandPos.x + 1, sandPos.y + 1}
			if !positions[bottom] {
				sandPos = bottom
			} else if !positions[bottomLeft] {
				sandPos = bottomLeft
			} else if !positions[bottomRight] {
				sandPos = bottomRight
			} else {
				positions[sandPos] = true
			}
			if sandPos.x == depth+1 {
				break
			}
		}
		if sandPos.x < depth+1 {
			sand++
		}
	}
	return sand
}

func simulateSand2(positions map[Position]bool, depth int, left int, right int) int {
	sand := 0
	for i := left - 1000; i <= right+1000; i++ {
		positions[Position{depth + 2, i}] = true
	}
	var sandPos Position
	for !positions[Position{0, 500}] {
		sandPos = Position{0, 500}
		for !positions[sandPos] {
			bottom := Position{sandPos.x + 1, sandPos.y}
			bottomLeft := Position{sandPos.x + 1, sandPos.y - 1}
			bottomRight := Position{sandPos.x + 1, sandPos.y + 1}
			if !positions[bottom] {
				sandPos = bottom
			} else if !positions[bottomLeft] {
				sandPos = bottomLeft
			} else if !positions[bottomRight] {
				sandPos = bottomRight
			} else {
				positions[sandPos] = true
			}
		}
		sand++
	}
	return sand
}
