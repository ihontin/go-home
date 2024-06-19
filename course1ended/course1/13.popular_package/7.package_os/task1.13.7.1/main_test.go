package main

import (
	"bytes"
	"os"
	"testing"
)

func TestMainFunc(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = old

	var reader = bytes.Buffer{}
	reader.ReadFrom(r)

	expected := ""
	if reader.String() != expected {
		t.Errorf("expected = %s, got = %s", expected, reader.String())
	}
}

//func TestWriteFile(t *testing.T) {
//	myPath := "/home/i/work/src/studentgit.kata.academy/Alkolex/go-kata/course1/13.popular_package/7.package_os/rest_API_4.13.7.1/file.txt"
//	err := WriteFile(myPath, []byte("Hello, World!"), os.FileMode(0644))
//	if err != nil {
//		fmt.Errorf("error: %v", err.Error())
//	}
//	if _, err = os.Stat(myPath); err != nil {
//		t.Errorf("file - %s \n", myPath)
//	}
//}
