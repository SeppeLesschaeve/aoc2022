package day7

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Entity struct {
	name        string
	Type        string
	Size        int
	subEntities []*Entity
}

func (e *Entity) GetSize() int {
	if e.Type == "File" {
		return e.Size
	} else if e.Type == "Dir" {
		sum := 0
		for _, entity := range e.subEntities {
			sum += entity.GetSize()
		}
		return sum
	}
	return 0
}

func Day7() {
	content, _ := os.ReadFile("day7.txt")
	day7Content := string(content)
	lines := strings.Split(day7Content, "\n")
	tree := constructTree(lines)
	sizes := getSizes(tree)
	firstPartSum := 0
	for _, size := range sizes {
		if size < 100000 {
			firstPartSum += size
		}
	}
	sizeToDelete := sizes[0] - 40000000
	secondPartSize := sizes[0]
	for _, size := range sizes {
		if size >= sizeToDelete && size <= secondPartSize {
			secondPartSize = size
		}
	}
	fmt.Println(firstPartSum)
	fmt.Println(secondPartSize)
}

func getSizes(dir Entity) []int {
	var sizes []int
	if dir.Type == "Dir" {
		sizes = append(sizes, dir.GetSize())
	}
	for _, entityStar := range dir.subEntities {
		if (*entityStar).Type == "Dir" {
			sizes = append(sizes, getSizes(*entityStar)...)
		}
	}
	return sizes
}

func constructTree(lines []string) Entity {
	currDirs := []*Entity{{strings.Split(lines[0], " ")[2], "Dir", 0, []*Entity{}}}
	tree := *currDirs[0]
	var indices []int
	for i := 1; i < len(lines); {
		if lines[i] == "$ cd .." {
			currDirs = currDirs[1:]
			newIndex := getIndexOfFirstDirAfterIndex(indices[len(indices)-1], *currDirs[0])
			if newIndex == -1 {
				indices = indices[:len(indices)-1]
			} else {
				indices[len(indices)-1] = newIndex
			}
			i++
		} else if lines[i][:4] == "$ cd" {
			currDirs = append([]*Entity{{strings.Split(lines[i], " ")[2], "Dir", 0, []*Entity{}}}, currDirs...)
			i++
		} else if lines[i][:4] == "$ ls" {
			currDirs[0].extend(getEntities(lines[i+1:]), nil)
			tree.extend(currDirs[0].subEntities, indices)
			index := getIndexOfFirstDirAfterIndex(-1, *currDirs[0])
			if index != -1 {
				indices = append(indices, index)
			}
			i += len(currDirs[0].subEntities) + 1
		}
	}
	return tree
}

func getIndexOfFirstDirAfterIndex(index int, curr Entity) int {
	for i, entityStar := range curr.subEntities {
		if (*entityStar).Type == "Dir" && i > index {
			return i
		}
	}
	return -1
}

func getEntities(lines []string) []*Entity {
	var entities []*Entity
	for _, line := range lines {
		if string(line[0]) == "$" {
			break
		} else {
			entitySplit := strings.Split(line, " ")
			if string(entitySplit[0]) == "dir" {
				entities = append(entities, &Entity{name: entitySplit[1], Type: "Dir", Size: 0, subEntities: []*Entity{}})
			} else {
				size, _ := strconv.Atoi(entitySplit[0])
				entities = append(entities, &Entity{name: entitySplit[1], Type: "File", Size: size, subEntities: nil})
			}
		}
	}
	return entities
}

func (e *Entity) extend(entities []*Entity, indices []int) {
	if len(indices) == 0 {
		e.subEntities = entities
	} else {
		(e.subEntities[indices[0]]).extend(entities, indices[1:])
	}
}
