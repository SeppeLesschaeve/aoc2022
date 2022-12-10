package day03

import (
	"fmt"
	"os"
	"strings"
)

func Day3() {
	content, err := os.ReadFile("day03.txt")
	if err != nil {
		panic(err)
	}
	day3Content := string(content)
	elvesData := strings.Split(day3Content, "\n")
	firstSum := 0
	for _, elfData := range elvesData {
		common := findCommon(elfData[:len(elfData)/2], elfData[len(elfData)/2:])
		firstSum += priority(common[0])
	}
	i := 0
	secondSum := 0
	for i < len(elvesData) {
		commons := elvesData[i]
		rucksacks := [2]string{elvesData[i+1], elvesData[i+2]}
		for _, rucksack := range rucksacks {
			commons = strings.Join(findCommons(rucksack, commons), "")
		}
		secondSum += priority(commons[0])
		i += 3
	}
	fmt.Println(firstSum)
	fmt.Println(secondSum)
}

func priority(common uint8) int {
	if strings.ToUpper(string(common)) == string(common) {
		return int(common - 65 + 27)
	} else {
		return int(common - 97 + 1)
	}
}

func findCommon(firstCompartment string, secondCompartment string) string {
	inter := findCommons(firstCompartment, secondCompartment)
	if len(inter) != 0 {
		return inter[0]
	}
	return ""
}

func findCommons(rucksack string, commons string) (inter []string) {
	m := make(map[uint8]bool)

	for i := 0; i < len(commons); i++ {
		m[commons[i]] = true
	}

	for j := 0; j < len(rucksack); j++ {
		charString := rucksack[j]
		if m[charString] && !contains(inter, string(charString)) {
			inter = append(inter, string(charString))
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
