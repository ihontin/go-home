package main

import (
	"fmt"
	"github.com/mattevans/dinero"
	"time"
)

func currencyPairRate(a, b string, val float64) float64 {
	client := dinero.NewClient(
		"8990687eff3648c5b53612f49fc5782a",
		a,
		20*time.Minute,
	)
	rsp, err := client.Rates.Get(b)
	if err != nil {
		fmt.Println("Error:", err)
		return 0.0
	}
	return val * *rsp
}

func main() {
	// Пример использования функции

	rate := currencyPairRate("USD", "EUR", 100.0)
	fmt.Println(rate) // 82.73

}
