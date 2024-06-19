package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func printTree(path string, prefix string, isLast bool, depth int) {
	entries, err := os.ReadDir(path) //чтение содержимого директории
	if err != nil {
		fmt.Println(err)
		return
	}
	//Перебераем содержимое в цикле
	for i, entry := range entries {
		entryName := entry.Name()
		entryPath := filepath.Join(path, entryName)

		// Определить префикс текущей записи
		var entryPrefix string
		if isLast {
			entryPrefix = prefix + "└── "
		} else {
			entryPrefix = prefix + "├── "
		}

		// Распечатать запись
		fmt.Println(entryPrefix + entryName)

		// Рекурсивный вызов катала, если entry является таковым
		if entry.IsDir() && depth > 0 {
			var newPrefix string
			if isLast {
				newPrefix = prefix + "    "
			} else {
				newPrefix = prefix + "│   "
			}
			printTree(entryPath, newPrefix, i == len(entries)-1, depth-1)
		}
	}
}

func main() {
	var path string
	var depth int
	flag.IntVar(&depth, "depth", 0, "tree depth")
	flag.Parse()

	if !strings.HasPrefix(path, "/") {
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		path = filepath.Join(wd, path)
	}

	printTree(path, "", true, depth)
}
