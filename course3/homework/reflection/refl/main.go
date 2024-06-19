package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num int = 42
	fmt.Println(reflect.TypeOf(num))
}

//Пример 2
//В этом примере мы используем рефлексию для получения информации о полях структуры.
//type Person struct {
//	Name string
//	Age  int
//}
//
//func main() {
//	p := Person{Name: "John", Age: 30}
//	t := reflect.TypeOf(p)
//
//	for i := 0; i < t.NumField(); i++ {
//		field := t.Field(i)
//		fmt.Printf("Field name: %s, Field type: %s\n", field.Name, field.Type)
//	}
//}

//Пример 3
//В этом примере мы используем рефлексию для вызова метода объекта.
//type Person struct {
//	Name string
//	Age  int
//}
//
//func (p Person) SayHello() {
//	fmt.Println("Hello, my name is", p.Name)
//}
//
//func main() {
//	p := Person{Name: "John", Age: 30}
//	v := reflect.ValueOf(p)
//	m := v.MethodByName("SayHello")
//	m.Call(nil)
//}

//Пример 4
//В этом примере мы используем рефлексию для изменения значения поля структуры.
//type Person struct {
//	Name string
//	Age  int
//}
//
//func main() {
//	p := Person{Name: "John", Age: 30}
//	v := reflect.ValueOf(&p).Elem()
//	f := v.FieldByName("Name")
//	f.SetString("Alice")
//	fmt.Println(p.Name)
//}

//Пример 5
//В этом примере мы используем рефлексию для получения информации о методах структуры.
//type Person struct {
//	Name string
//	Age  int
//}
//
//func (p Person) SayHello() {
//	fmt.Println("Hello, my name is", p.Name)
//}
//
//func main() {
//	p := Person{Name: "John", Age: 30}
//	t := reflect.TypeOf(p)
//
//	for i := 0; i < t.NumMethod(); i++ {
//		method := t.Method(i)
//		fmt.Println("Method name:", method.Name)
//	}
//}

//Пример 6
//В этом примере мы используем рефлексию для создания нового объекта по типу.
//type Person struct {
//	Name string
//	Age  int
//}
//
//func main() {
//	t := reflect.TypeOf(Person{})
//	v := reflect.New(t).Elem()
//	v.FieldByName("Name").SetString("John")
//	v.FieldByName("Age").SetInt(30)
//	p := v.Interface().(Person)
//	fmt.Println(p)
//}

//Пример 7
//В этом примере мы используем рефлексию для проверки, является ли значение указателем.

//func main() {
//	var num int = 42
//	var ptr *int = &num
//
//	fmt.Println(reflect.ValueOf(num).Kind() == reflect.Ptr)  // false
//	fmt.Println(reflect.ValueOf(ptr).Elem().Kind() == reflect.Ptr)  // true
//}

//Пример 8
//В этом примере мы используем рефлексию для вызова функции по имени.
//func SayHello(name string) {
//	fmt.Println("Hello,", name)
//}
//
//func main() {
//	funcValue := reflect.ValueOf(SayHello)
//	args := []reflect.Value{reflect.ValueOf("John")}
//	funcValue.Call(args)
//}
