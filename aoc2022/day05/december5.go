package day05

import (
	"fmt"
	"os"
	"strings"
)

func Day5() {
	content, _ := os.ReadFile("day05.txt")
	day5Content := string(content)
	configAndSteps := strings.Split(day5Content, "\n\n")
	configPart1 := getConfig(configAndSteps[0])
	configPart2 := getConfig(configAndSteps[0])
	steps := getSteps(configAndSteps[1])
	for _, step := range steps {
		for amount := step[0]; amount > 0; amount-- {
			changeConfig(configPart1, 1, step[1], step[2])
		}
	}
	for _, step := range steps {
		changeConfig(configPart2, step[0], step[1], step[2])
	}
	fmt.Println(getTop(configPart1))
	fmt.Println(getTop(configPart2))
}

func getTop(config [][]string) string {
	var top []string
	for i := 0; i < len(config); i++ {
		top = append(top, config[i][0])
	}
	return strings.Join(top, "")
}

func getConfig(configString string) [][]string {
	lines := strings.Split(configString, "\n")
	config := make([][]string, len(strings.Fields(lines[len(lines)-1])))
	for i := 0; i <= len(lines)-2; i++ {
		for l := 0; len(lines[i]) > 0; l++ {
			if len(lines[i]) >= 4 {
				if lines[i][:4] != "    " {
					config[l] = append(config[l], string(lines[i][1]))
				}
				lines[i] = lines[i][4:]
			} else {
				config[l] = append(config[l], string(lines[i][1]))
				break
			}
		}
	}
	return config
}

func getSteps(stepsString string) [][]int {
	lines := strings.Split(stepsString, "\n")
	steps := make([][]int, len(lines))
	for i, line := range lines {
		var amount, from, to int
		_, _ = fmt.Sscanf(line, "move %d from %d to %d", &amount, &from, &to)
		steps[i] = []int{amount, from - 1, to - 1}
	}
	return steps
}

func changeConfig(config [][]string, amount int, from int, to int) {
	var headFrom []string
	for amountToPop := 0; amountToPop < amount; amountToPop++ {
		headFrom = append(headFrom, config[from][amountToPop])
	}
	config[from] = config[from][amount:]
	config[to] = append(headFrom, config[to]...)
}
