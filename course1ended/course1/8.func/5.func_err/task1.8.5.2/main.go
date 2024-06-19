package main

import (
	"errors"
	"fmt"
)

// CheckDiscount (price, discount float64) (float64, error),
// которая будет проверять скидку на цену товара.
// Функция должна принимать два аргумента: цену товара и скидку в процентах. Если скидка не превышает 50%,
// функция должна вернуть цену товара с учетом скидки.
//
//	Если же скидка больше 50%, функция должна вернуть ошибку с сообщением “Скидка не может превышать 50%”.
func CheckDiscount(price, discount float64) (float64, error) {
	if discount > 50.0 || discount < 0 {
		err := errors.New("Скидка не может превышать 50%")
		return 0, err
	}
	discountMoney := (price / 100) * discount
	return price - discountMoney, nil
}

func main() {
	a, err := CheckDiscount(80.0, 41)
	checkErr(err)
	b, err := CheckDiscount(100.0, 0)
	checkErr(err)
	c, err := CheckDiscount(8000.50, 10)
	checkErr(err)
	fmt.Println(a, b, c)
}
func checkErr(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
