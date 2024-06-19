package main

import (
	"bufio"
	"os"
)

// ReadString принимает путь к файлу в качестве аргумента и возвращает содержимое файла в виде строки.
func ReadString(filePath string) string {
	var outString string           //creating a string to be returned
	file, err := os.Open(filePath) // open file

	// error handling
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()               // отложенное закрытие файла
	scaner := bufio.NewScanner(file) // создаем буфер в который положим содержимое файла
	for scaner.Scan() {
		//итерируемся построчно сохраняя текст в переменную
		outString += scaner.Text()
	}
	//check scaner for errors
	if err = scaner.Err(); err != nil {
		panic(err.Error())
	}
	return outString // out text from file
}

//func main() {
//	pathFile := "course1/13.popular_package/7.package_os/rest_API_4.13.7.3/file.txt"
//	fmt.Println(ReadString(pathFile))
//}
