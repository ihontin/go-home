package main

import (
	"encoding/json"
	"os"
	"path"
)

type User struct {
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Comments []Comment `json:"comments"`
}
type Comment struct {
	Text string `json:"text"`
}

// writeJSON(filePath string, data interface{}) error, которая будет записывать данные в формате JSON в файл.
// Функция должна создавать директорию, если она не существует, с использованием функции os.MkdirAll.
func writeJSON(filePath string, data interface{}) error {
	if err := os.MkdirAll(path.Dir(filePath), os.FileMode(0777)); err != nil {
		return err
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if err = os.WriteFile(filePath, jsonData, os.FileMode(0644)); err != nil {
		return err
	}
	return nil
}

//func main() {
//	// Write JSON data to file
//	data := []map[string]interface{}{
//		{
//			"name": "Elliot",
//			"age":  25,
//		},
//		{
//			"name": "Fraser",
//			"age":  30,
//		},
//	}
//	jsonFilePath, err := os.Getwd()
//	if err != nil {
//		panic(err)
//	}
//	err = writeJSON(jsonFilePath+"/course1/13.popular_package/8.package_json/rest_API_4.13.8.4/testdata/test.json", data)
//	if err != nil {
//		panic(err)
//	}
//}
