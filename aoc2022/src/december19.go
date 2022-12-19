package src

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func Day19() {
	content, _ := os.ReadFile("input/day19.txt")
	day19Content := string(content)
	bluePrints := getBluePrints(day19Content)
	sumQuality := getSumOfQuality(bluePrints)
	fmt.Println(sumQuality)
	productOfGeodes := getProductOfGeodes(bluePrints)
	fmt.Println(productOfGeodes)
}

func getSumOfQuality(prints []BluePrint) int {
	sum_quality := 0
	for _, bluePrint := range prints {
		costs := [][]int{{bluePrint.ore, 0, 0, 0},
			{bluePrint.clay, 0, 0, 0},
			{bluePrint.obsidian[0], bluePrint.obsidian[1], 0, 0},
			{bluePrint.geode[0], 0, bluePrint.geode[1], 0}}
		amountMined := mine(costs, []int{1, 0, 0, 0}, 24, 1000)
		sum_quality += bluePrint.id * amountMined
	}
	return sum_quality
}

func getProductOfGeodes(prints []BluePrint) int {
	productOfGeodes := 1
	for i, bluePrint := range prints {
		if i == 3 {
			break
		}
		costs := [][]int{{bluePrint.ore, 0, 0, 0},
			{bluePrint.clay, 0, 0, 0},
			{bluePrint.obsidian[0], bluePrint.obsidian[1], 0, 0},
			{bluePrint.geode[0], 0, bluePrint.geode[1], 0}}
		amountMined := mine(costs, []int{1, 0, 0, 0}, 32, 10000)
		productOfGeodes *= amountMined
	}
	return productOfGeodes
}

func mine(costs [][]int, robots []int, num_minutes int, top int) int {
	var queue []MineState
	queue = append(queue, MineState{0, robots, []int{0, 0, 0, 0}, []int{0, 0, 0, 0}})
	maxGeodeMined := 0
	depth := 0
	for len(queue) > 0 {
		mineState := queue[0]
		queue = queue[1:]
		robots = mineState.robots
		if mineState.minutes > depth {
			sort.Slice(queue, func(i, j int) bool {
				return qualityHeuristic(queue[i]) > qualityHeuristic(queue[j])
			})
			if len(queue) > top {
				queue = queue[:top]
			}
			depth = mineState.minutes
		}

		if mineState.minutes == num_minutes {
			maxGeodeMined = max(maxGeodeMined, mineState.mined[3])
			continue
		}

		var newInventory []int
		var newMined []int
		for i := 0; i < 4; i++ {
			newInventory = append(newInventory, mineState.inventory[i]+robots[i])
			newMined = append(newMined, mineState.mined[i]+robots[i])
		}
		queue = append(queue, MineState{mineState.minutes + 1, robots, newInventory, newMined})
		for i := 0; i < 4; i++ {
			costRobot := costs[i]
			if enoughToBuild(mineState.inventory, costRobot) {
				var newRobots []int
				for _, robot := range robots {
					newRobots = append(newRobots, robot)
				}
				newRobots[i] += 1
				var newInventoryState []int
				for j := 0; j < 4; j++ {
					newInventoryState = append(newInventoryState, newInventory[j]-costRobot[j])
				}
				queue = append(queue, MineState{mineState.minutes + 1, newRobots, newInventoryState, newMined})
			}
		}
	}
	return maxGeodeMined
}

func enoughToBuild(inventory []int, costRobot []int) bool {
	for i := 0; i < 4; i++ {
		if inventory[i] < costRobot[i] {
			return false
		}
	}
	return true
}

func qualityHeuristic(mineState MineState) int {
	return 1000*mineState.mined[3] + 100*mineState.mined[2] + 10*mineState.mined[1] + mineState.mined[0]
}

func getBluePrints(content string) []BluePrint {
	var bluePrints []BluePrint
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		var id, ore, clay, obsidianOre, obsidianClay, geodeOre, geodeObsidian int
		_, _ = fmt.Sscanf(line, "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian",
			&id, &ore, &clay, &obsidianOre, &obsidianClay, &geodeOre, &geodeObsidian)
		bluePrints = append(bluePrints, BluePrint{id, ore, clay,
			[]int{obsidianOre, obsidianClay}, []int{geodeOre, geodeObsidian}})
	}
	return bluePrints
}
