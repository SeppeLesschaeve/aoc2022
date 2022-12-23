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
	if monkey.value != math.MinInt {
		return monkey.value
	}
	switch monkey.operation {
	case "+":
		monkey.value = monkeys.Resolve(monkey.left) + monkeys.Resolve(monkey.right)
	case "-":
		monkey.value = monkeys.Resolve(monkey.left) - monkeys.Resolve(monkey.right)
	case "*":
		monkey.value = monkeys.Resolve(monkey.left) * monkeys.Resolve(monkey.right)
	case "/":
		monkey.value = monkeys.Resolve(monkey.left) / monkeys.Resolve(monkey.right)
	}
	monkey.operation = ""
	return monkey.value
}

func (monkeys *Monkeys) ResolveReversed(name string) int {
	for monkeyName, yellingMonkey := range *monkeys {
		if yellingMonkey.value != math.MinInt {
			continue
		}
		if yellingMonkey.left == name {
			other := monkeys.Resolve(yellingMonkey.right)
			if monkeyName == "root" {
				return other
			} else {
				result := monkeys.ResolveReversed(monkeyName)
				switch yellingMonkey.operation {
				case "+":
					return result - other
				case "-":
					return result + other
				case "*":
					return result / other
				case "/":
					return result * other
				}
			}
			break
		} else if yellingMonkey.right == name {
			other := monkeys.Resolve(yellingMonkey.left)
			if monkeyName == "root" {
				return other
			} else {
				result := monkeys.ResolveReversed(monkeyName)
				switch yellingMonkey.operation {
				case "+":
					return result - other
				case "-":
					return result + other
				case "*":
					return result / other
				case "/":
					return result * other
				}
			}
			break
		}
	}
	return 0
}
