package src

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day20() {
	content, _ := os.ReadFile("input/day20.txt")
	day20Content := string(content)
	groveCoordinates, groveZero := getGroveCoordinates(day20Content)
	setNextAndPrevious(groveCoordinates)
	mix(groveCoordinates)
	fmt.Println(sumGrove(groveZero))
	for _, n := range groveCoordinates {
		n.coordinate *= 811589153
	}
	setNextAndPrevious(groveCoordinates)
	for i := 0; i < 10; i++ {
		mix(groveCoordinates)
	}
	fmt.Println(sumGrove(groveZero))
}

func getGroveCoordinates(content string) ([]*groveCoordinate, *groveCoordinate) {
	var numbers []*groveCoordinate
	var zero *groveCoordinate
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		number , _ := strconv.Atoi(line)
		num := &groveCoordinate{coordinate: number}
		numbers = append(numbers, num)
		if number == 0 {
			zero = num
		}
	}
	return numbers, zero
}

func moveCoordinate(n *groveCoordinate, places int) *groveCoordinate {
	for places < 0 {
		n = n.previous
		places++
	}
	for places > 0 {
		n = n.next
		places--
	}
	return n
}

func mix(m []*groveCoordinate) {
	for _, n := range m {
		t := n.previous
		n.previous.next = n.next
		n.next.previous = n.previous

		t = moveCoordinate(t, n.coordinate%(len(m)-1))
		n.previous = t
		n.next = t.next
		n.previous.next = n
		n.next.previous = n
	}
}

func setNextAndPrevious(m []*groveCoordinate) {
	m[len(m)-1].next = m[0]
	m[0].previous = m[len(m)-1]
	for i := 1; i < len(m); i++ {
		m[i-1].next = m[i]
		m[i].previous = m[i-1]
	}
}

func sumGrove(groveZero *groveCoordinate) int {
	s := 0
	for i, n := 0, groveZero; i < 3; i++ {
		n = moveCoordinate(n, 1000)
		s += n.coordinate
	}
	return s
}
