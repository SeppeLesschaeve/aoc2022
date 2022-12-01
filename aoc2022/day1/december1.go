package day1

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Elf struct {
	Name string
	Cal  int
}

func Day1() {
	content, err := os.ReadFile("day1.txt")
	if err != nil {
		panic(err)
	}
	day1Content := string(content)
	elvesData := strings.Split(day1Content, "\n\n")
	var cals []Elf
	max := 0
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
		if sum > max {
			max = sum
		}
	}
	sort.Slice(cals, func(p, q int) bool {
		return cals[p].Cal > cals[q].Cal
	})
	threeCals := cals[:3]
	sumArray := sum(threeCals)
	fmt.Println(max)
	fmt.Println(sumArray)
}

func sum(array []Elf) int {
	result := 0
	for _, v := range array {
		result += v.Cal
	}
	return result
}
