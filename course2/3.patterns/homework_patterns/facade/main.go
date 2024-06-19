package main

import (
	"fmt"
	"net/http"
)

//Пример 1. Веб-фреймворк
//Фасад может быть использован в веб-фреймворках для предоставления удобного интерфейса при работе с различными
//компонентами, такими как маршрутизация, обработка запросов, авторизация и другие.
// Фасад для работы с веб-фреймворком

type WebFrameworkFacade struct {
	router     *Router
	authorizer *Authorizer
}

func NewWebFrameworkFacade() *WebFrameworkFacade {
	return &WebFrameworkFacade{
		router:     NewRouter(),
		authorizer: NewAuthorizer(),
	}
}

func (f *WebFrameworkFacade) HandleRequest(request *http.Request) {
	if f.authorizer.Authorize(request) {
		f.router.Route(request)
	} else {
		fmt.Println("Access denied")
	}
}

// Роутер
type Router struct{}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Route(request *http.Request) {
	fmt.Println("Routing request:", request.URL.Path)
}

// Авторизатор
type Authorizer struct{}

func NewAuthorizer() *Authorizer {
	return &Authorizer{}
}

func (a *Authorizer) Authorize(request *http.Request) bool {
	// Проверка авторизации пользователя
	return true
}

func main() {
	facade := NewWebFrameworkFacade()
	request, _ := http.NewRequest("GET", "/users", nil)
	facade.HandleRequest(request)
}

//Пример 2. Библиотека для работы с БД
//Фасад может быть применен для упрощения взаимодействия с базами данных, предоставляя унифицированный интерфейс
//для выполнения запросов, управления транзакциями и других операций.

// Фасад для работы с базой данных
//type DatabaseFacade struct {
//	connection *sql.DB
//}
//
//func NewDatabaseFacade() (*DatabaseFacade, error) {
//	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/database")
//	if err != nil {
//		return nil, err
//	}
//
//	return &DatabaseFacade{
//		connection: db,
//	}, nil
//}
//
//func (f *DatabaseFacade) ExecuteQuery(query string) ([]map[string]interface{}, error) {
//	rows, err := f.connection.Query(query)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//
//	columns, err := rows.Columns()
//	if err != nil {
//		return nil, err
//	}
//
//	result := make([]map[string]interface{}, 0)
//	values := make([]interface{}, len(columns))
//	valuePtrs := make([]interface{}, len(columns))
//	for rows.Next() {
//		for i := range columns {
//			valuePtrs[i] = &values[i]
//		}
//		err := rows.Scan(valuePtrs...)
//		if err != nil {
//			return nil, err
//		}
//
//		rowData := make(map[string]interface{})
//		for i, column := range columns {
//			rowData[column] = values[i]
//		}
//		result = append(result, rowData)
//	}
//
//	return result, nil
//}
//
//func main() {
//	facade, err := NewDatabaseFacade()
//	if err != nil {
//		fmt.Println("Failed to connect to the database:", err)
//		return
//	}
//
//	query := "SELECT * FROM users"
//	result, err := facade.ExecuteQuery(query)
//	if err != nil {
//		fmt.Println("Failed to execute query:", err)
//		return
//	}
//
//	fmt.Println("Query result:", result)
//}

//Мы рассмотрели два примера использования паттерна фасад в реальных проектах. В первом примере показано,
//как фасад может быть использован в веб-фреймворке для упрощения работы с различными компонентами.
//Во втором примере фасад применен для упрощения взаимодействия с базой данных. В обоих случаях
//фасад предоставляет унифицированный интерфейс, скрывая сложность и детали внутренней реализации подсистемы.
