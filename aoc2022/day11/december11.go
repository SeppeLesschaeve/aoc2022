package day11

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Operation struct {
	amount string
	op     string
}

type Monkey struct {
	items          []int
	operation      Operation
	div            int
	throwsTo       []int
	amountOfThrows int
}

func Day11() {
	content, _ := os.ReadFile("day11.txt")
	day11Content := string(content)
	monkeysSplit := strings.Split(day11Content, "\n\n")
	monkeysPart1 := getMonkeys(monkeysSplit)
	monkeysPart2 := getMonkeys(monkeysSplit)
	monkeysPart1 = simulateRounds(monkeysPart1, 20, true)
	monkeysPart2 = simulateRounds(monkeysPart2, 10000, false)
	monkeyBusinessPart1 := getMonkeyBusiness(monkeysPart1)
	monkeyBusinessPart2 := getMonkeyBusiness(monkeysPart2)
	fmt.Println(monkeyBusinessPart1)
	fmt.Println(monkeyBusinessPart2)
}

func simulateRounds(monkeys []*Monkey, rounds int, trust bool) []*Monkey {
	biggest := getBiggest(monkeys)
	for round := 0; round < rounds; round++ {
		for _, monkey := range monkeys {
			for itemInd, item := range monkey.items {
				var amount int
				if monkey.operation.amount == "old" {
					amount = item
				} else {
					amount, _ = strconv.Atoi(monkey.operation.amount)
				}
				if monkey.operation.op == "*" {
					monkey.items[itemInd] = item * amount
				} else if monkey.operation.op == "+" {
					monkey.items[itemInd] = item + amount
				}
				if trust {
					monkey.items[itemInd] /= 3
				} else {
					monkey.items[itemInd] %= biggest
				}
				if monkey.items[itemInd]%monkey.div == 0 {
					monkeys[monkey.throwsTo[0]].items = append(monkeys[monkey.throwsTo[0]].items, monkey.items[itemInd])
				} else {
					monkeys[monkey.throwsTo[1]].items = append(monkeys[monkey.throwsTo[1]].items, monkey.items[itemInd])
				}
			}
			monkey.amountOfThrows += len(monkey.items)
			monkey.items = nil
		}
	}
	return monkeys
}

func getBiggest(monkeys []*Monkey) int {
	biggest := 1
	for _, monkey := range monkeys {
		biggest *= monkey.div
	}
	return biggest
}

func getMonkeys(monkeysSplit []string) []*Monkey {
	var monkeys []*Monkey
	for _, monkeySplit := range monkeysSplit {
		monkeyInfo := strings.Split(monkeySplit, "\n")
		items := getItems(monkeyInfo[1])
		operation := getOperation(monkeyInfo[2])
		div := getDiv(monkeyInfo[3])
		throwsTo := getThrowsTo(monkeyInfo[4], monkeyInfo[5])
		monkeys = append(monkeys, &Monkey{items, operation, div, throwsTo, 0})
	}
	return monkeys
}

func getItems(itemsInfo string) []int {
	var items []int
	itemsInfo = itemsInfo[18:]
	itemsSplit := strings.Split(itemsInfo, ", ")
	for _, item := range itemsSplit {
		itemInt, _ := strconv.Atoi(item)
		items = append(items, itemInt)
	}
	return items
}

func getOperation(itemsInfo string) Operation {
	var op, amount string
	_, _ = fmt.Sscanf(itemsInfo[13:], "new = old %s %s", &op, &amount)
	return Operation{amount, op}
}

func getDiv(itemsInfo string) int {
	var div int
	_, _ = fmt.Sscanf(itemsInfo, "  Test: divisible by %d", &div)
	return div
}

func getThrowsTo(trueThrow string, falseThrow string) []int {
	var trueMonkey, falseMonkey int
	_, _ = fmt.Sscanf(trueThrow, "    If true: throw to monkey %d", &trueMonkey)
	_, _ = fmt.Sscanf(falseThrow, "    If false: throw to monkey %d", &falseMonkey)
	return []int{trueMonkey, falseMonkey}
}

func getMonkeyBusiness(monkeys []*Monkey) int {
	sort.Slice(monkeys, func(p, q int) bool {
		return monkeys[p].amountOfThrows > monkeys[q].amountOfThrows
	})
	return monkeys[0].amountOfThrows * monkeys[1].amountOfThrows
}
