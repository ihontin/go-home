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

// writeJSON (filePath string, data []User) error, которая будет записывать данные пользователей в формате JSON в указанный файл.
// Функция writeJSON должна создавать директорию, если она не существует, с помощью функции os.MkdirAll.
func writeJSON(filePath string, data []User) error {
	if len(data) < 1 {
		return nil
	}
	err := os.MkdirAll(path.Dir(filePath), 0777)
	if err != nil {
		return err
	}
	writJson, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = os.WriteFile(filePath, writJson, os.FileMode(0644))
	if err != nil {
		return err
	}
	return nil
}

//func main() {
//
//	var newUser = []User{
//		{"John Doe", 30, []Comment{
//			{"Comment 1"},
//			{"Comment 2"},
//		},
//		},
//		{"Jane Smith", 25, []Comment{
//			{"Comment 3"},
//			{"Comment 4"},
//		},
//		},
//	}
//	jsonFilePath, err := os.Getwd()
//	if err != nil {
//		fmt.Errorf("%v", err.Error())
//	}
//	err = writeJSON(jsonFilePath+"/course1/13.popular_package/8.package_json/rest_API_4.13.8.3/testdata/test.json", newUser)
//	if err != nil {
//		fmt.Errorf("%v", err.Error())
//	}
//}
