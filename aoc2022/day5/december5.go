package day5

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day5() {
	content, _ := os.ReadFile("day5.txt")
	day5Content := string(content)
	configAndSteps := strings.Split(day5Content, "\n\n")
	configPart1 := getConfig(configAndSteps[0])
	configPart2 := getConfig(configAndSteps[0])
	steps1 := getSteps(configAndSteps[1])
	steps2 := getSteps(configAndSteps[1])
	for _, step := range steps1 {
		for step[0] > 0 {
			changeConfigSingle(configPart1, step[1], step[2])
			step[0] -= 1
		}
	}
	for _, step := range steps2 {
		changeConfigMultiple(configPart2, step[0], step[1], step[2])
	}
	fmt.Println(getTop(configPart1))
	fmt.Println(getTop(configPart2))
}

func getTop(config [][]string) string {
	var top []string
	i := 0
	for i < len(config) {
		top = append(top, config[i][0])
		i += 1
	}
	return strings.Join(top, "")
}

func getAmountOfStacks(stacks string) int {
	amounts := strings.Fields(stacks)
	return len(amounts)
}

func firstN(s string, n int) string {
	if len(s) > n {
		return s[:n]
	}
	return s
}

func trimLeftChars(s string, n int) string {
	m := 0
	for i := range s {
		if m >= n {
			return s[i:]
		}
		m++
	}
	return s[:0]
}

func getLine(line string) []string {
	var stacks []string
	for len(line) > 0 {
		stack := firstN(line, 4)
		if stack == "    " {
			stacks = append(stacks, "")
		} else {
			stacks = append(stacks, string(stack[1]))
		}
		if len(line) == 3 {
			line = ""
		} else {
			line = trimLeftChars(line, 4)
		}
	}
	return stacks
}

func getConfig(configString string) [][]string {
	lines := strings.Split(configString, "\n")
	amountOfStacks := getAmountOfStacks(lines[len(lines)-1])
	config := make([][]string, amountOfStacks)
	i := len(lines) - 2
	for i >= 0 {
		line := getLine(lines[i])
		l := 0
		for l < len(line) {
			crate := line[l]
			if crate != "" {
				config[l] = append([]string{crate}, config[l]...)
			}
			l += 1
		}
		i -= 1
	}
	return config
}

func getStep(stepString string) []int {
	stepSplit := strings.Split(stepString, " ")
	amount, _ := strconv.Atoi(stepSplit[1])
	from, _ := strconv.Atoi(stepSplit[3])
	to, _ := strconv.Atoi(stepSplit[5])
	return []int{amount, from - 1, to - 1}
}

func getSteps(stepsString string) [][]int {
	lines := strings.Split(stepsString, "\n")
	steps := make([][]int, len(lines))
	i := 0
	for i < len(lines) {
		step := getStep(lines[i])
		steps[i] = step
		i += 1
	}
	return steps
}

func changeConfigSingle(config [][]string, from int, to int) {
	headFrom := config[from][0]
	config[from] = config[from][1:]
	config[to] = append([]string{headFrom}, config[to]...)
}

func changeConfigMultiple(config [][]string, amount int, from int, to int) {
	headFrom := config[from][:amount]
	config[from] = config[from][amount:]
	config[to] = append(headFrom, config[to]...)
}
