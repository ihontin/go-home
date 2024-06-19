package main

import "fmt"

//Необходимо написать функцию для генерации строки с математическими операциями с использованием fmt.Sprintf().
//Функция должна принимать входные параметры и возвращать сформированную строку с операциями.

func generateMathString(operands []int, operator string) string {
	if len(operands) == 0 || operator != "+" {
		return ""
	}
	var sum int
	var outStr string
	for i, oper := range operands {
		sum += oper
		outStr += fmt.Sprintf("%d ", oper)
		if len(operands) != i+1 {
			outStr += operator + " "
		}
	}
	return fmt.Sprintf("%s= %d", outStr, sum)
}

//func main() {
//	fmt.Println(generateMathString([]int{2, 4, 6}, "+")) // "2 + 4 + 6 = 12"
//}
