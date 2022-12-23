package src

import "fmt"

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

type ThrowingMonkey struct {
	items          []int
	operation      func(int, int, *ThrowingMonkey)
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

type groveCoordinate struct {
	coordinate     int
	next, previous *groveCoordinate
}

type Monkeys map[string]YellingMonkey
type YellingMonkey struct {
	name      string
	value     int
	operation string
	left      string
	right     string
}

type PosDir struct {
	pos Position
	dir int
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

type Interval struct{ Min, Max int }

func (i Interval) String() string {
	return fmt.Sprintf("[%d, %d]", i.Min, i.Max)
}

func (i Interval) Len() int {
	return i.Max - i.Min + 1
}

func (i Interval) Negate() Interval {
	return Interval{-i.Max, -i.Min}
}

func (i Interval) Add(b Interval) Interval {
	return Interval{i.Min + b.Min, i.Max + b.Max}
}

func (i Interval) Sub(b Interval) Interval {
	min := i.Min - b.Min
	max := i.Max - b.Max
	if min < max {
		return Interval{min, max}
	} else {
		return Interval{max, min}
	}
}

func (i Interval) Contains(x int) bool {
	return i.Min <= x && x <= i.Max
}

func minmax(a, b, c, d int) (int, int) {
	var min1, max1, min2, max2 int
	if a <= b {
		min1 = a
		max1 = b
	} else {
		min1 = b
		max1 = a
	}
	if c <= d {
		min2 = c
		max2 = d
	} else {
		min2 = d
		max2 = c
	}
	return min(min1, min2), max(max1, max2)
}

func (i Interval) Mul(b Interval) Interval {
	if i.Min >= 0 && b.Min >= 0 {
		return Interval{i.Min * b.Min, i.Max * b.Max}
	} else if i.Max <= 0 && b.Max <= 0 {
		return Interval{i.Max * b.Max, i.Min * b.Min}
	}

	min, max := minmax(i.Min*b.Min, i.Min*b.Max, i.Max*b.Min, i.Max*b.Max)
	return Interval{min, max}
}

func (i Interval) Div(b Interval) Interval {
	if i.Min >= 0 && b.Min >= 0 {
		return Interval{i.Min / b.Max, i.Max / b.Min}
	} else if i.Max <= 0 && b.Max <= 0 {
		return Interval{i.Max / b.Min, i.Min / b.Max}
	}

	min, max := minmax(i.Min/b.Min, i.Min/b.Max, i.Max/b.Min, i.Max/b.Max)
	return Interval{min, max}
}

func (i Interval) Inter(b Interval) Interval {
	if i.Max < b.Min || b.Max < i.Min {
		return Interval{0, -1}
	}
	minimum := max(i.Min, b.Min)
	max := min(i.Max, b.Max)
	return Interval{minimum, max}
}

func (i Interval) union(b Interval) Interval {
	min := min(i.Min, b.Min)
	max := max(i.Max, b.Max)
	return Interval{min, max}
}

func (i Interval) Mod1(m int) Interval {
	a := i.Min
	b := i.Max
	switch {
	case a > b || m == 0:
		return Interval{0, -1}
	case b < 0:
		return Interval{-b, -a}.Mod1(m).Negate()
	case a < 0:
		return Interval{a, -1}.Mod1(m).union(Interval{0, b}.Mod1(m))
	case b-a < abs(m) && a%m <= b%m:
		return Interval{a % m, b % m}
	default:
		return Interval{0, abs(m) - 1}
	}
}

func (i Interval) Mod2(i2 Interval) Interval {
	a := i.Min
	b := i.Max
	m := i2.Min
	n := i2.Max
	switch {
	case a > b || m > n:
		return Interval{0, -1}
	case b < 0:
		return Interval{-b, -a}.Mod2(Interval{m, n}).Negate()
	case a < 0:
		return Interval{a, -1}.Mod2(Interval{m, n}).union(Interval{0, b}.Mod2(Interval{m, n}))
	case m == n:
		return Interval{a, b}.Mod1(m)
	case n <= 0:
		return Interval{a, b}.Mod2(Interval{-n, -m})
	case m <= 0:
		return Interval{a, b}.Mod2(Interval{1, max(-m, n)})
	case b-a >= n:
		return Interval{0, n - 1}
	case b-a >= m:
		return Interval{0, b - a - 1}.union(Interval{a, b}.Mod2(Interval{b - a + 1, n}))
	case m > b:
		return Interval{a, b}
	case n > b:
		return Interval{0, b}
	default:
		return Interval{0, n - 1}
	}
}

type CubeMap struct {
	N    int
	grid [6]map[Position]uint8
}

type State3D struct {
	N    int
	face int
	pos  Position
	dir  int
}

type TransitionTable [6][4]struct{ face, rot int }

type Topology struct {
	row    map[int]Interval
	column map[int]Interval
	grid   map[Position]uint8
}

type State struct {
	pos  Position
	dir  int
	grid *Topology
}

func (s State) String() string {
	dir := []string{">", "v", "<", "^"}
	return fmt.Sprintf("%d,%d  %s", s.pos.y+1, s.pos.x+1, dir[s.dir])
}
