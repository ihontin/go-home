package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// CreateUserTable () - функция для создания таблицы пользователей в базе данных.
func CreateUserTable() error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		fmt.Println("Ошибка при подключении к базе данных:", err)
		return err
	}
	defer db.Close()
	_, err = db.Exec(`CREATE TABLE users (
id INTEGER PRIMARY KEY AUTOINCREMENT,
name TEXT,
age INTEGER
)`)
	if err != nil {
		return err
	}
	fmt.Println("Таблица успешно создана")
	return nil
}

// InsertUser(user User) - функция для добавления нового пользователя в таблицу.
func InsertUser(user User) error {

	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		fmt.Println("Ошибка при подключении к базе данных:", err)
		return err
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO users (name, age) VALUES (?, ?)", user.Name, user.Age)
	if err != nil {
		fmt.Println("Ошибка при вставке данных:", err)
		return err
	}
	fmt.Println("Данные успешно вставлены")
	return nil
}

// SelectUser(id int) User - функция для выборки пользователя по его идентификатору.
func SelectUser(id int) (User, error) {

	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		fmt.Println("Ошибка при подключении к базе данных:", err)
		return User{}, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM users WHERE id = ? ORDER BY RANDOM()", id)
	if err != nil {
		fmt.Println("Ошибка при выборке данных:", err)
		return User{}, err
	}
	defer rows.Close()
	var rowUser User
	for rows.Next() {
		err = rows.Scan(&rowUser.ID, &rowUser.Name, &rowUser.Age)
		if err != nil {
			fmt.Println("Ошибка при сканировании данных:", err)
			return User{}, err
		}
	}
	return rowUser, nil
}

// UpdateUser(user User) - функция для обновления информации о пользователе.
func UpdateUser(user User) error {

	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		fmt.Println("Ошибка при подключении к базе данных:", err)
		return err
	}
	defer db.Close()
	_, err = db.Exec("UPDATE users SET name = ?, age = ? WHERE id = ?", user.Name, user.Age, user.ID)
	if err != nil {
		fmt.Println("Ошибка при обновлении данных:", err)
		return err
	}
	fmt.Println("Данные успешно обновлены")
	return nil
}

// DeleteUser(id int) - функция для удаления пользователя из таблицы.
func DeleteUser(id int) error {

	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		fmt.Println("Ошибка при подключении к базе данных:", err)
		return err
	}
	defer db.Close()
	_, err = db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		fmt.Println("Ошибка при удалении данных:", err)
		return err
	}
	fmt.Println("Данные успешно удалены")
	return nil
}
func main() {
	var newUsers = []User{
		{1, "Kat", 7},
		{2, "Banan", 88},
		{3, "Salama", 101},
	}
	err := CreateUserTable()
	if err != nil {
		fmt.Println("Ошибка при создании базы данных:", err)
		return
	}
	for _, us := range newUsers {
		err = InsertUser(us)
		if err != nil {
			fmt.Println("Ошибка при вставке данных:", err)
			return
		}
	}
	getUser, err := SelectUser(2)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println(getUser.ID, getUser.Name, getUser.Age)
	err = UpdateUser(User{2, "Karakum", 35})
	if err != nil {
		fmt.Println("Ошибка при обновлении данных:", err)
		return
	}
	getUser2, err := SelectUser(2)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println(getUser2.ID, getUser2.Name, getUser2.Age)
	err = DeleteUser(2)
	if err != nil {
		fmt.Println("Ошибка при удалении данных:", err)
		return
	}
	getUser3, err := SelectUser(2)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println(getUser3.ID, getUser3.Name, getUser3.Age)
}
