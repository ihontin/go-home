package main

import "C"
import (
	"database/sql"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3"
	"strings"
)

type User struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	ID     int    `json:"id"`
	Text   string `json:"text"`
	UserID int    `json:"user_id"`
}

func CreateUserTable() error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		fmt.Println("Ошибка при подключении к базе данных:", err)
		return err
	}
	defer db.Close()
	statement, err := db.Prepare(`CREATE TABLE IF NOT EXISTS users (
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    Name TEXT NOT NULL, 
    Age INTEGER
)`)
	if err != nil {
		return err
	}
	statement.Exec()

	statement, err = db.Prepare(`CREATE TABLE IF NOT EXISTS comments (
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    Text TEXT NOT NULL, 
    UserID INTEGER, FOREIGN KEY(UserID) REFERENCES users(ID)
)`)
	if err != nil {
		return err
	}
	statement.Exec()

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
	sqlDb, args, err := prepareQuery("insert", "users", user).(sq.InsertBuilder).RunWith(db).ToSql()
	if err != nil {
		fmt.Println("Ошибка при вставке данных в функции prepareQuery:", err)
		return err
	}
	_, err = db.Exec(sqlDb, args...)
	if err != nil {
		fmt.Println("Ошибка при вставке данных:", err)
		return err
	}
	sqlDb1, args1, err := prepareQuery("insert", "comments", user).(sq.InsertBuilder).RunWith(db).ToSql()
	if err != nil {
		fmt.Println("Ошибка при вставке данных в функции prepareQuery:", err)
		return err
	}
	_, err = db.Exec(sqlDb1, args1...)
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
	userSelect := prepareQuery("Select", "users", emptUser)
	sqlDb, args, err := userSelect.(sq.SelectBuilder).RunWith(db).ToSql()
	if err != nil {
		fmt.Println("Ошибка при сканировании данных prepareQuery:", err)
		return User{}, err
	}
	rows, err := db.Query(sqlDb, args...)
	if err != nil {
		fmt.Println("Ошибка при сканировании данных Query:", err)
		return User{}, err
	}
	defer rows.Close()
	var scanUser User
	scanUser.Comments = []Comment{}
	for rows.Next() {
		var comment Comment
		err = rows.Scan(&scanUser.ID, &scanUser.Name, &scanUser.Age, &comment.ID, &comment.Text)
		if err != nil {
			fmt.Println("Ошибка при сканировании данных в цикле:", err)
			return User{}, err
		}
		scanUser.Comments = append(scanUser.Comments, comment)
	}
	return scanUser, nil
}

func UpdateUser(user User) error {
	fmt.Println("UpdateUser", user)
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		fmt.Println("Error with path creation: ", err)
		return err
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("Error db.Begin", err)
		return err
	}
	_, err = prepareQuery("update", "users", user).(sq.UpdateBuilder).RunWith(tx).Exec()
	if err != nil {
		tx.Rollback()
		fmt.Println("Ошибка при вставке данных table users:", err)
		return err
	}
	for _, comment := range user.Comments {
		_, err = sq.Update("comments").Set("Text", comment.Text).Where(sq.Eq{"ID": comment.ID}).
			PlaceholderFormat(sq.Question).RunWith(tx).Exec()
		if err != nil {
			tx.Rollback()
			fmt.Println("Ошибка при вставке данных table comments:", err)
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		fmt.Println("Ошибка при вставке данных tx.Commit:", err)
		return err
	}
	fmt.Println("User updated successfully")
	return nil
}
func DeleteUser(userID int) error {
	fmt.Println("DeleteUser", userID)
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		fmt.Println("Error with path creation: ", err)
		return err
	}
	var user User
	user, err = SelectUser(userID)
	if err != nil {
		fmt.Println("Удаление не возможно, пользователь не найден: ", err)
		return err
	}
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("Error db.Begin", err)
		return err
	}
	_, err = prepareQuery("delete", "users", user).(sq.DeleteBuilder).RunWith(tx).Exec()
	if err != nil {
		tx.Rollback()
		fmt.Println("Ошибка удаления данных, table users", err)
		return err
	}
	for _, comment := range user.Comments {
		delCom := sq.Delete("comments").Where(sq.Eq{"id": comment.ID})
		_, err = delCom.RunWith(tx).Exec()
		if err != nil {
			tx.Rollback()
			fmt.Println("Ошибка удаления данных, table comments", err)
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		fmt.Println("error tx.Commit", err)
		return err
	}
	return nil
}
func prepareQuery(operation string, table string, user User) interface{} {
	operation = strings.ToLower(operation)
	newErr := errors.New("wrong operation")
	if table == "users" {
		switch operation {
		case "insert":
			return sq.Insert("users").Columns("ID", "Name", "Age").Values(user.ID, user.Name, user.Age)
		case "select":
			return sq.Select("users.ID", "users.Name", "users.Age", "comments.ID", "Comments.Text").From("users").
				Join("comments ON users.ID = comments.UserID").Where(sq.Eq{"users.ID": user.ID})
		case "update":
			return sq.Update(table).Set("name", user.Name).
				Set("age", user.Age).Where(sq.Eq{"users.id": user.ID}).PlaceholderFormat(sq.Question)
		case "delete":
			return sq.Delete(table).Where(sq.Eq{"users.id": user.ID}).PlaceholderFormat(sq.Question)
		default:
			return newErr
		}
	} else if table == "comments" {
		switch operation {
		case "insert":
			comColuimns := sq.Insert("comments").Columns("ID", "Text", "UserID")
			for _, comment := range user.Comments {
				comColuimns = comColuimns.Values(comment.ID, comment.Text, comment.UserID)
			}
			return comColuimns
		default:
			return newErr
		}
	}
	return newErr
}

func main() {
	err := CreateUserTable()
	if err != nil {
		fmt.Println("Ошибка при создании базы данных:", err)
		return
	}
	var newUsers = []User{
		{ID: 0, Name: "Kat", Age: 13, Comments: []Comment{
			{0, "Something nice about his mother", 0},
			{1, "Life is GooD", 0},
			{2, "Nothing to say", 0},
		}},
		{ID: 1, Name: "Zortana", Age: 131, Comments: []Comment{
			{3, "We all gonna die", 1},
			{4, "Apocalypse is here", 1},
			{5, "Hell is near", 1},
		}},
	}
	for _, us := range newUsers {
		err = InsertUser(us)
		if err != nil {
			fmt.Println("Ошибка при вставке данных:", err)
			return
		}
	}
	for i := 0; i < 2; i++ {
		g, _ := SelectUser(i)
		fmt.Println(g.ID, g.Name, g.Age, "\nComments:  ", g.Comments[0].Text, "   ", g.Comments[1].Text, "   ", g.Comments[2].Text)
	}
	sam := User{ID: 1, Name: "Jojo", Age: 101, Comments: []Comment{
		{3, "Nothing else matters", 1},
		{4, "i forgot the head", 1},
		{5, "Rest of the rest", 1},
	}}
	err = UpdateUser(sam)
	if err != nil {
		fmt.Println("Ошибка при обновлении данных:", err)
		return
	}
	fmt.Println("----------------")
	for i := 0; i < 2; i++ {
		g, _ := SelectUser(i)
		fmt.Println(g.ID, g.Name, g.Age, "\nComments:  ", g.Comments[0].Text, "   ", g.Comments[1].Text, "   ", g.Comments[2].Text)
	}
	fmt.Println("User ID_1 - WILL BE DELETED")
	err = DeleteUser(1)
	if err != nil {
		fmt.Println("Ошибка при удалении данных:", err)
		return
	}
	g, _ := SelectUser(0)
	fmt.Println(g.ID, g.Name, g.Age, "\nComments:  ", g.Comments[0].Text, "   ", g.Comments[1].Text, "   ", g.Comments[2].Text)
}
