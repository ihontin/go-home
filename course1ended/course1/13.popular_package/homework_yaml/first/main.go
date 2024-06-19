package main

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"log"
)

//Пример 1: Пример YAML-конфигурационного файла
//database:
//host: localhost
//port: 5432
//username: admin
//password: password123

//--------------------------------2

//Пример 1: Чтение YAML-файла и преобразование в структуру данных Golang

//type Config struct {
//	Database struct {
//		Host     string `yaml:"host"`
//		Port     int    `yaml:"port"`
//		Username string `yaml:"username"`
//		Password string `yaml:"password"`
//	} `yaml:"database"`
//}
//
//func main() {
//	// Чтение YAML-файла
//	myPathPro, err := filepath.Abs("course1/13.popular_package/homework_yaml/first/config.yml")
//	if err != nil {
//		fmt.Printf("%v", err)
//	}
//	data, err := os.ReadFile(myPathPro)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Преобразование YAML в структуру данных Golang
//	var config Config
//	err = yaml.Unmarshal(data, &config)
//	if err != nil {
//		log.Fatal(err)
//	}
//	// Вывод данных
//	fmt.Println("Host:", config.Database.Host)
//	fmt.Println("Port:", config.Database.Port)
//	fmt.Println("Username:", config.Database.Username)
//	fmt.Println("Password:", config.Database.Password)
//}

//-----------------------------------------2

// Пример 2: Запись данных в YAML-файл
//type Config struct {
//	Database struct {
//		Host     string `yaml:"host"`
//		Port     int    `yaml:"port"`
//		Username string `yaml:"username"`
//		Password string `yaml:"password"`
//	} `yaml:"database"`
//}
//
//func main() {
//	// Создание структуры данных
//	config := Config{
//		Database: struct {
//			Host     string `yaml:"host"`
//			Port     int    `yaml:"port"`
//			Username string `yaml:"username"`
//			Password string `yaml:"password"`
//		}{
//			Host:     "localhost",
//			Port:     5432,
//			Username: "admin",
//			Password: "password123",
//		},
//	}
//
//	// Запись данных в YAML-файл
//	myPathPro, err := filepath.Abs("course1/13.popular_package/homework_yaml/first/config.yml")
//	if err != nil {
//		fmt.Printf("%v", err)
//	}
//	file, err := os.Create(myPathPro)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer file.Close()
//	encoder := yaml.NewEncoder(file)
//	err = encoder.Encode(config)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("Данные успешно записаны в файл config.yml")
//}

//------------------------------------3

//Пример 3: Работа с массивами в YAML
//type Config struct {
//	Items []string `yaml:"items"`
//}
//
//func main() {
//	// Чтение YAML-файла
//	data := []byte(`
//items:
//- apple
//- banana
//- orange
//`)
//
//	var config Config
//	err := yaml.Unmarshal(data, &config)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Вывод элементов массива
//	for _, item := range config.Items {
//		fmt.Println(item)
//	}
//}

//---------------------------------------------------4

// Пример 4: Работа с вложенными структурами в YAML
//type Person struct {
//	Name string `yaml:"name"`
//	Age  int    `yaml:"age"`
//}
//
//type Config struct {
//	People []Person `yaml:"people"`
//}
//
//func main() {
//	// Чтение YAML-файла
//	data := []byte(`
//people:
//- name: John
// age: 30
//- name: Alice
// age: 25
//`)
//	var config Config
//	err := yaml.Unmarshal(data, &config)
//	if err != nil {
//		log.Fatal(err)
//	}
//	// Вывод информации о людях
//	for _, person := range config.People {
//		fmt.Println("Name:", person.Name)
//		fmt.Println("Age:", person.Age)
//	}
//}

//--------------------------------------5

// Пример 5: Работа с числами в YAML
type Config struct {
	Number int `yaml:"number"`
}

func main() {
	// Чтение YAML-файла
	data := []byte(`number: 42`)

	var config Config
	err := yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}

	// Вывод числа
	fmt.Println("Number:", config.Number)
}

// Пример 6: Работа с логическими значениями в YAML
//type Config struct {
//	Enabled bool `yaml:"enabled"`
//}
//
//func main() {
//	// Чтение YAML-файла
//	data := []byte(`
//enabled: true
//`)
//
//	var config Config
//	err := yaml.Unmarshal(data, &config)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Вывод логического значения
//	fmt.Println("Enabled:", config.Enabled)
//}

// Пример 7: Работа с пустыми значениями в YAML
//type Config struct {
//	Value string `yaml:"value"`
//}
//
//func main() {
//	// Чтение YAML-файла
//	data := []byte("value:")
//
//	var config Config
//	err := yaml.Unmarshal(data, &config)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Вывод пустого значения
//	fmt.Println("Value:", config.Value)
//}

//Пример 8: Обработка ошибок при чтении YAML-файла
//type Config struct {
//	Database struct {
//		Host string `yaml:"host"`
//		Port int `yaml:"port"`
//		Username string `yaml:"username"`
//		Password string `yaml:"password"`
//	} `yaml:"database"`
//}
//
//func main() {
//	// Чтение YAML-файла
//	data := []byte(`
//database:
//host: localhost
//port: 5432
//username: admin
//`)
//
//	var config Config
//	err := yaml.Unmarshal(data, &config)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Вывод данных
//	fmt.Println("Host:", config.Database.Host)
//	fmt.Println("Port:", config.Database.Port)
//	fmt.Println("Username:", config.Database.Username)
//	fmt.Println("Password:", config.Database.Password)
//}

// Пример 9: Поддержка различных форматов YAML
//type Config struct {
//	Value string `yaml:"value"`
//}
//
//func main() {
//	// Чтение YAML-файла в формате YAML 1.1
//	data := []byte(`
//value: example
//`)
//
//	var config Config
//	err := yaml.Unmarshal(data, &config)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Вывод значения
//	fmt.Println("Value:", config.Value)
//
//	// Чтение YAML-файла в формате YAML 1.2
//	data = []byte(`
//value: example
//`)
//
//	err = yaml.Unmarshal(data, &config)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Вывод значения
//	fmt.Println("Value:", config.Value)
//}
