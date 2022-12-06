package day6

import (
	"fmt"
	"os"
)

func Day6() {
	content, _ := os.ReadFile("day6.txt")
	day6Content := string(content)
	unique(day6Content, 4)
	unique(day6Content, 14)

}

func unique(content string, amount int) {
	var markers []string
	for i := 0; i < amount; i++ {
		markers = append(markers, string(content[i]))
	}
	j := amount
	for ; !allDifferent(markers); j++ {
		markers = markers[1:]
		markers = append(markers, string(content[j]))
	}
	fmt.Println(j)

}

func allDifferent(markers []string) bool {
	for i := 0; i < len(markers); i++ {
		for j := i + 1; j < len(markers); j++ {
			if markers[i] == markers[j] && i != j {
				return false
			}
		}
	}
	return true
}
