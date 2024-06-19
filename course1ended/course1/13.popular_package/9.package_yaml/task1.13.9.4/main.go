package main

import (
	"github.com/go-yaml/yaml"
	"os"
	"path"
)

//writeYAML(filePath string, data interface{}) error, которая будет создавать файл YAML и записывать в него данные.
//	Если директория, в которой должен быть создан файл, не существует, функция должна создать ее с помощью os.MkdirAll.

func writeYAML(filePath string, data interface{}) error {
	if err := os.MkdirAll(path.Dir(filePath), os.FileMode(0777)); err != nil {
		return err
	}
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	enc := yaml.NewEncoder(file)
	if err = enc.Encode(data); err != nil {
		return err
	}
	return nil
}
