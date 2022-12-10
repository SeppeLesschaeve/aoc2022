package day04

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day4() {
	content, _ := os.ReadFile("day04.txt")
	day4Content := string(content)
	elvesData := strings.Split(day4Content, "\n")
	sumPart1 := 0
	sumPart2 := 0
	for _, elfData := range elvesData {
		sections := strings.Split(elfData, ",")
		section1 := strings.Split(sections[0], "-")
		section2 := strings.Split(sections[1], "-")
		section1First, _ := strconv.Atoi(section1[0])
		section1Second, _ := strconv.Atoi(section1[1])
		section2First, _ := strconv.Atoi(section2[0])
		section2Second, _ := strconv.Atoi(section2[1])
		if doesOverlapFully(section1First, section1Second, section2First, section2Second) {
			sumPart1 += 1
		}
		if doesOverlap(section1First, section1Second, section2First, section2Second) {
			sumPart2 += 1
		}
	}
	fmt.Println(sumPart1)
	fmt.Println(sumPart2)
}

func doesOverlapFully(section1First int, section1Second int, section2First int, section2Second int) bool {
	return (section1First <= section2First && section2Second <= section1Second) ||
		(section2First <= section1First && section1Second <= section2Second)
}

func doesOverlap(section1First int, section1Second int, section2First int, section2Second int) bool {
	return section2First <= section1Second && section1First <= section2Second
}
