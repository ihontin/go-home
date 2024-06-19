package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// В этом примере мы создадим новый файл с указанным именем с помощью функции os.Create.
func main() {
	path, _ := os.Getwd()
	path += "/course1/13.popular_package/homework_os/first/"
	fmt.Println(path)
	file, err := os.Create(path + "example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() // отложенное закрытие файла
	fmt.Println("Файл успешно создан.")

	//В этом примере мы откроем существующий файл для чтения с помощью функции os.Open.
	file, err = os.Open(path + "example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() // отложенное закрытие файла
	fmt.Println("Файл успешно открыт.")

	//В этом примере мы откроем файл с указанными флагами доступа с помощью функции os.OpenFile.
	file, err = os.OpenFile("example.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() // отложенное закрытие файла
	fmt.Println("Файл успешно открыт.")

	//В этом примере мы создадим новую директорию с указанным именем с помощью функции os.Mkdir.
	err = os.Mkdir("example_dir", 0755)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Директория успешно создана.")

	//В этом примере мы создадим директорию и все промежуточные директории с помощью функции os.MkdirAll.
	err = os.MkdirAll("path/to/example_dir", 0755)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Директория успешно создана.")

	//В этом примере мы удалим файл или директорию с помощью функции os.Remove.
	err = os.Remove("example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Файл успешно удален.")

	//В этом примере мы удалим директорию и все ее содержимое с помощью функции os.RemoveAll.
	err = os.RemoveAll("example_dir")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Директория успешно удалена.")

	//В этом примере мы получим текущую рабочую директорию с помощью функции os.Getwd.
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Текущая рабочая директория:", wd)

	//Рекурсивный вывод файлов
	dir := "." // Здесь можно указать путь к нужной директории
	err = listFiles(dir, 0)
	if err != nil {
		fmt.Printf("Ошибка: %v", err)
	}
}

func listFiles(dir string, level int) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		fileInfo, err := os.Stat(filepath.Join(dir, file.Name()))
		if err != nil {
			return err
		}

		indent := ""
		for i := 0; i < level; i++ {
			indent += " "
		}

		fmt.Printf("%s%s\n", indent, file.Name())

		if fileInfo.IsDir() {
			subdir := filepath.Join(dir, file.Name())
			listFiles(subdir, level+1)
		}
	}
	return nil
}
