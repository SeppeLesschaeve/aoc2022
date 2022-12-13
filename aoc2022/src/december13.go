package src

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readPacket(input string) Packet {
	root := Packet{-1, []*Packet{}, nil}
	temp := &root
	var currentNumber string
	for _, r := range input {
		switch r {
		case '[':
			temp = updateTemp(temp)
		case ']':
			temp, currentNumber = updateTempWithNumber(temp, currentNumber)
		case ',':
			temp, currentNumber = updateTempWithNumber(temp, currentNumber)
			temp = updateTemp(temp)
		default:
			currentNumber += string(r)
		}
	}
	return root
}

func updateTemp(temp *Packet) *Packet {
	newPacket := Packet{-1, []*Packet{}, temp}
	temp.elements = append(temp.elements, &newPacket)
	return &newPacket
}

func updateTempWithNumber(temp *Packet, currentNumber string) (*Packet, string) {
	if len(currentNumber) > 0 {
		number, _ := strconv.Atoi(currentNumber)
		temp.value = number
		return temp.root, ""
	}
	return temp.root, currentNumber
}

func areOrdered(first, second Packet) int {
	switch {
	case len(first.elements) == 0 && len(second.elements) == 0: //Single
		return sign(first.value - second.value)
	case first.value >= 0:
		return areOrdered(Packet{-1, []*Packet{&first}, nil}, second)
	case second.value >= 0:
		return areOrdered(first, Packet{-1, []*Packet{&second}, nil})
	default:
		var i int
		for i = 0; i < len(first.elements) && i < len(second.elements); i++ {
			ordered := areOrdered(*first.elements[i], *second.elements[i])
			if ordered != 0 {
				return ordered
			}
		}
		if i < len(first.elements) {
			return 1
		} else if i < len(second.elements) {
			return -1
		}
	}
	return 0
}

func Day13() {
	content, _ := os.ReadFile("input/day13.txt")
	day13Content := string(content)
	pairs := strings.Split(day13Content, "\n\n")
	var packets []Packet
	for _, pair := range pairs {
		packet := strings.Split(pair, "\n")
		packet1 := readPacket(packet[0])
		packet2 := readPacket(packet[1])
		packets = append(packets, packet1, packet2)
	}
	fmt.Println(sumOfIndicesOfOrderedPairs(packets))
	fmt.Println(productOfIndicesOf2and6Ordered(packets))
}

func sumOfIndicesOfOrderedPairs(packets []Packet) int {
	result := 0
	for i := 0; i < len(packets); i += 2 {
		if areOrdered(packets[i], packets[i+1]) == -1 {
			result += (i / 2) + 1
		}
	}
	return result
}

func productOfIndicesOf2and6Ordered(packets []Packet) int {
	packetSingle2 := readPacket("[[2]]")
	packets = append(packets, packetSingle2)
	packetSingle6 := readPacket("[[6]]")
	packets = append(packets, packetSingle6)
	sort.Slice(packets, func(i, j int) bool {
		return areOrdered(packets[i], packets[j]) <= 0
	})
	result := 1
	for i, packet := range packets {
		if packet.value == -1 && len(packet.elements) == 1 {
			packetSub := packet.elements[0]
			if packetSub.value == -1 && len(packetSub.elements) == 1 {
				packetSubSub := packetSub.elements[0]
				if len(packetSubSub.elements) == 0 && (packetSubSub.value == 2 || packetSubSub.value == 6) {
					result *= i + 1
				}
			}
		}
	}
	return result
}
