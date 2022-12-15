package src

type Elf struct {
	Name string
	Cal  int
}

type Entity interface {
	GetSize() int
}

type File struct {
	name string
	Size int
}

type Directory struct {
	name        string
	subEntities []*Entity
}

type Void struct{}

var void Void

type Action struct {
	direction string
	distance  int
}
type Position struct {
	x int
	y int
}

func getHammingDistance(from Position, to Position) int {
	return abs(to.x-from.x) + abs(to.y-from.y)
}

type Monkey struct {
	items          []int
	operation      func(int, int, *Monkey)
	div            int
	throwsTo       []int
	amountOfThrows int
}

type Packet struct {
	value    int
	elements []*Packet
	root     *Packet
}

func abs(i int) int {
	if i >= 0 {
		return i
	} else {
		return -i
	}
}

func sign(dif int) int {
	if dif < 0 {
		return -1
	} else if dif == 0 {
		return 0
	}
	return 1
}

func min(i int, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

func max(i int, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
