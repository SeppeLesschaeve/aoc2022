package day4

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day4() {
	content, _ := os.ReadFile("day4.txt")
	day4Content := string(content)
	elvesData := strings.Split(day4Content, "\n")
	sumPart1 := 0
	sumPart2 := 0
	for _, elfData := range elvesData {
		sections := strings.Split(elfData, ",")
		section1 := strings.Split(sections[0], "-")
		section2 := strings.Split(sections[1], "-")
		if doesOverlapFully(section1, section2) {
			sumPart1 += 1
		}
		if doesOverlap(section1, section2) {
			sumPart2 += 1
		}
	}
	fmt.Println(sumPart1)
	fmt.Println(sumPart2)
}

func doesOverlapFully(section1 []string, section2 []string) bool {
	section1First, _ := strconv.Atoi(section1[0])
	section1Second, _ := strconv.Atoi(section1[1])
	section2First, _ := strconv.Atoi(section2[0])
	section2Second, _ := strconv.Atoi(section2[1])
	if section1First <= section2First && section2Second <= section1Second {
		return true
	} else if section2First <= section1First && section1Second <= section2Second {
		return true
	} else {
		return false
	}
}

func doesOverlap(section1 []string, section2 []string) bool {
	section1First, _ := strconv.Atoi(section1[0])
	section1Second, _ := strconv.Atoi(section1[1])
	section2First, _ := strconv.Atoi(section2[0])
	section2Second, _ := strconv.Atoi(section2[1])
	if section1First <= section2Second && section2First <= section1Second {
		return true
	} else if section2First <= section1Second && section1First <= section2Second {
		return true
	} else {
		return false
	}
}
