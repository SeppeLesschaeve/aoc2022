package src

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func (monkeys *Monkeys) confirmYell(monkey YellingMonkey, yelling map[string]int) {
	delete(*monkeys, monkey.name)
	yelling[monkey.name] = monkey.value
}

func yell(monkey YellingMonkey, left int, right int) YellingMonkey {
	value := 0
	switch monkey.operation {
	case "+":
		value = left + right
	case "-":
		value = left - right
	case "*":
		value = left * right
	case "/":
		value = left / right
	case "=":
		if left == right {
			value = 1
		} else {
			value = 0
		}
	}
	monkey.value = value
	monkey.operation = ""
	return monkey
}

func (monkeys *Monkeys) iterateBusy(yelling map[string]int) {
	for len(*monkeys) > 0 {
		for _, monkey := range *monkeys {
			if monkey.operation == "" {
				monkeys.confirmYell(monkey, yelling)
			} else {
				valLeft, okLeft := yelling[monkey.left]
				valRight, okRight := yelling[monkey.right]
				if okLeft && okRight {
					(*monkeys)[monkey.name] = yell(monkey, valLeft, valRight)
				}
			}
		}
	}
}

func Day21() {
	content, _ := os.ReadFile("input/day21.txt")
	day21Content := string(content)
	monkeys := getYellingMonkeys(day21Content)
	yelling := make(map[string]int)
	monkeys.iterateBusy(yelling)
	fmt.Println(yelling["root"])
	monkeys = getYellingMonkeys(day21Content)
	you := monkeys.ResolveReversed("humn")
	fmt.Println(you)
}

func getYellingMonkeys(content string) Monkeys {
	monkeys := make(Monkeys)
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		colonSplit := strings.Split(line, ": ")
		var monkey YellingMonkey
		monkey.name = colonSplit[0]
		opSplit := strings.Fields(colonSplit[1])
		if len(opSplit) == 1 {
			monkey.value, _ = strconv.Atoi(opSplit[0])
		} else {
			monkey.value = math.MinInt
			monkey.left = opSplit[0]
			monkey.operation = opSplit[1]
			monkey.right = opSplit[2]
		}
		monkeys[monkey.name] = monkey
	}
	return monkeys

}

func (monkeys *Monkeys) Resolve(name string) int {
	monkey := (*monkeys)[name]
	// Monkey has a number
	if monkey.value != math.MinInt {
		return monkey.value
	}
	// Calculate this monkey's number and cache/return it
	switch monkey.operation {
	case "+": // result = op1 + op2
		monkey.value = monkeys.Resolve(monkey.left) + monkeys.Resolve(monkey.right)
	case "-": // result = op1 - op2
		monkey.value = monkeys.Resolve(monkey.left) - monkeys.Resolve(monkey.right)
	case "*": // result = op1 * op2
		monkey.value = monkeys.Resolve(monkey.left) * monkeys.Resolve(monkey.right)
	case "/": // result = op1 / op2
		monkey.value = monkeys.Resolve(monkey.left) / monkeys.Resolve(monkey.right)
	}
	monkey.operation = ""
	return monkey.value
}

func (monkeys *Monkeys) ResolveReversed(name string) int {
	// Find the monkey waiting for the given name/monkey
	for monkeyName, yellingMonkey := range *monkeys {
		// Skip non-computed monkeys
		if yellingMonkey.value != math.MinInt {
			continue
		}
		if yellingMonkey.left == name {
			// Solve the second operand the regular way
			other := monkeys.Resolve(yellingMonkey.right)
			if monkeyName == "root" {
				// This is root, so first operand is the same
				return other
			} else {
				// Get the expected result for this monkey
				result := monkeys.ResolveReversed(monkeyName)
				// Reverse the operation and return the first operand
				switch yellingMonkey.operation {
				case "+": // result = x + other
					return result - other
				case "-": // result = x - other
					return result + other
				case "*": // result = x * other
					return result / other
				case "/": // result = x / other
					return result * other
				}
			}
			break
		} else if yellingMonkey.right == name {
			// Solve the first operand the regular way
			other := monkeys.Resolve(yellingMonkey.left)
			if monkeyName == "root" {
				// This is root, so second operand is the same
				return other
			} else {
				// Get the expected result for this monkey
				result := monkeys.ResolveReversed(monkeyName)
				// Reverse the operation and return the second operand
				switch yellingMonkey.operation {
				case "+": // result = x + other
					return result - other
				case "-": // result = x - other
					return result + other
				case "*": // result = x * other
					return result / other
				case "/": // result = x / other
					return result * other
				}
			}
			break
		}
	}
	return 0
}
