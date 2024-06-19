package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	file, _ := os.Create("log.txt")
	defer file.Close()
	expected := "well done"
	fileLogger := FileLogger{file: file}
	fileLogger.Log(expected)
	file, err := os.OpenFile("log.txt", os.O_RDWR, os.FileMode(0755))
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(file)
	got, err := r.ReadString('\n')
	if got != expected+"\n" {
		t.Errorf("expected = %v, got = %v", expected, got)
	}
}
func TestLogSystemLog(t *testing.T) {
	file, _ := os.Create("log.txt")
	defer file.Close()
	expected := "well done"
	fileLogger := FileLogger{file: file}
	logSystem := NewLogSystem(WithLogger(fileLogger))
	logSystem.Log(expected)
	file, err := os.OpenFile("log.txt", os.O_RDWR, os.FileMode(0755))
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(file)
	got, err := r.ReadString('\n')
	if got != expected+"\n" {
		t.Errorf("expected = %v, got = %v", expected, got)
	}
}

func TestMainFunc(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()
	w.Close()
	os.Stdout = old

	expected := "all wright\n"
	var osStdout bytes.Buffer
	osStdout.ReadFrom(r)
	if expected != osStdout.String() {
		t.Errorf("expected = %v, got = %v", expected, osStdout.String())
	}
}

func TestWithLogger(t *testing.T) {
	file, _ := os.Create("log.txt")
	defer file.Close()
	fileLogger := FileLogger{file: file}
	expecteds := []Logger{fileLogger, &FileLogger{}}
	testlogS := &LogSystem{}
	for _, expected := range expecteds {
		WithLogger(expected)(testlogS)
		if expected != testlogS.Logger {
			t.Errorf("expected = %v, got = %v", expected, testlogS.Logger)
		}
	}
}

func TestNewLogSystem(t *testing.T) {
	file, _ := os.Create("log.txt")
	defer file.Close()

	fileLogger := FileLogger{file: file}
	expected := NewLogSystem(WithLogger(fileLogger))

	testO := &LogSystem{fileLogger}
	if expected.Logger != testO.Logger {
		t.Errorf("expected = %v, got = %v", expected.Logger, testO.Logger)
	}
}
