package main

import (
	"database/sql"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3"
	"strings"
)

type User struct {
	ID       int
	Username string
	Email    string
}

func CreateUserTable() error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		fmt.Println("Ошибка при подключении к базе данных:", err)
		return err
	}
	defer db.Close()
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    Username TEXT NOT NULL, 
    Email TEXT NOT NULL 
)`)
	if err != nil {
		return err
	}
	fmt.Println("Таблица успешно создана")
	return nil
}

func InsertUser(user User) error {
	fmt.Println("InsertUser", user)
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		fmt.Println("Ошибка при подключении к базе данных:", err)
		return err
	}
	defer db.Close()
	sqlDb, args, err := PrepareQuery("Insert", "users", user)
	if err != nil {
		fmt.Println("Ошибка при вставке данных в функции PrepareQuery:", err)
		return err
	}
	_, err = db.Exec(sqlDb, args...)
	if err != nil {
		fmt.Println("Ошибка при вставке данных:", err)
		return err
	}
	fmt.Println("Данные успешно вставлены")
	return nil
}

// SelectUser Выборка пользователя из таблицы
func SelectUser(userID int) (User, error) {
	fmt.Println("SelectUser", userID)
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		fmt.Println("Ошибка при создании пути:", err)
		return User{}, err
	}
	defer db.Close()
	var emptUser = User{ID: userID}
	sqlDb, _, err := PrepareQuery("Select", "users", emptUser)
	if err != nil {
		fmt.Println("Ошибка при сканировании данных:", err)
		return User{}, err
	}
	rows := db.QueryRow(sqlDb, userID)
	var scanUser User
	err = rows.Scan(&scanUser.ID, &scanUser.Username, &scanUser.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Ошибка поиска, юзер не найден:", err)
			return User{}, err
		} else {
			fmt.Println("Ошибка при сравнении данных:", err)
			return User{}, err
		}
	}
	return scanUser, nil
}

// UpdateUser Обновление информации о пользователе
func UpdateUser(user User) error {
	fmt.Println("UpdateUser", user)
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		fmt.Println("Ошибка при подключении к базе данных:", err)
		return err
	}
	defer db.Close()
	sqlDb, args, err := PrepareQuery("Select", "users", user)
	if err != nil {
		fmt.Println("Ошибка при сканировании данных:", err)
		return err
	}
	row := db.QueryRow(sqlDb, args...)
	var scanUser User
	err = row.Scan(&scanUser.ID, &scanUser.Username, &scanUser.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			sqlD, ars, err1 := PrepareQuery("Insert", "users", user)
			if err1 != nil {
				fmt.Println("Ошибка при вставке данных в функции PrepareQuery:", err1)
				return err1
			}
			_, err1 = db.Exec(sqlD, ars...)
			if err1 != nil {
				fmt.Println("Ошибка при вставке данных:", err1)
				return err1
			}
			fmt.Println("Данные успешно вставлены")
			return nil
		} else {
			fmt.Println("Ошибка при сравнении данных:", err)
			return err
		}
	}
	sqlDb2, args2, err := PrepareQuery("update", "users", user)
	if err != nil {
		fmt.Println("Ошибка при обновлении базы данных:", err)
		return err
	}
	_, err = db.Exec(sqlDb2, args2...)
	if err != nil {
		fmt.Println("Ошибка при обновлении данных:", err)
		return err
	}
	fmt.Println("Данные успешно обновлены")
	return nil
}

// Удаление пользователя из таблицы
func DeleteUser(userID int) error {
	fmt.Println("DeleteUser", userID)
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		fmt.Println("Ошибка при подключении к базе данных:", err)
		return err
	}
	sqlDb, _, err := PrepareQuery("delete", "users", User{ID: userID})
	if err != nil {
		fmt.Println("Ошибка при создании запроса на удаление:", err)
		return err
	}
	_, err = db.Exec(sqlDb, userID)
	if err != nil {
		fmt.Println("Ошибка при удалении данных:", err)
		return err
	}
	return nil
}

// PrepareQuery Функция для подготовки запроса
func PrepareQuery(operation string, table string, user User) (string, []interface{}, error) {
	operation = strings.ToLower(operation)
	newErr := errors.New("wrong operation")
	switch operation {
	case "insert":
		return sq.Insert(table).
			Columns("Username", "Email").
			Values(user.Username, user.Email).
			ToSql()
	case "select":
		return sq.Select("ID", "Username", "Email").
			From(table).
			Where(sq.Eq{"ID": user.ID}).
			ToSql()
	case "delete":
		return sq.Delete(table).
			Where(sq.Eq{"ID": user.ID}).
			ToSql()
	case "update":
		return sq.Update(table).
			Set("Username", user.Username).
			Set("Email", user.Email).
			Where(sq.Eq{"ID": user.ID}).
			ToSql()
	default:
		return "", []interface{}{}, newErr
	}
}

func main() {
	var newUsers = []User{
		{ID: 0, Username: "Kat", Email: "Kat@p.ru"},
		//{ID: 1, Username: "Bonon", Email: "Bonon@o.com"},
		//{ID: 2, Username: "Silimi", Email: "Silimi@lecum.i"},
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
	getUser, err := SelectUser(1)
	if err != nil {
		fmt.Println("Ошибка при выборке данных:", err)
		return
	}
	sam := User{ID: 1, Username: "Gaga", Email: "laga@gul.aga"}
	fmt.Println(getUser.ID, getUser.Username, getUser.Email)
	err = UpdateUser(sam)
	if err != nil {
		fmt.Println("Ошибка при обновлении данных:", err)
		return
	}
	fmt.Println("----------------")
	for i := 1; i < 3; i++ {
		g, _ := SelectUser(i)
		fmt.Println(g.ID, g.Username, g.Email)
	}
	getUser2, _ := SelectUser(1)
	fmt.Println(getUser2.ID, getUser2.Username, getUser2.Email, "- WILL BE DELETED")
	err = DeleteUser(1)
	if err != nil {
		fmt.Println("Ошибка при удалении данных:", err)
		return
	}
	getUser3, _ := SelectUser(1)
	fmt.Println(getUser3.ID, getUser3.Username, getUser3.Email)
}

//type User struct {
//	ID       int       `json:"id"`
//	Name     string    `json:"name"`
//	Age      int       `json:"age"`
//	Comments []Comment `json:"comments"`
//}
//
//type Comment struct {
//	ID     int    `json:"id"`
//	Text   string `json:"text"`
//	UserID int    `json:"user_id"`
//}
