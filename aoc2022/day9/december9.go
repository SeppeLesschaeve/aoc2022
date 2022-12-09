package day9

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Action struct {
	dir    string
	amount int
}

type Position struct {
	x int
	y int
}

func Day9() {
	content, _ := os.ReadFile("day9.txt")
	day9Content := string(content)
	actions := getActions(day9Content)
	firstPartTail := []Position{{0, 0}}
	secondPartTail := []Position{{0, 0}, {0, 0}, {0, 0},
		{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
	firstPartVisited := followActions(actions, firstPartTail)
	secondPartVisited := followActions(actions, secondPartTail)
	fmt.Println(firstPartVisited)
	fmt.Println(secondPartVisited)
}

func getActions(content string) []Action {
	lines := strings.Split(content, "\n")
	actions := make([]Action, len(lines))
	for i, line := range lines {
		action := strings.Split(line, " ")
		amount, _ := strconv.Atoi(action[1])
		actions[i] = Action{action[0], amount}
	}
	return actions
}

func followActions(actions []Action, tail []Position) interface{} {
	head := Position{0, 0}
	tailVisited := []Position{tail[len(tail)-1]}
	for _, action := range actions {
		for i := 0; i < action.amount; i++ {
			head = moveHead(head, action.dir)
			tail[0] = moveTail(head, tail[0])
			for t := 1; t < len(tail); t++ {
				tail[t] = moveTail(tail[t-1], tail[t])
				if t == len(tail)-1 {
					if isNewTailPos(tail[t], tailVisited) {
						tailVisited = append(tailVisited, tail[t])
					}
				}
			}
			if len(tail) == 1 {
				if isNewTailPos(tail[0], tailVisited) {
					tailVisited = append(tailVisited, tail[0])
				}
			}
		}
	}
	return len(tailVisited)
}

func isNewTailPos(pos Position, visited []Position) bool {
	for _, position := range visited {
		if pos.x == position.x && pos.y == position.y {
			return false
		}
	}
	return true
}

func moveHead(pos Position, dir string) Position {
	if dir == "U" {
		return Position{pos.x, pos.y + 1}
	} else if dir == "D" {
		return Position{pos.x, pos.y - 1}
	} else if dir == "L" {
		return Position{pos.x - 1, pos.y}
	} else if dir == "R" {
		return Position{pos.x + 1, pos.y}
	}
	return pos
}

func moveTail(head Position, tail Position) Position {
	deltaX := head.x - tail.x
	deltaY := head.y - tail.y

	if abs(deltaX) > 1 || abs(deltaY) > 1 {
		return Position{tail.x + sign(deltaX), tail.y + sign(deltaY)}
	}
	return tail
}

func sign(x int) int {
	if x < 0 {
		return -1
	}
	if x == 0 {
		return 0
	}
	return 1
}

func abs(i int) int {
	if i >= 0 {
		return i
	} else {
		return -i
	}
}
