package main

import (
	"math/rand"
	"strings"
	"time"
)

func generateActivationKey() string {
	rand.Seed(time.Now().UnixNano())
	rendGen := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	tir := []rune("-")
	var builder strings.Builder
	for i := 0; i < 16; i++ {
		if i > 0 && (i)%4 == 0 {
			builder.WriteRune(tir[0])
		}
		builder.WriteRune(rendGen[rand.Intn(len(rendGen))])
	}
	return builder.String()
}

//func main() {
//	activationKey := generateActivationKey()
//	fmt.Println(activationKey) // UQNI-NYSI-ZVYB-ZEFQ
//}
