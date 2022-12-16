package src

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Day16() {
	content, _ := os.ReadFile("input/day16.txt")
	day16Content := string(content)
	valves := getValves(day16Content)
	valveNamesNonZero := getValveNamesNonZero(valves)
	valves = bfsAll(valveNamesNonZero, valves)
	opened := make(map[string]Void)
	opened["AA"] = void
	best1 := search1(opened, 0, "AA", 29, 0, valves)
	fmt.Println(best1)
	best2 := search2(opened, 0, "AA", 25, 0, valves, false)
	fmt.Println(best2)
}

func getValves(content string) map[string]Valve {
	valves := make(map[string]Valve)
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		ratesAndValves := strings.Split(line, "; ")
		ratesSplit := strings.Split(ratesAndValves[0], " ")
		name := ratesSplit[1]
		rateEquals := strings.Split(ratesSplit[4], "=")
		rate, _ := strconv.Atoi(rateEquals[1])
		valvesTogether := strings.Replace(ratesAndValves[1], ", ", ",", -1)
		valvesSplit := strings.Split(valvesTogether, " ")
		valvesTos := strings.Split(valvesSplit[4], ",")
		valvesTo := make(map[string]Void)
		for _, valveTo := range valvesTos {
			valvesTo[valveTo] = void
		}
		valve := Valve{name, rate, valvesTo, make(map[string]int)}
		valves[valve.name] = valve
	}
	return valves
}

func getValveNamesNonZero(valves map[string]Valve) []string {
	var valveNamesNonZero []string
	for _, valve := range valves {
		if valve.rate != 0 {
			valveNamesNonZero = append(valveNamesNonZero, valve.name)
		}
	}
	sort.Slice(valveNamesNonZero, func(i, j int) bool {
		return valveNamesNonZero[i] < valveNamesNonZero[j]
	})
	return valveNamesNonZero
}

func bfs(valvesTo map[string]Void, endValveName string, valves map[string]Valve) int {
	depth := 1
	for true {
		nextValvesTo := make(map[string]Void)
		for valveTo := range valvesTo {
			if valveTo == endValveName {
				return depth
			}
			for valveValveTo := range valves[valveTo].valvesTo {
				nextValvesTo[valveValveTo] = void
			}
		}
		valvesTo = nextValvesTo
		depth += 1
	}
	return depth
}

func bfsAll(valveNamesNonZero []string, valves map[string]Valve) map[string]Valve {
	valveNamesNonZeroAndStart := append(valveNamesNonZero, "AA")
	for _, valveNameAndStart := range valveNamesNonZeroAndStart {
		for _, valveName := range valveNamesNonZero {
			if valveName != valveNameAndStart {
				valves[valveNameAndStart].paths[valveName] = bfs(valves[valveNameAndStart].valvesTo, valveName, valves)
			}
		}
	}
	return valves
}

func search1(opened map[string]Void, flowed int, currentValveName string, minutesToGo, best int, valves map[string]Valve) int {
	if flowed > best {
		best = flowed
	}
	if minutesToGo <= 0 {
		return best
	}
	_, okValves := opened[currentValveName]
	if !okValves {
		newOpened := make(map[string]Void)
		for valve, _ := range opened {
			newOpened[valve] = void
		}
		newOpened[currentValveName] = void
		best = search1(newOpened, flowed+valves[currentValveName].rate*minutesToGo, currentValveName, minutesToGo-1, best, valves)
	} else {
		for path, _ := range valves[currentValveName].paths {
			_, okOpened := opened[path]
			if !okOpened {
				best = search1(opened, flowed, path, minutesToGo-valves[currentValveName].paths[path], best, valves)
			}
		}
	}
	return best
}

func search2(opened map[string]Void, flowed int, currentValveName string, minutesToGo, best int, valves map[string]Valve, turn bool) int {
	if flowed > best {
		best = flowed
	}
	if minutesToGo <= 0 {
		return best
	}
	_, okValves := opened[currentValveName]
	if !okValves {
		newOpened := make(map[string]Void)
		for valve, _ := range opened {
			newOpened[valve] = void
		}
		newOpened[currentValveName] = void
		best = search2(newOpened, flowed+valves[currentValveName].rate*minutesToGo, currentValveName, minutesToGo-1, best, valves, turn)
		if !turn {
			best = search2(newOpened, flowed+valves[currentValveName].rate*minutesToGo, "AA", 25, best, valves, true)
		}
	} else {
		for path, _ := range valves[currentValveName].paths {
			_, okOpened := opened[path]
			if !okOpened {
				best = search2(opened, flowed, path, minutesToGo-valves[currentValveName].paths[path], best, valves, turn)
			}
		}
	}
	return best
}
