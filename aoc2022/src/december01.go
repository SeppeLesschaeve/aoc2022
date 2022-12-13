package src

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Day01() {
	content, _ := os.ReadFile("input/day01.txt")
	day1Content := string(content)
	elvesData := strings.Split(day1Content, "\n\n")
	var cals []Elf
	for elf, elfData := range elvesData {
		elfData := strings.Split(elfData, "\n")
		sum := 0
		for _, elfCal := range elfData {
			cal, err := strconv.Atoi(elfCal)
			if err == nil {
				sum += cal
			}
		}
		cals = append(cals, Elf{strconv.Itoa(elf), sum})
	}
	sort.Slice(cals, func(p, q int) bool {
		return cals[p].Cal > cals[q].Cal
	})
	threeCals := cals[:3]
	sumArray := sum(threeCals)
	fmt.Println(cals[0].Cal)
	fmt.Println(sumArray)
}

func sum(array []Elf) int {
	result := 0
	for _, v := range array {
		result += v.Cal
	}
	return result
}
