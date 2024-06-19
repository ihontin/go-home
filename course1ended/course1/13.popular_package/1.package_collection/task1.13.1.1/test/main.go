//go:generate grizzly generate main.go

package main

import (
	"fmt"
)

//grizzly:generate
type User struct {
	Nickname string
	Age      int
	Id       int
	Name     string
}

func getUsersByCondition(users []*User, condition string) []*User {
	newUsers := NewUserCollection(users)
	var chaseUsers []*User
	switch condition {
	case "age > 18":
		chaseUsers = newUsers.Filter(func(user *User) bool {
			return user.Age > 18
		}).Items
	case "age < 18":
		chaseUsers = newUsers.Filter(func(user *User) bool {
			return user.Age < 18
		}).Items
	case "age >= 18":
		chaseUsers = newUsers.Filter(func(user *User) bool {
			return user.Age >= 18
		}).Items
	case "age <= 18":
		chaseUsers = newUsers.Filter(func(user *User) bool {
			return user.Age <= 18
		}).Items
	case "age = 18":
		chaseUsers = newUsers.Filter(func(user *User) bool {
			return user.Age == 18
		}).Items
	default:
		return nil
	}
	return chaseUsers
}

func getUsersByAge(users []*User, age int) []*User {
	newUsers := NewUserCollection(users)
	return newUsers.Filter(func(user *User) bool {
		return user.Age == age
	}).Items
}

func getUsersByNickName(users []*User, nickName string) []*User {
	newUsers := NewUserCollection(users)
	return newUsers.Filter(func(user *User) bool {
		return user.Nickname == nickName
	}).Items
}

func getUsersUniqueNickName(users []*User) []*User {
	newUsers := NewUserCollection(users)
	return newUsers.UniqByNickname().Items
}
func main() {
	users := []*User{
		{Id: 1, Name: "John", Age: 20, Nickname: "Purple"},
		{Id: 2, Name: "Tom", Age: 22, Nickname: "Green"},
		{Id: 3, Name: "Billy", Age: 20, Nickname: "Black"},
		{Id: 4, Name: "Mister X", Age: 30, Nickname: "Green"},
	}

	getUniqNickUsers := getUsersUniqueNickName(users)
	for _, us := range getUniqNickUsers {
		fmt.Println("Unique NickName = ", *us)
	}
	fmt.Println()
	getByCondition := getUsersByCondition(users, "age > 18")
	for _, us := range getByCondition {
		fmt.Println("Users Condition age > 18 = ", (*us).Name, (*us).Age)
	}
	fmt.Println()
	getByAge := getUsersByAge(users, 20)
	for _, us := range getByAge {
		fmt.Println("Users Age 20 = ", (*us).Name, (*us).Age)
	}
	fmt.Println()
	getByNickName := getUsersByNickName(users, "Green")
	for _, us := range getByNickName {
		fmt.Println("getUsersByNickName = ", (*us).Name, (*us).Age)
	}

}
