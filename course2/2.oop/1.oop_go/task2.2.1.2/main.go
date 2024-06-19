package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Logger interface {
	Log(message string) error
}

type ConsoleLogger struct {
	Writer io.Writer
}

func (c *ConsoleLogger) Log(message string) error {
	_, err := fmt.Fprintln(c.Writer, message)
	return err
}

type FileLogger struct {
	File *os.File
}

func (f *FileLogger) Log(message string) error {
	writer := bufio.NewWriter(f.File)
	_, err := writer.Write([]byte(message))
	if err != nil {
		return err
	}
	err = writer.Flush()
	return err
}

type RemoteLogger struct {
	Address string
}

func (r *RemoteLogger) Log(message string) error {
	// Имитация отправки сообщения на удаленный сервер
	url := r.Address

	text := bytes.NewBuffer([]byte(message))
	req, err := http.NewRequest("POST", url, text)
	if err != nil {
		return err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("ошибка, код ответа: %v", resp.StatusCode)
	}

	fmt.Println("Sending log message to remote server:", r.Address)
	return nil
}

func LogAll(loggers []Logger, message string) {
	for _, logger := range loggers {
		err := logger.Log(message)
		if err != nil {
			log.Println("Failed to log message:", err)
		}
	}
}

func main() {
	consoleLogger := &ConsoleLogger{Writer: os.Stdout}
	file, err := os.OpenFile("logfile.log", os.O_CREATE|os.O_RDWR, os.FileMode(0655))
	if err != nil {
		fmt.Println("ошибка открытия файла:", err)
	}
	defer file.Close()
	fileLogger := &FileLogger{File: file}
	remoteLogger := &RemoteLogger{Address: "http://example.com"}

	loggers := []Logger{consoleLogger, fileLogger, remoteLogger}
	LogAll(loggers, "This is a test log message.")
}
