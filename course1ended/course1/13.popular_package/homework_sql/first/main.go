package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

// Подключение к базе данных SQLite
//func main() {
//	db, err := sql.Open("sqlite3", "test.db")
//	if err != nil {
//		fmt.Println("Ошибка при подключении к базе данных:", err)
//		return
//	}
//	defer db.Close()
//
//	// Дальнейшая работа с базой данных...
//}

// Пример 2: Подключение к базе данных MySQL
//func main() {
//	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/database")
//	if err != nil {
//		fmt.Println("Ошибка при подключении к базе данных:", err)
//		return
//	}
//	defer db.Close()
//
//	// Дальнейшая работа с базой данных...
//}

// //Пример 3: Подключение к базе данных PostgreSQL
//func main() {
//	db, err := sql.Open("postgres", "user=postgres password=123456 dbname=testdb sslmode=disable")
//	if err != nil {
//		fmt.Println("Ошибка при подключении к базе данных:", err)
//		return
//	}
//	defer db.Close()
//
//	// Дальнейшая работа с базой данных...
//}

// //Пример 4: Создание таблицы
//func main() {
//	db, err := sql.Open("sqlite3", "test.db")
//	if err != nil {
//		fmt.Println("Ошибка при подключении к базе данных:", err)
//		return
//	}
//	defer db.Close()
//
//	// Создание таблицы
//	_, err = db.Exec(`CREATE TABLE users (
//id INTEGER PRIMARY KEY AUTOINCREMENT,
//name TEXT,
//age INTEGER
//)`)
//	if err != nil {
//		fmt.Println("Ошибка при создании таблицы:", err)
//		return
//	}
//
//	fmt.Println("Таблица успешно создана")
//}

// //Пример 5: Вставка данных
//func main() {
//	db, err := sql.Open("sqlite3", "test.db")
//	if err != nil {
//		fmt.Println("Ошибка при подключении к базе данных:", err)
//		return
//	}
//	defer db.Close()
//
//	// Вставка данных
//	_, err = db.Exec("INSERT INTO users (name, age) VALUES (?, ?)", "John Doe", 25)
//	if err != nil {
//		fmt.Println("Ошибка при вставке данных:", err)
//		return
//	}
//
//	fmt.Println("Данные успешно вставлены")
//}

// //Пример 6: Выборка данных
//func main() {
//	db, err := sql.Open("sqlite3", "test.db")
//	if err != nil {
//		fmt.Println("Ошибка при подключении к базе данных:", err)
//		return
//	}
//	defer db.Close()
//
//	// Выборка данных
//	rows, err := db.Query("SELECT * FROM users")
//	if err != nil {
//		fmt.Println("Ошибка при выборке данных:", err)
//		return
//	}
//	defer rows.Close()
//
//	for rows.Next() {
//		var id int
//		var name string
//		var age int
//		err = rows.Scan(&id, &name, &age)
//		if err != nil {
//			fmt.Println("Ошибка при сканировании данных:", err)
//			return
//		}
//		fmt.Printf("ID: %d, Имя: %s, Возраст: %d\n", id, name, age)
//	}
//}

// //Пример 7: Обновление данных
//func main() {
//	db, err := sql.Open("sqlite3", "test.db")
//	if err != nil {
//		fmt.Println("Ошибка при подключении к базе данных:", err)
//		return
//	}
//	defer db.Close()
//
//	// Обновление данных
//	_, err = db.Exec("UPDATE users SET age = ? WHERE name = ?", 30, "John Doe")
//	if err != nil {
//		fmt.Println("Ошибка при обновлении данных:", err)
//		return
//	}
//
//	fmt.Println("Данные успешно обновлены")
//}

// //Пример 8: Удаление данных
//func main() {
//	db, err := sql.Open("sqlite3", "test.db")
//	if err != nil {
//		fmt.Println("Ошибка при подключении к базе данных:", err)
//		return
//	}
//	defer db.Close()
//
//	// Удаление данных
//	_, err = db.Exec("DELETE FROM users WHERE name = ?", "John Doe")
//	if err != nil {
//		fmt.Println("Ошибка при удалении данных:", err)
//		return
//	}
//
//	fmt.Println("Данные успешно удалены")
//}

////Пример 9: Транзакция
//func main() {
//	db, err := sql.Open("sqlite3", "test.db")
//	if err != nil {
//		fmt.Println("Ошибка при подключении к базе данных:", err)
//		return
//	}
//	defer db.Close()
//
//	// Начало транзакции
//	tx, err := db.Begin()
//	if err != nil {
//		fmt.Println("Ошибка при начале транзакции:", err)
//		return
//	}
//}

//	// Выполнение запросов в рамках транзакции
//	_, err = tx.Exec("INSERT INTO users (name, age) VALUES (?, ?)", "John Doe", 25)
//	if err != nil {
//		tx.Rollback()
//		fmt.Println("Ошибка при вставке данных:", err)
//		return
//	}
//
//	_, err = tx.Exec("UPDATE users SET age = ? WHERE name = ?", 30, "John Doe")
//	if err != nil {
//		tx.Rollback()
//		fmt.Println("Ошибка при обновлении данных:", err)
//		return
//	}
//
//	// Фиксация транзакции
//	err = tx.Commit()
//	if err != nil {
//		fmt.Println("Ошибка при фиксации транзакции:", err)
//		return
//	}
//
//	fmt.Println("Транзакция успешно выполнена")
//}

////Пример 10: Использование Prepared Statements
//func main() {
//	db, err := sql.Open("sqlite3", "test.db")
//	if err != nil {
//		fmt.Println("Ошибка при подключении к базе данных:", err)
//		return
//	}
//	defer db.Close()
//
//	// Подготовка запроса
//	stmt, err := db.Prepare("INSERT INTO users (name, age) VALUES (?, ?)")
//	if err != nil {
//		fmt.Println("Ошибка при подготовке запроса:", err)
//		return
//	}
//	defer stmt.Close()
//
//	// Выполнение запроса с разными значениями
//	_, err = stmt.Exec("John Doe", 25)
//	if err != nil {
//		fmt.Println("Ошибка при вставке данных:", err)
//		return
//	}
//
//	_, err = stmt.Exec("Jane Smith", 30)
//	if err != nil {
//		fmt.Println("Ошибка при вставке данных:", err)
//		return
//	}
//
//	fmt.Println("Данные успешно вставлены")
//}

////Пример 11: Использование транзакций с Prepared Statements
//func main() {
//	db, err := sql.Open("sqlite3", "test.db")
//	if err != nil {
//		fmt.Println("Ошибка при подключении к базе данных:", err)
//		return
//	}
//	defer db.Close()
//
//	// Начало транзакции
//	tx, err := db.Begin()
//	if err != nil {
//		fmt.Println("Ошибка при начале транзакции:", err)
//		return
//	}
//
//	// Подготовка запроса
//	stmt, err := tx.Prepare("INSERT INTO users (name, age) VALUES (?, ?)")
//	if err != nil {
//		tx.Rollback()
//		fmt.Println("Ошибка при подготовке запроса:", err)
//		return
//	}
//	defer stmt.Close()
//
//	// Выполнение запроса с разными значениями
//	_, err = stmt.Exec("John Doe", 25)
//	if err != nil {
//		tx.Rollback()
//		fmt.Println("Ошибка при вставке данных:", err)
//		return
//	}
//
//	_, err = stmt.Exec("Jane Smith", 30)
//	if err != nil {
//		tx.Rollback()
//		fmt.Println("Ошибка при вставке данных:", err)
//		return
//	}
//
//	// Фиксация транзакции
//	err = tx.Commit()
//	if err != nil {
//		fmt.Println("Ошибка при фиксации транзакции:", err)
//		return
//	}
//
//	fmt.Println("Транзакция успешно выполнена")
//}

////Пример 12: Получение информации о результате выполнения запроса
//func main() {
//	db, err := sql.Open("sqlite3", "test.db")
//	if err != nil {
//		fmt.Println("Ошибка при подключении к базе данных:", err)
//		return
//	}
//	defer db.Close()
//
//	// Выполнение запроса
//	result, err := db.Exec("DELETE FROM users WHERE age > ?", 30)
//	if err != nil {
//		fmt.Println("Ошибка при удалении данных:", err)
//		return
//	}
//
//	// Получение информации о результате выполнения запроса
//	rowsAffected, err := result.RowsAffected()
//	if err != nil {
//		fmt.Println("Ошибка при получении информации о результате выполнения запроса:", err)
//		return
//	}
//
//	fmt.Println("Количество затронутых строк:", rowsAffected)
//}

////Пример 13:Работа с транзакциями
//func main() {
//	db, err := sql.Open("sqlite3", "test.db")
//	if err != nil {
//		fmt.Println("Ошибка при подключении к базе данных:", err)
//		return
//	}
//	defer db.Close()
//
//	// Начало транзакции
//	tx, err := db.Begin()
//	if err != nil {
//		fmt.Println("Ошибка при начале транзакции:", err)
//		return
//	}
//
//	// Выполнение запросов в рамках транзакции
//	_, err = tx.Exec("INSERT INTO users (name, age) VALUES (?, ?)", "John Doe", 25)
//	if err != nil {
//		tx.Rollback()
//		fmt.Println("Ошибка при вставке данных:", err)
//		return
//	}
//
//	_, err = tx.Exec("UPDATE users SET age = ? WHERE name = ?", 30, "John Doe")
//	if err != nil {
//		tx.Rollback()
//		fmt.Println("Ошибка при обновлении данных:", err)
//		return
//	}
//
//	// Фиксация транзакции
//	err = tx.Commit()
//	if err != nil {
//		fmt.Println("Ошибка при фиксации транзакции:", err)
//		return
//	}
//
//	fmt.Println("Транзакция успешно завершена")
//}
