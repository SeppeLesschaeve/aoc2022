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
		amount := step[0]
		for amount > 0 {
			changeConfigSingle(configPart1, step[1], step[2])
			amount -= 1
		}
	}
	for _, step := range steps {
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
	i := 0
	for i <= len(lines)-2 {
		line := getLine(lines[i])
		l := 0
		for l < len(line) {
			crate := line[l]
			if crate != "" {
				config[l] = append(config[l], crate)
			}
			l += 1
		}
		i += 1
	}
	return config
}

func getStep(stepString string) []int {
	var amount, from, to int
	_, _ = fmt.Sscanf(stepString, "move %d from %d to %d", &amount, &from, &to)
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

func pop(stack []string, amount int) ([]string, []string) {
	if len(stack) < amount {
		panic("stack is empty")
	}
	var popped []string
	amountToPop := 0
	for amountToPop < amount {
		popped = append(popped, stack[amountToPop])
		amountToPop += 1
	}
	stack = stack[amount:]
	return popped, stack
}

func changeConfigSingle(config [][]string, from int, to int) {
	var headFrom []string
	headFrom, config[from] = pop(config[from], 1)
	config[to] = append(headFrom, config[to]...)
}

func changeConfigMultiple(config [][]string, amount int, from int, to int) {
	var headFrom []string
	headFrom, config[from] = pop(config[from], amount)
	config[to] = append(headFrom, config[to]...)
}
