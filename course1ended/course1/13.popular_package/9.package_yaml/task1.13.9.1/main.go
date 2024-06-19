package main

import (
	"fmt"
	"github.com/go-yaml/yaml"
)

type Config struct {
	Server Server `yaml:"server"`
	Db     Db     `yaml:"db"`
}

type Server struct {
	Port string `yaml:"port"`
}

type Db struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// getYAML, которая принимает на вход срез структур типа Config и возвращает строку в формате YAML.
// Функция должна преобразовывать данные из среза Config в YAML-строку.
func getYAML(c []Config) (string, error) {
	res, err := yaml.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

func main() {
	config := []Config{{
		Server: Server{"8080"},
		Db: Db{
			"localhost",
			"5432",
			"admin",
			"password123",
		},
	},
	}
	res, err := getYAML(config)
	if err != nil {
		fmt.Printf("%v", err)
	} else {
		fmt.Println(res)
	}
}

// - server:
// port: "8080"
// db:
// host: localhost
// port: "5432"
// user: admin
// password: password123
