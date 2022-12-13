package src

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day09() {
	content, _ := os.ReadFile("input/day09.txt")
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
	tailVisited := make(map[Position]Void)
	tailVisited[knots[len(knots)-1]] = void
	for _, action := range actions {
		for i := 0; i < action.distance; i++ {
			knots[0] = moveHead(knots[0], action.direction)
			for t := 1; t < len(knots); t++ {
				if abs(knots[t-1].x-knots[t].x) > 1 || abs(knots[t-1].y-knots[t].y) > 1 {
					knots[t] = moveTail(knots[t-1], knots[t])
					if t == len(knots)-1 {
						tailVisited[knots[t]] = void
					}
				}
			}
		}
	}
	return len(tailVisited)
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
