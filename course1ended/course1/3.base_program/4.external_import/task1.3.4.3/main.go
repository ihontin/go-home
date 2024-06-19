package main

import (
	"fmt"
	"github.com/icrowley/fake"
)

func GenerateFakeData() string {
	name, address, phone, email := fake.FullName(), fake.StreetAddress(), fake.Phone(), fake.EmailAddress()
	return fmt.Sprintf("Name: %s\nAddress: %s\nPhone: %s\nEmail: %s\n", name, address, phone, email)
}

func main() {
	fmt.Println(GenerateFakeData())
}
