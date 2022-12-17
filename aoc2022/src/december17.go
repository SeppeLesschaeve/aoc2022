package src

import (
	"fmt"
	"os"
	"strconv"
)

func Day17() {
	content, _ := os.ReadFile("input/day17.txt")
	pattern := string(content)
	shapes := initializeShapes()
	landing1 := initializeLanding()
	height1 := getHeight(landing1)
	for i, jetIndex := 0, 0; i < 2022; i++ {
		shape := getInitialShape(shapes[i%5], height1+4)
		jetIndex = chamberAfterShape(pattern, jetIndex, shape, landing1)
		height1 = getHeight(landing1)
	}
	fmt.Println(height1)
	var hash string
	var hashes []string
	i := 0
	amountsLanded := 0
	landing2 := initializeLanding()
	height2 := getHeight(landing2)
	shape := getInitialShape(shapes[0], height2+4)
	hashVals := make(map[string][]int)
	for true {
		hash = getHash(landing2, height2, i)
		if contains(hashes, hash) {
			break
		}
		hashes = append(hashes, hash)
		hashVals[hash] = []int{amountsLanded, height2}
		if canBePushed(shape, rune(pattern[i]), landing2) {
			shape = updateShapeAfterPush(shape, rune(pattern[i]))
		}
		i = (i + 1) % len(pattern)
		if canFall(shape, landing2) {
			shape = updateShapeAfterFall(shape)
		} else {
			updateLanding(shape, landing2)
			height2 = getHeight(landing2)
			amountsLanded += 1
			shape = getInitialShape(shapes[amountsLanded%5], height2+4)
		}
	}
	rocksPerCycle := amountsLanded - hashVals[hash][0]
	heightPerCycle := height2 - hashVals[hash][1]
	rocks := 1000000000000
	amountOfCycles := (rocks - amountsLanded) / rocksPerCycle
	rocksLeft := rocks - amountsLanded - rocksPerCycle*amountOfCycles

	amountsLanded = 0
	for amountsLanded < rocksLeft {
		if canBePushed(shape, rune(pattern[i]), landing2) {
			shape = updateShapeAfterPush(shape, rune(pattern[i]))
		}
		i = (i + 1) % len(pattern)
		if canFall(shape, landing2) {
			shape = updateShapeAfterFall(shape)
		} else {
			updateLanding(shape, landing2)
			height2 = getHeight(landing2)
			amountsLanded += 1
			shape = getInitialShape(shapes[amountsLanded%5], height2+4)
		}
	}
	fmt.Println(height2 + heightPerCycle*amountOfCycles)
}

func getHash(landing map[Position]bool, height int, i int) string {
	var diffHeight []int
	for col := 0; col < 7; col++ {
		d := height
		var cols []int
		for position, _ := range landing {
			if position.y == col {
				cols = append(cols, d-position.x)
			}
		}
		minimumOfCol := cols[0]
		for _, colDiff := range cols {
			minimumOfCol = min(minimumOfCol, colDiff)
		}
		diffHeight = append(diffHeight, minimumOfCol)
	}
	hash := strconv.Itoa(i)
	hash += strconv.Itoa(i % 5)
	for _, diff := range diffHeight {
		hash += strconv.Itoa(diff)
	}
	return hash
}

func getHeight(landing map[Position]bool) int {
	height := 0
	for position, _ := range landing {
		if position.x > height {
			height = position.x
		}
	}
	return height
}

func chamberAfterShape(jet string, index int, shape []Position, landing map[Position]bool) int {
	newIndex := index
	for true {
		if canBePushed(shape, rune(jet[newIndex]), landing) {
			shape = updateShapeAfterPush(shape, rune(jet[newIndex]))
		}
		newIndex = (newIndex + 1) % len(jet)
		if canFall(shape, landing) {
			shape = updateShapeAfterFall(shape)
		} else {
			break
		}
	}
	updateLanding(shape, landing)
	return newIndex
}

func updateLanding(shape []Position, landing map[Position]bool) {
	for _, position := range shape {
		landing[position] = true
	}
}

func canBePushed(shape []Position, dir rune, landing map[Position]bool) bool {
	for _, position := range shape {
		if (dir == '<' && position.y == 0) || (dir == '>' && position.y == 6) {
			return false
		}
		if (dir == '<' && landing[Position{position.x, position.y - 1}]) || (dir == '>' && landing[Position{position.x, position.y + 1}]) {
			return false
		}
	}
	return true
}

func updateShapeAfterPush(shape []Position, dir rune) []Position {
	var newShape []Position
	for _, position := range shape {
		if dir == '<' {
			newShape = append(newShape, Position{position.x, position.y - 1})
		} else if dir == '>' {
			newShape = append(newShape, Position{position.x, position.y + 1})
		}
	}
	return newShape
}

func canFall(shape []Position, landing map[Position]bool) bool {
	for _, position := range shape {
		fallPosition := Position{position.x - 1, position.y}
		if landing[fallPosition] {
			return false
		}
	}
	return true
}

func updateShapeAfterFall(shape []Position) []Position {
	var newShape []Position
	for _, position := range shape {
		newShape = append(newShape, Position{position.x - 1, position.y})
	}
	return newShape
}

func getInitialShape(shape []Position, height int) []Position {
	var initialShape []Position
	for _, pos := range shape {
		initialShape = append(initialShape, Position{pos.x + height, pos.y})
	}
	return initialShape
}

func initializeLanding() map[Position]bool {
	landing := make(map[Position]bool)
	for i := 0; i < 7; i++ {
		landing[Position{0, i}] = true
	}
	return landing
}

func initializeShapes() [5][]Position {
	horln := []Position{{0, 2}, {0, 3}, {0, 4}, {0, 5}}
	cross := []Position{{0, 3}, {1, 2}, {1, 3}, {1, 4}, {2, 3}}
	angle := []Position{{0, 2}, {0, 3}, {0, 4}, {1, 4}, {2, 4}}
	verln := []Position{{0, 2}, {1, 2}, {2, 2}, {3, 2}}
	block := []Position{{0, 2}, {0, 3}, {1, 2}, {1, 3}}
	return [5][]Position{horln, cross, angle, verln, block}
}
