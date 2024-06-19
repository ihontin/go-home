package main

import (
	"fmt"
	"testing"
)

func TestGetYAML(t *testing.T) {
	config := []Config{{
		Server: Server{"8080"},
		Db: Db{
			"localhost",
			"5432",
			"admin",
			"password123",
		},
	},
	}
	expected := "- server:\n    port: \"8080\"\n  db:\n    host: localhost\n    port: \"5432\"\n    user: admin\n    password: password123\n"
	got, err := getYAML(config)
	if err != nil {
		fmt.Println(err)
	}
	if expected != got {
		t.Errorf("expected = %s, got = %s\n", expected, got)
	}

}
