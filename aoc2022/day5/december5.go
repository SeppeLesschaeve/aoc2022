package day5

import (
	"fmt"
	"os"
	"strings"
)

func Day5() {
	content, _ := os.ReadFile("day5.txt")
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

func getLine(line string) []string {
	var stacks []string
	for len(line) > 0 {
		if len(line) >= 4 {
			stack := line[:4]
			if stack == "    " {
				stacks = append(stacks, "")
			} else {
				stacks = append(stacks, string(stack[1]))
			}
			line = line[4:]
		} else {
			stacks = append(stacks, string(line[1]))
			break
		}
	}
	return stacks
}

func getConfig(configString string) [][]string {
	lines := strings.Split(configString, "\n")
	config := make([][]string, len(strings.Fields(lines[len(lines)-1])))
	for i := 0; i <= len(lines)-2; i++ {
		line := getLine(lines[i])
		for l, crate := range line {
			if crate != "" {
				config[l] = append(config[l], crate)
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
