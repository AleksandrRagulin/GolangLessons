package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

var level = 0

var isLast []bool

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "[-f]"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, path string, files bool) error {
	dirs, err := os.ReadDir(path)

	var resDirs []os.DirEntry
	isLastDir := false

	if err != nil {
		return nil
	}
	for _, dir := range dirs {
		if dir.IsDir() || files {
			resDirs = append(resDirs, dir)
		}
	}

	sort.Slice(resDirs, func(i, j int) bool {
		return resDirs[i].Name() < resDirs[j].Name()
	})

	for idx, dir := range resDirs {

		tabs := ""
		for i := 0; i < level; i++ {
			if isLast[i] {
				tabs += "\t"
			} else {
				tabs += "│\t"
			}
		}

		if idx == len(resDirs)-1 {
			tabs += "└───"
			isLastDir = true

		} else {
			tabs += "├───"
		}

		_, _ = fmt.Fprintf(out, tabs)

		if dir.IsDir() {
			_, _ = fmt.Fprintf(out, "%s\n", dir.Name())
			level++
			isLast = append(isLast, isLastDir)
			newPath := path + "/" + dir.Name()
			err := dirTree(out, newPath, files)
			if err != nil {
				fmt.Println("err")
			}
		} else {
			info, _ := dir.Info()
			if info.Size() > 0 {
				_, _ = fmt.Fprintf(out, "%s (%vb)\n", dir.Name(), info.Size())
			} else {
				_, _ = fmt.Fprintf(out, "%s (empty)\n", dir.Name())
			}
		}
	}

	if level > 0 {
		level--
	}

	lenIsLast := len(isLast)
	if lenIsLast > 0 {
		isLast = isLast[:lenIsLast-1]
	}
	return nil
}
