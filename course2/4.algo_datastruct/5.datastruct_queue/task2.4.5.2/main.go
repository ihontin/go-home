package main

import "fmt"

type BrowserHistory struct {
	stack []string
}

func (h *BrowserHistory) Visit(url string) {
	h.stack = append(h.stack, url)
	fmt.Printf("Посещение [%s]\n", url)
}

func (h *BrowserHistory) Back() {
	if len(h.stack) != 0 {
		url := h.stack[len(h.stack)-1]
		h.stack = h.stack[:len(h.stack)-1]
		fmt.Printf("Возврат к [%s]\n", url)
	}
	fmt.Println("Нет больше истории для возврата")

}

func (h *BrowserHistory) PrintHistory() {
	fmt.Println("История браузера:")
	for i := 0; i < len(h.stack); i++ {
		fmt.Println(h.stack[len(h.stack)-1-i])
	}
}

func main() {
	history := &BrowserHistory{}
	history.Visit("www.google.com")
	history.Visit("www.github.com")
	history.Visit("www.openai.com")
	history.Back()
	history.PrintHistory()
}
