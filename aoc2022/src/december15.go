package src

import (
	"fmt"
	"os"
	"strings"
)

func Day15() {
	content, _ := os.ReadFile("input/day15.txt")
	day15Content := string(content)
	sensors, beacons := getSensorsAndBeacons(day15Content)
	noBeaconFirstPositions := getNoBeaconFirstPositions(sensors, beacons)
	pos := getNoBeaconSecondPosition(sensors, beacons)
	tuningFrequency := getTuningFrequency(pos)
	fmt.Println(len(noBeaconFirstPositions))
	fmt.Println(tuningFrequency)
}

func getTuningFrequency(pos Position) int {
	return 4000000*pos.x + pos.y
}

func getNoBeaconFirstPositions(sensors []Position, beacons []Position) map[Position]Void {
	noBeaconPositions := make(map[Position]Void)
	for i := 0; i < len(sensors); i++ {
		noBeaconPositions = getNoBeaconPosOn2000000(sensors[i], beacons[i], noBeaconPositions)
	}
	return noBeaconPositions
}

func getNoBeaconPosOn2000000(sensor Position, beacon Position, noBeaconPositions map[Position]Void) map[Position]Void {
	distance := getHammingDistance(sensor, beacon)
	j := -distance
	for i := 0; i <= distance; i++ {
		col := sensor.y + j
		if col == 2000000 {
			for row := sensor.x - i; row <= sensor.x+i; row++ {
				if !(row == beacon.x && col == beacon.y) {
					noBeaconPositions[Position{row, col}] = void
				}
			}
		}
		j++
	}
	for i := distance - 1; i >= 0; i-- {
		col := sensor.y + j
		if col == 2000000 {
			for row := sensor.x - i; row <= sensor.x+i; row++ {
				if !(row == beacon.x && col == beacon.y) {
					noBeaconPositions[Position{row, col}] = void
				}
			}
		}
		j++
	}
	return noBeaconPositions
}

func getNoBeaconSecondPosition(sensors []Position, beacons []Position) Position {
	var pos Position
	for y := 0; y <= 4000000; y++ {
		for x := 0; x <= 4000000; x++ {
			pos = Position{x, y}
			if isNotInRange(sensors, beacons, pos) {
				return pos
			} else {
				for i := 0; i < len(sensors); i++ {
					distance := getHammingDistance(sensors[i], pos)
					if distance <= getHammingDistance(sensors[i], beacons[i]) {
						x += getHammingDistance(sensors[i], beacons[i]) - distance
						break
					}
				}
			}
		}
	}
	return pos
}

func isNotInRange(sensors []Position, beacons []Position, position Position) bool {
	for i := 0; i < len(sensors); i++ {
		if getHammingDistance(sensors[i], position) <= getHammingDistance(sensors[i], beacons[i]) {
			return false
		}
	}
	return true
}

func getSensorsAndBeacons(content string) ([]Position, []Position) {
	var sensors, beacons []Position
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		var sensorX, sensorY, beaconX, beaconY int
		_, _ = fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
			&sensorX, &sensorY, &beaconX, &beaconY)
		sensors = append(sensors, Position{sensorX, sensorY})
		beacons = append(beacons, Position{beaconX, beaconY})
	}
	return sensors, beacons
}
