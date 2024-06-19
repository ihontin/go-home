package main

import "fmt"

//Необходимо создать интерфейс для телевизоров, который будет иметь метод switchOFF,
//позволяющий выключить телевизор. У вас есть две модели телевизоров - Samsunger и LGer.

// Каждая модель имеет свои дополнительные методы - SamsungHub() для Samsunger и LGHub() для LGer.
// Обе модели также имеют методы GetStatus(), возвращающий состояние телевизора (true - включен, false - выключен),
// и GetModel(), возвращающий название модели телевизора.

type TVer interface {
	switchOFF()
	GetStatus()
	GetModel()
}

type SamsungTV struct {
	status bool
	model  string
}

func (tv *SamsungTV) SamsungHub() string {
	return "SamsungHub"
}

func (tv *SamsungTV) switchOn() bool {
	tv.status = true
	return true
}
func (tv *SamsungTV) switchOFF() bool {
	tv.status = false
	return true
}
func (tv *SamsungTV) GetStatus() bool {
	return tv.status
}
func (tv *SamsungTV) GetModel() string {
	return tv.model
}

type LgTV struct {
	status bool
	model  string
}

func (tv *LgTV) LGHub() string {
	return "SamsungHub"
}

func (tv *LgTV) switchOn() bool {
	tv.status = true
	return true
}
func (tv *LgTV) switchOFF() bool {
	tv.status = false
	return true
}
func (tv *LgTV) GetStatus() bool {
	return tv.status
}
func (tv *LgTV) GetModel() string {
	return tv.model
}

//—  Есть общий интерфейс TVer для телевизоров, с методом switchOFF, GetStatus и GetModel.
//—  Функция switchOFF(tv Samsunger) переводит статус структуры в false и возвращает true.
//—  Функция switchOFF(tv LGer) переводит статус структуры в false и возвращает true.
//—  Функция switchOn переводит статус структуры в true и возвращает true.
//—  Объект типа Samsunger должен иметь метод SamsungHub().
//—  Объект типа LGer должен иметь метод LGHub().
//—  Все объекты типа Samsunger и LGer должны иметь методы GetStatus() (возвращающий состояние телевизора)
//и GetModel() (возвращающий название модели телевизора).

func main() {
	tv := &SamsungTV{
		status: true,
		model:  "Samsung XL-100500",
	}
	fmt.Println(tv.GetStatus())  // true
	fmt.Println(tv.GetModel())   // Samsung XL-100500
	fmt.Println(tv.SamsungHub()) // SamsungHub
	fmt.Println(tv.GetStatus())  // false
	fmt.Println(tv.switchOn())   // true
	fmt.Println(tv.GetStatus())  // true
	fmt.Println(tv.switchOFF())  // true
	fmt.Println(tv.GetStatus())  // false
}
