package main

import (
	"fmt"
	"os/exec"
)

// ExecBin которая принимает путь к бинарному файлу binPath и возвращает вывод программы в виде строки.
func ExecBin(binPath string, args ...string) string {
	cmd := exec.Command(binPath, args...)
	output, err := cmd.Output()
	if err != nil {
		return fmt.Sprintf("Error executing binary: %v", err)
	}
	return string(output)
}

//func main() {
//	fmt.Println(ExecBin("ls", "-la"))
//	fmt.Println(ExecBin("nonexistent-binary"))
//}
