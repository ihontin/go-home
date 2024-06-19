package main

import (
	"fmt"
	"testing"
)

func TestGetConfigFromYAML(t *testing.T) {
	var servData = []byte(`server:
  port: 8080
db:
  host: localhost
  port: 5432
  user: admin
  password: password123
`)
	got, err := getConfigFromYAML(servData)
	if err != nil {
		fmt.Errorf("error: %v", err)
	}
	gotStr := fmt.Sprintf("%s%s%s%s%s", got.Server.Port, got.Db.Host, got.Db.Port, got.Db.User, got.Db.Password)

	expected := "8080localhost5432adminpassword123"
	if expected != gotStr {
		t.Errorf("expected = %s, got = %s\n", expected, got)
	}

}
