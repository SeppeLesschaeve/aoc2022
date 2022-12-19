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

type Valve struct {
	name     string
	rate     int
	valvesTo map[string]Void
	paths    map[string]int
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

type CaveState struct {
	blockIndex int
	shape      []Position
	diffHeight []int
}

type Cube struct {
	x int
	y int
	z int
}

type CubeState struct {
	cubes            map[Cube]bool
	surfaceArea      int
	minCube, maxCube Cube
}

type BluePrint struct {
	id       int
	ore      int
	clay     int
	obsidian []int
	geode    []int
}

type MineState struct {
	minutes   int
	robots    []int
	inventory []int
	mined     []int
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
