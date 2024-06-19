package main

import (
	"fmt"
)

type phoneReader string

//func (p phoneReader) Read(bs []byte) (int, error) {
//	count := 0
//	for i := 0; i < len(p); i++ {
//		if p[i] >= '0' && p[i] <= '9' {
//			bs[count] = p[i]
//			count++
//		}
//	}
//	return count, io.EOF
//}

func main() {
	//phone1 := phoneReader("+1(234)567 90-10")
	//io.Copy(os.Stdout, phone1)
	fmt.Print("expected output")
	//fmt.Println(revString("expected output"))
}

// рекурсивный разворот строки
//func revString(str string) string {
//	if len(str) <= 1 {
//		return str
//	}
//	return revString(str[1:]) + string(str[0])
//}
