package src

import (
	_ "embed"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func parse(input string) (topology *Topology, path string) {
	input = strings.TrimSuffix(input, "\n")
	parts := strings.Split(input, "\n\n")
	lines := strings.Split(parts[0], "\n")

	grid := make(map[Position]uint8)
	row := make(map[int]Interval)
	column := make(map[int]Interval)

	for j, line := range lines {
		ymin := math.MaxInt
		ymax := 0
		xmin := math.MaxInt
		xmax := 0
		for i, c := range line {
			if _, ok := column[i]; !ok {
				column[i] = Interval{math.MaxInt, 0}
			}
			if c == '#' || c == '.' {
				p := Position{i, j}
				grid[p] = uint8(c)
				xmin = min(xmin, i)
				xmax = max(xmax, i)
				ymin = min(column[i].Min, j)
				ymax = max(column[i].Max, j)
				column[i] = Interval{ymin, ymax}
			}
		}
		row[j] = Interval{xmin, xmax}
	}

	return &Topology{row, column, grid}, parts[1]
}

func step(state *State, order string) {
	switch order {
	case "L":
		state.dir = (state.dir + 3) % 4
	case "R":
		state.dir = (state.dir + 1) % 4
	default:
		nbSteps, _ := strconv.Atoi(order)
		for i := 0; i < nbSteps; i++ {
			X, Y := state.pos.x, state.pos.y
			var nextX, nextY int
			switch state.dir {
			case 0:
				nextX = X + 1
				nextY = Y
				if nextX > state.grid.row[Y].Max {
					nextX = state.grid.row[Y].Min
				}
			case 1:
				nextX = X
				nextY = Y + 1
				if nextY > state.grid.column[X].Max {
					nextY = state.grid.column[X].Min
				}
			case 2:
				nextX = X - 1
				nextY = Y
				if nextX < state.grid.row[Y].Min {
					nextX = state.grid.row[Y].Max
				}
			case 3:
				nextX = X
				nextY = Y - 1
				if nextY < state.grid.column[X].Min {
					nextY = state.grid.column[X].Max
				}
			}
			if c, ok := state.grid.grid[Position{nextX, nextY}]; ok && c == '.' {
				state.pos.x = nextX
				state.pos.y = nextY
			} else {
				break
			}
		}

	}
}

func Part1(input string) int {
	topology, path := parse(input)
	state := State{Position{topology.row[0].Min, 0}, 0, topology}
	path = strings.ReplaceAll(path, "L", " L ")
	path = strings.ReplaceAll(path, "R", " R ")
	orders := strings.Split(path, " ")
	for _, order := range orders {
		step(&state, order)
	}

	res := 1000*(state.pos.y+1) + 4*(state.pos.x+1) + state.dir
	return res
}

func (s State3D) String() string {
	dir := []string{">", "v", "<", "^"}
	return fmt.Sprintf("face %d -- %d,%d  %s", s.face+1, s.pos.y+1, s.pos.x+1, dir[s.dir])
}

func (s *State3D) rotate90() {
	s.dir = (s.dir + 1) % 4
	s.pos.x, s.pos.y = s.N-s.pos.y-1, s.pos.x
}

func (s *State3D) move(cube *CubeMap, transitionTable [6][4]struct{ face, rot int }) bool {
	start := *s
	switch s.dir {
	case 0:
		s.pos.x++
	case 1:
		s.pos.y++
	case 2:
		s.pos.x--
	case 3:
		s.pos.y--
	}
	if s.pos.x >= 0 && s.pos.x < s.N && s.pos.y >= 0 && s.pos.y < s.N {
		if cube.grid[s.face][s.pos] == '#' {
			*s = start
			return false
		}
		return true
	}

	if s.pos.x < 0 {
		s.pos.x = s.N - 1
	}
	if s.pos.x >= s.N {
		s.pos.x = 0
	}
	if s.pos.y < 0 {
		s.pos.y = s.N - 1
	}
	if s.pos.y >= s.N {
		s.pos.y = 0
	}
	s.face = transitionTable[start.face][start.dir].face
	switch transitionTable[start.face][start.dir].rot {
	case 90:
		s.rotate90()
	case 180:
		s.rotate90()
		s.rotate90()
	case 270:
		s.rotate90()
		s.rotate90()
		s.rotate90()
	}
	if cube.grid[s.face][s.pos] == '#' {
		*s = start
		return false
	}
	return true
}

func step3D(s *State3D, cube *CubeMap, order string, transitionTable TransitionTable) {
	switch order {
	case "L":
		s.dir = (s.dir + 3) % 4
	case "R":
		s.dir = (s.dir + 1) % 4
	default:
		nbSteps, _ := strconv.Atoi(order)
		for i := 0; i < nbSteps; i++ {
			ok := s.move(cube, transitionTable)
			if !ok {
				break
			}
		}

	}
}

func Part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	parts := strings.Split(input, "\n\n")
	lines := strings.Split(parts[0], "\n")
	path := parts[1]
	N := 50
	faces := [6]struct {
		X Interval
		Y Interval
	}{
		{X: Interval{1 * N, 2*N - 1}, Y: Interval{0 * N, 1*N - 1}},
		{X: Interval{2 * N, 3*N - 1}, Y: Interval{0 * N, 1*N - 1}},
		{X: Interval{1 * N, 2*N - 1}, Y: Interval{1 * N, 2*N - 1}},
		{X: Interval{0 * N, 1*N - 1}, Y: Interval{2 * N, 3*N - 1}},
		{X: Interval{1 * N, 2*N - 1}, Y: Interval{2 * N, 3*N - 1}},
		{X: Interval{0 * N, 1*N - 1}, Y: Interval{3 * N, 4*N - 1}},
	}

	transition := TransitionTable{
		/* face 1 */ {{face: 2 - 1, rot: 0}, {face: 3 - 1, rot: 0}, {face: 4 - 1, rot: 180}, {face: 6 - 1, rot: 90}},
		/* face 2 */ {{face: 5 - 1, rot: 180}, {face: 3 - 1, rot: 90}, {face: 1 - 1, rot: 0}, {face: 6 - 1, rot: 0}},
		/* face 3 */ {{face: 2 - 1, rot: 270}, {face: 5 - 1, rot: 0}, {face: 4 - 1, rot: 270}, {face: 1 - 1, rot: 0}},
		/* face 4 */ {{face: 5 - 1, rot: 0}, {face: 6 - 1, rot: 0}, {face: 1 - 1, rot: 180}, {face: 3 - 1, rot: 90}},
		/* face 5 */ {{face: 2 - 1, rot: 180}, {face: 6 - 1, rot: 90}, {face: 4 - 1, rot: 0}, {face: 3 - 1, rot: 0}},
		/* face 6 */ {{face: 5 - 1, rot: 270}, {face: 2 - 1, rot: 0}, {face: 1 - 1, rot: 270}, {face: 4 - 1, rot: 0}},
	}

	cube := CubeMap{N: N}
	for i := 0; i < 6; i++ {
		cube.grid[i] = make(map[Position]uint8)
	}
	for j, line := range lines {
		for i, c := range line {
			for k := 0; k < 6; k++ {
				if faces[k].X.Contains(i) && faces[k].Y.Contains(j) {
					cube.grid[k][Position{i % N, j % N}] = uint8(c)
				}
			}
		}
	}
	path = strings.ReplaceAll(path, "L", " L ")
	path = strings.ReplaceAll(path, "R", " R ")
	orders := strings.Split(path, " ")

	state := State3D{N: N, face: 0, pos: Position{0, 0}, dir: 0}
	for _, order := range orders {
		step3D(&state, &cube, order, transition)
	}

	return 1000*(faces[state.face].Y.Min+state.pos.y+1) + 4*(faces[state.face].X.Min+state.pos.x+1) + state.dir

}

func Day22() {
	content, _ := os.ReadFile("input/day22.txt")
	day22Content := string(content)
	fmt.Println(Part1(day22Content))
	fmt.Println(Part2(day22Content))
}
