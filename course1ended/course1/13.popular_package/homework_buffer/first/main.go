package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	//В этом примере мы создадим новый буфер с помощью пакета bytes и запишем в него данные.
	buf := bytes.NewBuffer([]byte("some string")) // create buffer with data
	fmt.Println(buf.String())

	//В этом примере мы объединим два среза байтов с помощью пакета bytes.
	buf2 := []byte(" second note!")
	bufer := bytes.NewBuffer(buf.Bytes()) // create buffer with data from another buffer
	bufer.Write(buf2)                     // write to second buffer
	fmt.Println(bufer.String())

	//В этом примере мы разделим срез байтов на две части с помощью пакета bytes.
	cut := make([]byte, 12)
	numByte, err := bufer.Read(cut) //cut of 12 bytes from buffer to variable cut
	if err != nil {
		fmt.Errorf("error = %v", err.Error())
	}
	fmt.Printf("%d bites in new []byte = %v, left in bufer = %v\n", numByte, string(cut), bufer.String())

	//В этом примере мы добавим данные в существующий срез байтов с помощью пакета bytes.
	message := make([]byte, 11)
	bufer.Read(message)                  //cut of 11 bytes from buffer to variable cut
	bufer.Reset()                        // dell all data from bufer
	bufer.Write(message)                 // add cut []byte to buffer
	bufer.WriteString(" + another one)") // add new string to buffer
	fmt.Println(bufer.String())

	//В этом примере мы найдем позицию первого вхождения подстроки в срез байтов с помощью пакета bytes.
	plusFind := []byte("+")
	indexPlus := bytes.Index(bufer.Bytes(), plusFind) // find index of string "+" in buffer
	cut = make([]byte, indexPlus+2)
	bufer.Read(cut) // cut all text before string "+" and next 2 bites from buffer
	fmt.Println(string(cut), "- is cut, and what's left - ", bufer.String())

	//В этом примере мы будем буферизованно считывать данные из источника ввода с помощью пакета bufio
	//reader := bufio.NewReader(os.Stdin) // create reader from console input
	//fmt.Print("Input your text here: ")
	//text, _ := reader.ReadString('\n')                      // save input text in variable text
	text := " - 123 - "
	bufer.WriteString(text[:len(text)-1] + "end str test") //удалить последний символ строки - text[:len(text)-1]
	fmt.Println(bufer.String())

	//В этом примере мы будем буферизованно записывать данные в целевой источник вывода с помощью пакета bufio
	file, err := os.Create("course1/13.popular_package/homework_buffer/first/file.txt") // create file
	if err != nil {
		fmt.Errorf("error %v", err.Error())
		os.Exit(1)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)                                                  // create writer to the file
	writer.WriteString(bufer.String() + "\n")                                        // writing to the file
	writer.Flush()                                                                   // save data in file
	file, err = os.Open("course1/13.popular_package/homework_buffer/first/file.txt") // create file
	if err != nil {
		fmt.Errorf("unable to open file %v", err.Error())
		os.Exit(1)
	}
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Errorf("error %v", err.Error())
				os.Exit(1)
			}
		}
		fmt.Println(line)
	}
	input := "Hello, world!"
	reader = bufio.NewReader(strings.NewReader(input)) // буферизированное чтение из источника input
	line, _ := reader.ReadString('d')                  // читает до символа d
	line2, _ := reader.ReadString(1)                   //читает 1 байт
	//fmt.Println(line, "\n", line2, "\n", bufer.String())
	scanerLine := bufio.NewScanner(strings.NewReader(line + "\n" + line2 + "\n" + bufer.String()))
	scanerWord := bufio.NewScanner(strings.NewReader(line + "\n" + line2 + "\n" + bufer.String()))
	scanerLine.Split(bufio.ScanLines)
	scanerWord.Split(bufio.ScanWords)
	var textLines = make([]string, 0, 10)
	var textWords = make([]string, 0, 10)
	for scanerLine.Scan() {
		textLines = append(textLines, scanerLine.Text())
	}
	for scanerWord.Scan() {
		textWords = append(textWords, scanerWord.Text())
	}
	rangeLists := [][]string{textLines, textWords}
	for i, l := range rangeLists {
		for j, val := range l {
			fmt.Printf("list %d val %d = '%s'; ", i+1, j+1, val)
		}
		fmt.Println()
	}

}
