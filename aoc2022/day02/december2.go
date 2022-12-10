package day02

import (
	"fmt"
	"os"
	"strings"
)

func Day2() {
	content, _ := os.ReadFile("day02.txt")
	day2Content := string(content)
	rounds := strings.Split(day2Content, "\n")
	pointsPart1 := 0
	pointsPart2 := 0
	for _, roundData := range rounds {
		currentRound := strings.Split(roundData, " ")
		pointsPart1 += pointsStrategy1(currentRound[1], currentRound[0])
		pointsPart2 += pointsStrategy2(currentRound[1], currentRound[0])
	}
	fmt.Println(pointsPart1)
	fmt.Println(pointsPart2)
}

func pointsStrategy1(you string, other string) int {
	if you == "X" {
		return 3*btoi(other == "A") + 6*btoi(other == "C") + 1
	} else if you == "Y" {
		return 3*btoi(other == "B") + 6*btoi(other == "A") + 2
	} else if you == "Z" {
		return 3*btoi(other == "C") + 6*btoi(other == "B") + 3
	} else {
		return 0
	}
}

func pointsStrategy2(you string, other string) int {
	if you == "X" {
		return 1*btoi(other == "B") + 2*btoi(other == "C") + 3*btoi(other == "A") + 0
	} else if you == "Y" {
		return 1*btoi(other == "A") + 2*btoi(other == "B") + 3*btoi(other == "C") + 3
	} else if you == "Z" {
		return 1*btoi(other == "C") + 2*btoi(other == "A") + 3*btoi(other == "B") + 6
	} else {
		return 0
	}
}

func btoi(boolean bool) int {
	if boolean {
		return 1
	} else {
		return 0
	}
}
