package main

import (
	"github.com/go-yaml/yaml"
	"os"
)

// writeYAML(filePath string, data []User) error, которая будет записывать данные в формате YAML в файл.
func writeYAML(filePath string, data []User) error {
	if len(data) < 1 {
		return nil
	}
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := yaml.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return err
	}
	return nil
}

type User struct {
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	Text string `json:"text"`
}
