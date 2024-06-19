package models

// easyjson -all models.go
//команда терминала для генерации файла easyjson

type Users struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender bool   `json:"gender"`
}
