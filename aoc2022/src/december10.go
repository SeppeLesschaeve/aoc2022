package src

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day10() {
	content, _ := os.ReadFile("input/day10.txt")
	day10Content := string(content)
	lines := strings.Split(day10Content, "\n")
	valuesOfX := getValues(lines)
	firstPartSum := getCyclesSum(valuesOfX)
	fmt.Println(firstPartSum)
	secondPartImage := getRendering(valuesOfX)
	for _, row := range secondPartImage {
		fmt.Println(row)
	}
}

func getValues(lines []string) []int {
	var valuesOfX []int
	newX := 1
	for _, line := range lines {
		if line == "noop" {
			valuesOfX = append(valuesOfX, newX)
		} else {
			addSplit := strings.Split(line, " ")
			toAddX, _ := strconv.Atoi(addSplit[1])
			valuesOfX = append(valuesOfX, newX, newX)
			newX = newX + toAddX
		}
	}
	return valuesOfX
}

func getCyclesSum(valuesOfX []int) interface{} {
	cycles := []int{20, 60, 100, 140, 180, 220}
	sum := 0
	for _, cycle := range cycles {
		sum += cycle * valuesOfX[cycle-1]
	}
	return sum
}

func getRendering(valuesOfX []int) []string {
	rendering := make([]string, 6)
	var x int
	for rowI, _ := range rendering {
		rendering[rowI] = strings.Repeat(".", 40)
	}
	for i := 0; i < 240; i++ {
		rowI := i / 40
		colI := i % 40
		x = valuesOfX[0]
		if colI >= x-1 && colI <= x+1 {
			rendering[rowI] = rendering[rowI][:colI] + "#" + rendering[rowI][colI+1:]
		}
		valuesOfX = valuesOfX[1:]
	}
	return rendering
}
