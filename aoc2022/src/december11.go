package src

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Day11() {
	content, _ := os.ReadFile("input/day11.txt")
	day11Content := string(content)
	monkeysSplit := strings.Split(day11Content, "\n\n")
	monkeysPart1 := getMonkeys(monkeysSplit)
	monkeysPart2 := getMonkeys(monkeysSplit)
	monkeysPart1 = simulateRounds(monkeysPart1, 20, func(i int, monkey *ThrowingMonkey) { monkey.items[i] /= 3 })
	biggest := getBiggest(monkeysPart2)
	monkeysPart2 = simulateRounds(monkeysPart2, 10000, func(i int, monkey *ThrowingMonkey) { monkey.items[i] %= biggest })
	fmt.Println(getMonkeyBusiness(monkeysPart1))
	fmt.Println(getMonkeyBusiness(monkeysPart2))
}

func simulateRounds(monkeys []*ThrowingMonkey, rounds int, trust func(int, *ThrowingMonkey)) []*ThrowingMonkey {
	for round := 0; round < rounds; round++ {
		for _, monkey := range monkeys {
			for itemInd, item := range monkey.items {
				monkey.operation(item, itemInd, monkey)
				trust(itemInd, monkey)
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

func getBiggest(monkeys []*ThrowingMonkey) int {
	biggest := 1
	for _, monkey := range monkeys {
		biggest *= monkey.div
	}
	return biggest
}

func getMonkeys(monkeysSplit []string) []*ThrowingMonkey {
	var monkeys []*ThrowingMonkey
	for _, monkeySplit := range monkeysSplit {
		monkeyInfo := strings.Split(monkeySplit, "\n")
		var items []int
		var itemsInfo, op, a string
		var div, trueMonkey, falseMonkey int
		monkeyInfo[1] = strings.ReplaceAll(monkeyInfo[1], ", ", ",")
		_, _ = fmt.Sscanf(monkeyInfo[1], "  Starting items: %s", &itemsInfo)
		_, _ = fmt.Sscanf(monkeyInfo[2], "  Operation: new = old %s %s", &op, &a)
		_, _ = fmt.Sscanf(monkeyInfo[3], "  Test: divisible by %d", &div)
		_, _ = fmt.Sscanf(monkeyInfo[4], "    If true: throw to monkey %d", &trueMonkey)
		_, _ = fmt.Sscanf(monkeyInfo[5], "    If false: throw to monkey %d", &falseMonkey)
		itemsSplit := strings.Split(itemsInfo, ",")
		for _, item := range itemsSplit {
			itemInt, _ := strconv.Atoi(item)
			items = append(items, itemInt)
		}
		if op == "*" {
			monkeys = append(monkeys, &ThrowingMonkey{items, func(item int, itemInd int, monkey *ThrowingMonkey) {
				monkey.items[itemInd] = item * getAmount(a, item)
			}, div, []int{trueMonkey, falseMonkey}, 0})
		} else if op == "+" {
			monkeys = append(monkeys, &ThrowingMonkey{items, func(item int, itemInd int, monkey *ThrowingMonkey) {
				monkey.items[itemInd] = item + getAmount(a, item)
			}, div, []int{trueMonkey, falseMonkey}, 0})
		}
	}
	return monkeys
}

func getAmount(a string, item int) int {
	if a != "old" {
		amount, _ := strconv.Atoi(a)
		return amount
	}
	return item
}

func getMonkeyBusiness(monkeys []*ThrowingMonkey) int {
	sort.Slice(monkeys, func(p, q int) bool { return monkeys[p].amountOfThrows > monkeys[q].amountOfThrows })
	return monkeys[0].amountOfThrows * monkeys[1].amountOfThrows
}
