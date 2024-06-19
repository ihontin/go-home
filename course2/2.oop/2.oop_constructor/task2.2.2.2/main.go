package main

type User struct {
	ID       int
	Username string
	Email    string
	Role     string
}

type UserOption func(*User)

func WithUsername(s string) UserOption {
	return func(user *User) {
		user.Username = s
	}
}
func WithEmail(s string) UserOption {
	return func(user *User) {
		user.Email = s
	}
}
func WithRole(s string) UserOption {
	return func(user *User) {
		user.Role = s
	}
}

func NewUser(id int, options ...UserOption) *User {
	theUser := &User{
		id, "", "", "",
	}
	for _, option := range options {
		option(theUser)
	}
	return theUser
}

//func main() {
//	user := NewUser(1, WithUsername("testuser"), WithEmail("testuser@example.com"), WithRole("admin"))
//	fmt.Printf("User: %+v\n", user)
//}
