package src

import (
	"fmt"
	"os"
	"strconv"
)

func Day17() {
	content, _ := os.ReadFile("input/day17.txt")
	day17Content := string(content)
	shapes := initializeShapes()
	part1(day17Content, shapes)
	part2(day17Content, shapes)
}

func part1(jets string, shapes [5][]Position) {
	landing := initializeLanding()
	height := 0
	for i, jetIndex := 0, 0; i < 2022; i++ {
		shape := getInitialShape(shapes[i%5], height+4)
		for true {
			if canBePushed(shape, rune(jets[jetIndex]), landing) {
				shape = updateShapeAfterPush(shape, rune(jets[jetIndex]))
			}
			jetIndex = (jetIndex + 1) % len(jets)
			if !canFall(shape, landing) {
				break
			}
			shape = updateShapeAfterFall(shape)
		}
		updateLanding(shape, landing)
		height = getHeight(height, shape)
	}
	fmt.Println(height)
}

func part2(jets string, shapes [5][]Position) {
	var hash string
	var hashes []string
	i, amountsLanded, currentShape, height := 0, 0, 0, 0
	landing := initializeLanding()
	shape := getInitialShape(shapes[0], height+4)
	hashVals := make(map[string][]int)
	for true {
		hash = getHash(landing, height, i)
		if contains(hashes, hash) {
			break
		}
		hashes = append(hashes, hash)
		hashVals[hash] = []int{amountsLanded, height}
		shape, i, height, amountsLanded, currentShape =
			move(shape, i, height, amountsLanded, currentShape, jets, shapes, landing)
	}
	rocksPerCycle := amountsLanded - hashVals[hash][0]
	heightPerCycle := height - hashVals[hash][1]
	rocks := 1000000000000
	amountOfCycles := (rocks - amountsLanded) / rocksPerCycle
	rocksLeft := rocks - amountsLanded - rocksPerCycle*amountOfCycles
	amountsLanded = 0
	for amountsLanded < rocksLeft {
		shape, i, height, amountsLanded, currentShape =
			move(shape, i, height, amountsLanded, currentShape, jets, shapes, landing)
	}
	fmt.Println(height + heightPerCycle*amountOfCycles)
}

func move(shape []Position, i int, height int, amountsLanded int, currentShape int,
	pattern string, shapes [5][]Position, landing map[Position]bool) ([]Position, int, int, int, int) {
	if canBePushed(shape, rune(pattern[i]), landing) {
		shape = updateShapeAfterPush(shape, rune(pattern[i]))
	}
	i = (i + 1) % len(pattern)
	if canFall(shape, landing) {
		shape = updateShapeAfterFall(shape)
	} else {
		updateLanding(shape, landing)
		height = getHeight(height, shape)
		amountsLanded += 1
		currentShape += 1
		shape = getInitialShape(shapes[currentShape%5], height+4)
	}
	return shape, i, height, amountsLanded, currentShape
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
	hash := strconv.Itoa(i) + strconv.Itoa(i%5)
	for _, diff := range diffHeight {
		hash += strconv.Itoa(diff)
	}
	return hash
}

func getHeight(height int, shape []Position) int {
	for _, position := range shape {
		height = max(height, position.x)
	}
	return height
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
		if (dir == '<' && landing[Position{position.x, position.y - 1}]) ||
			(dir == '>' && landing[Position{position.x, position.y + 1}]) {
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
		if landing[Position{position.x - 1, position.y}] {
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
