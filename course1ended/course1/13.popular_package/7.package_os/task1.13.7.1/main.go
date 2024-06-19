package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

//Необходимо реализовать функцию WriteFile, которая будет записывать данные в файл с указанным путем и правами доступа.

func WriteFile(filePath string, data []byte, perm os.FileMode) error {
	re, _ := regexp.Compile(`(/[a-zA-Z]+.txt)`)
	createPath := re.ReplaceAllString(filePath, "")
	err := os.MkdirAll(createPath, 0777)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, perm)
	if err != nil {
		fmt.Println("+++", filePath)
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file) // create writer to the file
	writer.Write(data)              // writing to the file
	writer.Flush()
	return nil
}

func main() {
	this_dir, err := os.Getwd()
	if err != nil {
		fmt.Errorf("error: %v", err.Error())
	}
	err = WriteFile(this_dir+"/course1/13.popular_package/7.package_os/rest_API_4.13.7.1/newdirfile/file.txt", []byte("Hello, World!"), os.FileMode(0644))
	if err != nil {
		fmt.Errorf("error: %v", err.Error())
	}
}
