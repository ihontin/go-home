package main

//-------------------------1

//В этом примере мы сериализуем структуру данных в формат JSON с использованием пакета json.
//type Person struct {
//	Name  string `json:"name"`
//	Age   int    `json:"age"`
//	Email string `json:"email"`
//}
//func main() {
//	person := Person{
//		Name:  "Иван",
//		Age:   30,
//		Email: "ivan@example.com",
//	}
//	jsonData, err := json.Marshal(person)
//	if err != nil {
//		fmt.Println("Ошибка сериализации в JSON:", err)
//		return
//	}
//	fmt.Println(string(jsonData))
//}

//-------------------------2

// В этом примере мы десериализуем данные из формата JSON в структуру данных с использованием пакета json.
//type Person struct {
//	Name  string `json:"name"`
//	Age   int    `json:"age"`
//	Email string `json:"email"`
//}
//func main() {
//	jsonData := []byte(`{"name":"Анна","age":25,"email":"anna@example.com"}`)
//	var person Person
//	err := json.Unmarshal(jsonData, &person)
//	if err != nil {
//		fmt.Println("Ошибка десериализации из JSON:", err)
//		return
//	}
//	fmt.Println(person.Name, person.Age, person.Email)
//}

//-------------------------3

// В этом примере мы используем теги структур для контроля имён полей в JSON.
//type Person struct {
//	Name  string `json:"full_name"`
//	Age   int    `json:"age"`
//	Email string `json:"email"`
//}
//func main() {
//	person := Person{
//		Name:  "Мария",
//		Age:   35,
//		Email: "maria@example.com",
//	}
//	jsonData, err := json.Marshal(person)
//	if err != nil {
//		fmt.Println("Ошибка сериализации в JSON:", err)
//		return
//	}
//	fmt.Println(string(jsonData))
//}

//-------------------------4

//В этом примере мы читаем данные из файла в формате JSON с использованием пакета json.

//type Person struct {
//	Name string `json:"name"`
//	Age int `json:"age"`
//	Email string `json:"email"`
//}
//func main() {
//	fileData, err := ioutil.ReadFile("data.json")
//	if err != nil {
//		fmt.Println("Ошибка чтения файла:", err)
//		return
//	}
//	var person Person
//	err = json.Unmarshal(fileData, &person)
//	if err != nil {
//		fmt.Println("Ошибка десериализации из JSON:", err)
//		return
//	}
//	fmt.Println(person.Name, person.Age, person.Email)
//}

//------------------------------------5

//В этом примере мы записываем данные в файл в формате JSON с использованием пакета json.
//type Person struct {
//	Name  string `json:"name"`
//	Age   int    `json:"age"`
//	Email string `json:"email"`
//}
//func main() {
//	person := Person{
//		Name:  "Петр",
//		Age:   40,
//		Email: "petr@example.com",
//	}
//	jsonData, err := json.Marshal(person)
//	if err != nil {
//		fmt.Println("Ошибка сериализации в JSON:", err)
//		return
//	}
//	err = ioutil.WriteFile("data.json", jsonData, 0644)
//	if err != nil {
//		fmt.Println("Ошибка записи в файл:", err)
//		return
//	}
//	fmt.Println("Данные успешно записаны в файл.")
//}

//------------------------------------6

// В этом примере мы работаем с потоками данных JSON при чтении данных из сетевого соединения.
//type Person struct {
//	Name  string `json:"name"`
//	Age   int    `json:"age"`
//	Email string `json:"email"`
//}
//
//func main() {
//	resp, err := http.Get("https://api.example.com/person/1")
//	if err != nil {
//		fmt.Println("Ошибка при выполнении HTTP-запроса:", err)
//		return
//	}
//	defer resp.Body.Close()
//	var person Person
//	err = json.NewDecoder(resp.Body).Decode(&person)
//	if err != nil {
//		fmt.Println("Ошибка десериализации из JSON:", err)
//		return
//	}
//	fmt.Println(person.Name, person.Age, person.Email)
//}

//------------------------------------7

// В этом примере мы работаем с нестандартными типами данных при сериализации и десериализации JSON.
//type Person struct {
//	Name      string    `json:"name"`
//	BirthDate time.Time `json:"birth_date"`
//}
//
//func main() {
//	person := Person{
//		Name:      "Алексей",
//		BirthDate: time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
//	}
//	jsonData, err := json.Marshal(person)
//	if err != nil {
//		fmt.Println("Ошибка сериализации в JSON:", err)
//		return
//	}
//	fmt.Println(string(jsonData))
//}

//------------------------------------8

//В этом примере мы обрабатываем ошибки при десериализации данных из формата JSON.
//type Person struct {
//	Name string `json:"name"`
//	Age int `json:"age"`
//	Email string `json:"email"`
//}
//func main() {
//	jsonData := []byte(`{"name":"John","age":"30","email":"john@example.com"}`)
//	var person Person
//	err := json.Unmarshal(jsonData, &person)
//	if err != nil {
//		fmt.Println("Ошибка десериализации из JSON:", err)
//		return
//	}
//	fmt.Println(person.Name, person.Age, person.Email)
//}
