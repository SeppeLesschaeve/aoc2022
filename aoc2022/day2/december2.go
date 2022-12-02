package day2

import (
	"fmt"
	"os"
	"strings"
)

func Day2() {
	content, err := os.ReadFile("day2.txt")
	if err != nil {
		panic(err)
	}
	day2Content := string(content)
	rounds := strings.Split(day2Content, "\n")
	pointsPart1 := 0
	pointsPart2 := 0
	for _, roundData := range rounds {
		currentRound := strings.Split(roundData, " ")
		if isWin(currentRound[1], currentRound[0]) {
			pointsPart1 += pointsWinning(currentRound[1])
		} else if isDraw(currentRound[1], currentRound[0]) {
			pointsPart1 += pointsDraw(currentRound[1])
		} else if isLose(currentRound[1], currentRound[0]) {
			pointsPart1 += pointsLose(currentRound[1])
		}
		if currentRound[1] == "X" {
			pointsPart2 += pointsMustLose(currentRound[0])
		} else if currentRound[1] == "Y" {
			pointsPart2 += pointsMustDraw(currentRound[0])
		} else if currentRound[1] == "Z" {
			pointsPart2 += pointsMustWin(currentRound[0])
		}
	}
	fmt.Println(pointsPart1)
	fmt.Println(pointsPart2)
}

func isWin(you string, other string) bool {
	return (you == "X" && other == "C") || (you == "Y" && other == "A") || (you == "Z" && other == "B")
}

func pointsWinning(you string) int {
	if you == "X" {
		return 7
	} else if you == "Y" {
		return 8
	} else if you == "Z" {
		return 9
	}
	return 0
}

func pointsMustWin(other string) int {
	if other == "C" {
		return 7
	} else if other == "A" {
		return 8
	} else if other == "B" {
		return 9
	}
	return 0
}

func isDraw(you string, other string) bool {
	return (you == "X" && other == "A") || (you == "Y" && other == "B") || (you == "Z" && other == "C")
}

func pointsDraw(you string) int {
	if you == "X" {
		return 4
	} else if you == "Y" {
		return 5
	} else if you == "Z" {
		return 6
	}
	return 0
}

func pointsMustDraw(other string) int {
	if other == "A" {
		return 4
	} else if other == "B" {
		return 5
	} else if other == "C" {
		return 6
	}
	return 0
}

func isLose(you string, other string) bool {
	return (you == "X" && other == "B") || (you == "Y" && other == "C") || (you == "Z" && other == "A")
}

func pointsLose(you string) int {
	if you == "X" {
		return 1
	} else if you == "Y" {
		return 2
	} else if you == "Z" {
		return 3
	}
	return 0
}

func pointsMustLose(other string) int {
	if other == "B" {
		return 1
	} else if other == "C" {
		return 2
	} else if other == "A" {
		return 3
	}
	return 0
}
