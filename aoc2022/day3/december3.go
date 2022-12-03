package day3

import (
	"fmt"
	"os"
	"strings"
)

func Day3() {
	content, err := os.ReadFile("day3.txt")
	if err != nil {
		panic(err)
	}
	day3Content := string(content)
	elvesData := strings.Split(day3Content, "\n")
	firstSum := 0
	for _, elfData := range elvesData {
		firstCompartment := elfData[:len(elfData)/2]
		secondCompartment := elfData[len(elfData)/2:]
		common := findCommon(firstCompartment, secondCompartment)
		firstSum += priority(common)
	}
	i := 0
	secondSum := 0
	for i < len(elvesData) {
		commons := elvesData[i]
		rucksacks := [3]string{elvesData[i], elvesData[i+1], elvesData[i+2]}
		for _, rucksack := range rucksacks {
			common := findCommons(rucksack, commons)
			commons = strings.Join(common, "")
		}
		secondSum += priority(commons)
		i += 3
	}
	fmt.Println(firstSum)
	fmt.Println(secondSum)
}

func priority(common string) int {
	for i, comm := range common {
		if i == 0 {
			if strings.ToUpper(string(comm)) == common {
				return int(comm - 65 + 27)
			} else {
				return int(comm - 97 + 1)
			}
		}
	}
	return 0
}

func findCommon(firstCompartment string, secondCompartment string) string {
	for _, first := range firstCompartment {
		for _, second := range secondCompartment {
			if second == first {
				return string(first)
			}
		}
	}
	return ""
}

func findCommons(rucksack string, commons string) (inter []string) {
	ruckSackSlice := strings.Split(rucksack, ",")
	commonsSlice := strings.Split(commons, ",")
	m := make(map[string]bool)

	for _, item := range commonsSlice {
		for _, char := range item {
			m[string(char)] = true
		}
	}

	for _, item := range ruckSackSlice {
		for _, char := range item {
			charString := string(char)
			if _, ok := m[charString]; ok && !contains(inter, charString) {
				inter = append(inter, charString)
			}
		}

	}
	return
}

func contains(slice []string, charString string) bool {
	for _, v := range slice {
		if v == charString {
			return true
		}
	}
	return false
}
