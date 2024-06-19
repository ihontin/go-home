package main

func mutate(a *int) {
	*a = 42
}

func ReverseString(str *string) {
	var sideStr string
	for _, letter := range *str {
		sideStr = string(letter) + sideStr
	}
	*str = sideStr
}

func main() {
	a := 51
	s := "rakatau"
	mutate(&a)
	ReverseString(&s)

}
