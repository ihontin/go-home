package main

import (
	"encoding/json"
	"github.com/go-yaml/yaml"
)

// unmarshal(data []byte, v interface{}) error, которая будет выполнять декодирование данных
// в формате YAML или JSON в соответствующую структуру данных.
func unmarshal(data []byte, v interface{}) error {
	if err := yaml.Unmarshal(data, v); err == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return err
	}
	return nil
}

//type Person struct {
//	Name string `json:"name" yaml:"name"`
//	Age  int    `json:"age" yaml:"age"`
//}
//
//func main() {
//	data := []byte(`{"name":"John","age":30}`)
//	var person Person
//	err := unmarshal(data, &person)
//	if err != nil {
//		fmt.Println("Ошибка декодирования данных:", err)
//		return
//	}
//	fmt.Println("Имя:", person.Name)
//	fmt.Println("Возраст:", person.Age)
//	dataYml := []byte(`name: FIlth
//age: 11`)
//	err = unmarshal(dataYml, &person)
//	if err != nil {
//		fmt.Println("Ошибка декодирования данных:", err)
//		return
//	}
//	fmt.Println("Имя:", person.Name)
//	fmt.Println("Возраст:", person.Age)
//}
