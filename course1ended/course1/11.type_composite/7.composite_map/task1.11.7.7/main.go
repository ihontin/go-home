package main

type User struct {
	Nickname string
	Age      int
	Email    string
}

// Функция getUniqueUsers должна принимать срез пользователей типа []User и возвращать срез уникальных пользователей по никнейму.
// Порядок пользователей в возвращаемом срезе должен быть сохранен.
func getUniqueUsers(users []User) []User {
	var usersMap = make(map[string]bool)
	var listUsers []User
	if len(users) == 0 {
		return users
	}
	for _, us := range users {
		if _, ok := usersMap[us.Nickname]; !ok {
			usersMap[us.Nickname] = true
			listUsers = append(listUsers, us)
		}
	}
	var outList = make([]User, len(listUsers))
	copy(outList, listUsers)
	return outList
}

//func main() {
//	var users = []User{
//		{"Jaba", 10, "e@ru"},
//		{"Doratain d", 30, "e@ru"},
//		{"Jaba", 50, "e@ru"},
//		{"", 12, ""},
//		{"", 12, ""},
//		{"Doratain d", 30, "e@ru"},
//	}
//	fmt.Println(getUniqueUsers(users))
//}
