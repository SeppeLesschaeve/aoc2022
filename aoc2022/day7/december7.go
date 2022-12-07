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
		entity := *entityStar
		if entity.Type == "Dir" {
			sizes = append(sizes, getSizes(entity)...)
		}
	}
	return sizes
}

func constructTree(lines []string) Entity {
	var tree Entity
	var curr Entity
	var indices []int
	var currDirs []Entity
	for i := 0; i < len(lines); {
		if lines[i] == "$ cd .." {
			currDirs = currDirs[1:]
			if len(currDirs) != 0 {
				curr = currDirs[0]
			}
			newIndex := getIndexOfFirstDirAfterIndex(indices[len(indices)-1], curr)
			if newIndex == -1 {
				indices = indices[:len(indices)-1]
			} else {
				indices[len(indices)-1] = newIndex
			}
			i++
		} else if lines[i][:4] == "$ cd" {
			curr = Entity{strings.Split(lines[i], " ")[2], "Dir", 0, []*Entity{}}
			currDirs = append([]Entity{curr}, currDirs...)
			if indices == nil {
				tree = curr
			}
			i++
		} else if lines[i][:4] == "$ ls" {
			entities := getEntities(lines[i+1:])
			curr.extend(entities, nil)
			currDirs[0] = curr
			tree.extend(curr.subEntities, indices)
			index := getIndexOfFirstDirAfterIndex(-1, curr)
			if index != -1 {
				indices = append(indices, index)
			}
			i += len(entities) + 1
		}
	}
	return tree
}

func getIndexOfFirstDirAfterIndex(index int, curr Entity) int {
	for i, entityStar := range curr.subEntities {
		entity := *entityStar
		if entity.Type == "Dir" && i > index {
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
			var entity Entity
			if string(entitySplit[0]) == "dir" {
				entity = Entity{name: entitySplit[1], Type: "Dir", Size: 0, subEntities: []*Entity{}}
			} else {
				size, _ := strconv.Atoi(entitySplit[0])
				entity = Entity{name: entitySplit[1], Type: "File", Size: size, subEntities: nil}
			}
			entities = append(entities, &entity)
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
