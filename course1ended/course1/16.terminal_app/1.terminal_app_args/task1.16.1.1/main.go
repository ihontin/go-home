package main

import (
	"fmt"
	"os"
)

func getArgs() []string {
	return os.Args[1:]
}
func main() {
	fmt.Println(getArgs())
}
