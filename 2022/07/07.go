package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	content, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	if err := logic(content); err != nil {
		fmt.Printf("%v\n", err)
	}
}

func logic(input []byte) error {
	tree, err := parseTree(input)
	fmt.Println(tree)
	if err != nil {
		return err
	}
	size := tree.calculateSize()
	fmt.Printf("Part 1: %d\n", size)
	return nil
}

type File struct {
	Name string
	Size int
}

type Directory struct {
	Name   string
	Childs []*Directory
	Parent *Directory
	Files  []*File
}

func parseTree(input []byte) (*Directory, error) {
	scanner := bufio.NewScanner(bytes.NewReader(input))

	rootDir := Directory{
		Name: "/",
	}

	currentDir := &rootDir

	for scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(s, "$") {
			// command
			// remove leading $ to get plain command
			s = strings.TrimLeft(s, "$ ")
			command := strings.SplitN(s, " ", 2)
			// fmt.Printf("Got command: %s\n", command)
			switch command[0] {
			case "ls":
				// nothing to do with ls
				continue
			case "cd":
				targetDir := command[1]
				if targetDir == "/" {
					continue
				}
				if targetDir == ".." {
					currentDir = currentDir.Parent
					continue
				}
				foundDir := false
				for _, child := range currentDir.Childs {
					if child.Name == targetDir {
						if child.Parent == nil {
							child.Parent = currentDir
						}
						currentDir = child
						foundDir = true
						// fmt.Printf("switching current dir to already listed %s\n", targetDir)
						break
					}
				}
				if !foundDir {
					newChild := Directory{
						Name:   targetDir,
						Parent: currentDir,
					}
					// fmt.Printf("switching current dir to new %s\n", targetDir)
					currentDir = &newChild
				}
			default:
				panic(fmt.Sprintf("unknown command %s", command[0]))
			}
		} else {
			// command output
			output := strings.SplitN(s, " ", 2)
			// fmt.Printf("Got output: %s\n", output)
			switch output[0] {
			case "dir":
				// dir in ls output
				dir := Directory{
					Name:   output[1],
					Parent: currentDir,
				}
				currentDir.Childs = append(currentDir.Childs, &dir)
			default:
				// file
				size, err := strconv.Atoi(output[0])
				if err != nil {
					return nil, fmt.Errorf("invalid size %v: %v", output[0], err)
				}
				file := File{
					Name: output[1],
					Size: size,
				}
				currentDir.Files = append(currentDir.Files, &file)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &rootDir, nil
}

func (d *Directory) String() string {
	return d.printContents(0)
}

func (d *Directory) printContents(indent int) string {
	var sb strings.Builder
	pre := strings.Repeat("\t", indent)
	_, err := sb.WriteString(strings.TrimLeft(fmt.Sprintf("%s - %s (dir)\n", pre, d.Name), " "))
	if err != nil {
		panic(err)
	}
	for _, child := range d.Childs {
		_, err := sb.WriteString(child.printContents(indent + 1))
		if err != nil {
			panic(err)
		}
	}
	pre = strings.Repeat("\t", indent+1)
	for _, file := range d.Files {
		_, err := sb.WriteString(fmt.Sprintf("%s - %s\t\t%d\n", pre, file.Name, file.Size))
		if err != nil {
			panic(err)
		}
	}
	return sb.String()
}

func (d *Directory) calculateSize() int {
	size := 0
	for _, f := range d.Files {
		size += f.Size
	}
	for _, d := range d.Childs {
		size += d.calculateSize()
	}
	return size
}
