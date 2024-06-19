package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Пример 1
// В этом примере мы продемонстрируем использование HTTP-клиента для отправки GET-запроса и получения ответа от сервера.
func main() {
	// Создаем новый HTTP-клиент
	client := &http.Client{}

	// Создаем GET-запрос
	req, err := http.NewRequest("GET", "https://mholt.github.io/curl-to-go/", nil)
	if err != nil {
		fmt.Println("Ошибка при создании запроса:", err)
		return
	}

	// Отправляем запрос и получаем ответ
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка при отправке запроса:", err)
		return
	}
	defer resp.Body.Close()

	// Читаем тело ответа
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}

	// Выводим ответ на экран
	fmt.Println(string(body))
}

//Пример 2
//В этом примере мы продемонстрируем использование HTTP-клиента для отправки POST-запроса с данными в формате JSON и получения ответа от сервера.

//func main() {
//	// Создаем новый HTTP-клиент
//	client := &http.Client{}
//
//	// Создаем тело запроса в формате JSON
//	jsonData := []byte(`{"name": "John", "age": 30}`)
//
//	// Создаем POST-запрос
//	req, err := http.NewRequest("POST", "https://api.example.com/data", bytes.NewBuffer(jsonData))
//	if err != nil {
//		fmt.Println("Ошибка при создании запроса:", err)
//		return
//	}
//
//	// Устанавливаем заголовок Content-Type
//	req.Header.Set("Content-Type", "application/json")
//
//	// Отправляем запрос и получаем ответ
//	resp, err := client.Do(req)
//	if err != nil {
//		fmt.Println("Ошибка при отправке запроса:", err)
//		return
//	}
//	defer resp.Body.Close()
//
//	// Читаем тело ответа
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		fmt.Println("Ошибка при чтении ответа:", err)
//		return
//	}
//
//	// Выводим ответ на экран
//	fmt.Println(string(body))
//}

//Пример 3
//Unmarshall JSON в структуру

//type Response struct {
//	UserId int    `json:"userId"`
//	Id     int    `json:"id"`
//	Title  string `json:"title"`
//	Body   string `json:"body"`
//}
//
//func main() {
//	// Создаем экземпляр HTTP клиента
//	client := &http.Client{}
//
//	// Создаем GET запрос к API
//	req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts/1", nil)
//	if err != nil {
//		fmt.Println("Ошибка при создании GET запроса:", err)
//		return
//	}
//
//	// Отправляем запрос и получаем ответ
//	resp, err := client.Do(req)
//	if err != nil {
//		fmt.Println("Ошибка при выполнении GET запроса:", err)
//		return
//	}
//	defer resp.Body.Close()
//
//	// Распаковываем JSON в структуру
//	var response Response
//	err = json.NewDecoder(resp.Body).Decode(&response)
//	if err != nil {
//		fmt.Println("Ошибка при распаковке JSON:", err)
//		return
//	}
//
//	// Выводим данные из структуры
//	fmt.Println("UserID:", response.UserId)
//	fmt.Println("ID:", response.Id)
//	fmt.Println("Title:", response.Title)
//	fmt.Println("Body:", response.Body)
//}
