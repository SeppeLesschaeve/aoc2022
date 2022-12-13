package src

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func (file *File) GetSize() int {
	return file.Size
}

func (dir *Directory) GetSize() int {
	sum := 0
	for _, entity := range dir.subEntities {
		sum += (*entity).GetSize()
	}
	return sum
}

func Day07() {
	content, _ := os.ReadFile("input/day07.txt")
	day7Content := string(content)
	lines := strings.Split(day7Content, "\n")
	tree := constructTree(lines)
	sizes := tree.getSizes()
	firstPartSum := 0
	for _, size := range sizes {
		if size < 100000 {
			firstPartSum += size
		}
	}
	secondPartSize := sizes[0]
	for _, size := range sizes {
		if size >= sizes[0]-40000000 && size <= secondPartSize {
			secondPartSize = size
		}
	}
	fmt.Println(firstPartSum)
	fmt.Println(secondPartSize)
}

func (dir *Directory) getSizes() []int {
	var sizes []int
	sizes = append(sizes, dir.GetSize())
	for _, entityStar := range dir.subEntities {
		if reflect.TypeOf(*entityStar).String() == "*src.Directory" {
			sizes = append(sizes, (*entityStar).(*Directory).getSizes()...)
		}
	}
	return sizes
}

func constructTree(lines []string) Directory {
	currDirs := []*Directory{{strings.Split(lines[0], " ")[2], []*Entity{}}}
	tree := *currDirs[0]
	var indices []int
	for i := 1; i < len(lines); {
		if lines[i] == "$ cd .." {
			currDirs = currDirs[1:]
			newIndex := currDirs[0].getIndexOfFirstDirAfterIndex(indices[len(indices)-1])
			if newIndex == -1 {
				indices = indices[:len(indices)-1]
			} else {
				indices[len(indices)-1] = newIndex
			}
			i++
		} else if lines[i][:4] == "$ cd" {
			currDirs = append([]*Directory{{strings.Split(lines[i], " ")[2], []*Entity{}}}, currDirs...)
			i++
		} else if lines[i][:4] == "$ ls" {
			currDirs[0].extend(getEntities(lines[i+1:]), nil)
			tree.extend(currDirs[0].subEntities, indices)
			index := currDirs[0].getIndexOfFirstDirAfterIndex(-1)
			if index != -1 {
				indices = append(indices, index)
			}
			i += len(currDirs[0].subEntities) + 1
		}
	}
	return tree
}

func (dir *Directory) getIndexOfFirstDirAfterIndex(index int) int {
	for i, entityStar := range dir.subEntities {
		if reflect.TypeOf(*entityStar).String() == "*src.Directory" && i > index {
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
		}
		entitySplit := strings.Split(line, " ")
		var entity Entity
		if string(entitySplit[0]) == "dir" {
			entity = (Entity)(&Directory{name: entitySplit[1], subEntities: []*Entity{}})
		} else {
			size, _ := strconv.Atoi(entitySplit[0])
			entity = (Entity)(&File{name: entitySplit[1], Size: size})
		}
		entities = append(entities, &entity)
	}
	return entities
}

func (dir *Directory) extend(entities []*Entity, indices []int) {
	if len(indices) == 0 {
		dir.subEntities = entities
	} else {
		(*dir.subEntities[indices[0]]).(*Directory).extend(entities, indices[1:])
	}
}
