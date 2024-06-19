package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// WriteFile (filePath string, data io.Reader, fd io.Writer) error,
// которая принимает путь к файлу, данные для записи и дескриптор файла, и возвращает ошибку.
func WriteFile(filePath string, data io.Reader, fd io.Writer) error {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Errorf("ошибка при записи файла: %v", err.Error())
	}
	defer file.Close()
	scanerLine := bufio.NewReader(data)
	dataStr, err := scanerLine.ReadString('\n')
	if err != nil {
		fmt.Errorf("ошибка при записи файла: %v", err.Error())
	}
	fd = file
	fd.Write([]byte(dataStr))
	return nil
}

func main() {
	filePath := "course1/13.popular_package/7.package_os/rest_API_4.13.7.2/file.txt"

	err := WriteFile(filePath, strings.NewReader("Hello, World!"), os.Stdout)
	if err != nil {
		fmt.Errorf("ошибка при записи файла: %v", err.Error())
	}
}
