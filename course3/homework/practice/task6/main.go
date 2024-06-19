package main

import (
	"fmt"
)

func main() {
	group := "ratosum"
	f := fmt.Sprintf("Group %s Привет, мир", string(group[len(group)-1]))
	fmt.Println(f)
}
