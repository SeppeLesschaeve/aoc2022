package day9

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Action struct {
	direction string
	distance  int
}

type Position struct {
	x int
	y int
}

func Day9() {
	content, _ := os.ReadFile("day9.txt")
	day9Content := string(content)
	actions := getActions(day9Content)
	firstPartKnots := []Position{{0, 0}, {0, 0}}
	secondPartKnots := []Position{{0, 0}, {0, 0}, {0, 0}, {0, 0},
		{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
	firstPartVisited := followActions(actions, firstPartKnots)
	secondPartVisited := followActions(actions, secondPartKnots)
	fmt.Println(firstPartVisited)
	fmt.Println(secondPartVisited)
}

func getActions(content string) []Action {
	lines := strings.Split(content, "\n")
	actions := make([]Action, len(lines))
	for i, line := range lines {
		action := strings.Split(line, " ")
		distance, _ := strconv.Atoi(action[1])
		actions[i] = Action{action[0], distance}
	}
	return actions
}

func followActions(actions []Action, knots []Position) interface{} {
	tailVisited := []Position{knots[len(knots)-1]}
	for _, action := range actions {
		for i := 0; i < action.distance; i++ {
			knots[0] = moveHead(knots[0], action.direction)
			for t := 1; t < len(knots); t++ {
				if abs(knots[t-1].x-knots[t].x) > 1 || abs(knots[t-1].y-knots[t].y) > 1 {
					knots[t] = moveTail(knots[t-1], knots[t])
					if t == len(knots)-1 {
						if isNewTailPos(knots[t], tailVisited) {
							tailVisited = append(tailVisited, knots[t])
						}
					}
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
	return Position{newCoordinate(tail.x, head.x), newCoordinate(tail.y, head.y)}
}

func newCoordinate(tailCoordinate int, headCoordinate int) int {
	if headCoordinate-tailCoordinate > 0 {
		return tailCoordinate + 1
	} else if headCoordinate-tailCoordinate == 0 {
		return tailCoordinate
	} else {
		return tailCoordinate - 1
	}
}

func abs(i int) int {
	if i >= 0 {
		return i
	} else {
		return -i
	}
}
