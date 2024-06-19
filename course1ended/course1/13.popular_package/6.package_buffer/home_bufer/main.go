package main

import (
	"bufio"
	"fmt"
	"strings"
)

//В этом примере мы создадим новый буфер с помощью пакета bytes и запишем в него данные.

//func main() {
//	buffer := bytes.NewBuffer([]byte("Hello, world!")) // хранит инфу в буффере в виде среза байт
//	fmt.Println(buffer.String())
//	m := bytes.Buffer{}                       // создаем новый буфер m
//	n, _ := m.Write([]byte("world Hello, !")) // записываем данные в буфер m
//	fmt.Println("данные в буфере:", m.String())
//	var sliseByte = make([]byte, n) // создать слайс заполненный
//	_, _ = m.Read(sliseByte)        // читаем буфер в слайс
//	fmt.Println(string(sliseByte))
//	fmt.Println("буфер буст:", m.String()) // после прочтения буфер пуст
//}

//Пример 2

//В этом примере мы объединим два среза байтов с помощью пакета bytes.
//func main() {
//	slice1 := []byte("Hello, ")
//	slice2 := []byte("world!")
//	buffer := bytes.NewBuffer(slice1) // создать буфер с данными из первого среза
//	buffer.Write(slice2) // добавить в буфер данные из второго среза
//	fmt.Println(buffer.String())
//}

//Пример 3

// В этом примере мы разделим срез байтов на две части с помощью пакета bytes.
//func main() {
//	slice := []byte("Hello, world!")
//	buffer := bytes.NewBuffer(slice)
//	part1 := make([]byte, 5)
//	buffer.Read(part1) // заберет часть данных из буфера запишет в срез part1
//	part2 := buffer.Bytes()
//	fmt.Println(string(part1))
//	fmt.Println(string(part2))
//}

//Пример 4

//В этом примере мы добавим данные в существующий срез байтов с помощью пакета bytes.
//func main() {
//	slice := []byte("Hello, ")
//	buffer := bytes.NewBuffer(slice) // создаст буфер с данными
//	buffer.WriteString("world!") // добавит данные в буфер
//	fmt.Println(buffer.String())
//}

//Пример 5

//В этом примере мы найдем позицию первого вхождения подстроки в срез байтов с помощью пакета bytes.
//func main() {
//	slice := []byte("Hello, world!")
//	buffer := bytes.NewBuffer(slice)
//	index := bytes.Index(buffer.Bytes(), []byte("world")) // Найдет индекс вхождения подстроки
//	fmt.Println(index)
//}

//Пример 6

//В этом примере мы будем буферизованно считывать данные из источника ввода с помощью пакета bufio.
//func main() {
//	reader := bufio.NewReader(os.Stdin) // буферизированный ввод
//	fmt.Print("Введите текст: ")
//	text, _ := reader.ReadString('\n') // чтение из буфера в переменную по разделителю окончания строки
//	fmt.Println(text)
//}

//Пример 7

//В этом примере мы будем буферизованно записывать данные в целевой источник вывода с помощью пакета bufio.
//func main() {
//	file, _ := os.Create("output.txt")
//	writer := bufio.NewWriter(file) // создаст буферизированный поток для записи в файл
//	writer.WriteString("Hello, world!") // запись данных в буферизированный поток
//	writer.Flush() запись данных в файл
//	fmt.Println("Данные записаны в файл.")
//}

//Пример 8

// В этом примере мы будем буферизованно считывать строки из буфера с указанным префиксом и суффиксом с помощью пакета bufio.
func main() {
	input := "Hello, world!"
	reader := bufio.NewReader(strings.NewReader(input)) // буферизованный поток для чтения данных из строки
	line, _ := reader.ReadString(',')
	fmt.Println(line)
}
