package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	jiter "github.com/json-iterator/go"
	"github.com/mailru/easyjson"
	"io"
	"net/http"
	"studentgit.kata.academy/Alkolex/go-kata/course2/5.optimization/2.optimization_jsoniter/task2.5.2.1/models"
)

type MarshalUnmarshaler interface {
	Marshal(v any) ([]byte, error)
	Unmarshal(data []byte, v any) error
}

type jsonMarsh struct {
	models.Whetherer
}

func (j *jsonMarsh) Marshal(v any) ([]byte, error) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(v); err != nil {
		fmt.Println("Error marshaling:", err)
		return []byte{}, err
	}
	return buf.Bytes(), nil

	//outB, err := json.Marshal(v)
	//if err != nil {
	//	return []byte{}, err
	//}
	//return outB, err
}
func (j *jsonMarsh) Unmarshal(data []byte, v any) error {
	var buf = bytes.NewBuffer(data)
	enc := json.NewDecoder(buf)
	if err := enc.Decode(v); err != nil {
		fmt.Println("Error marshaling:", err)
		return err
	}
	return nil
	//err := json.Unmarshal(data, v)
	//if err != nil {
	//	return err
	//}
	//return err
}

// -------------------------------------------json-jiter
type jsonIter struct {
	models.Whetherer
}

func (j *jsonIter) Marshal(v any) ([]byte, error) {
	outB, err := jiter.Marshal(v)
	if err != nil {
		return []byte{}, err
	}
	return outB, err
}
func (j *jsonIter) Unmarshal(data []byte, v any) error {
	err := jiter.Unmarshal(data, v)
	if err != nil {
		return err
	}
	return err
}

// --------------------------------------------easyjson

type EasyMarshJson struct {
	MarJs models.Whetherer
}

func (j *EasyMarshJson) Marshal(v any) ([]byte, error) {
	s := v.(models.Whetherer)
	outB, err := easyjson.Marshal(s)
	if err != nil {
		return []byte{}, err
	}
	return outB, err
}
func (j *EasyMarshJson) Unmarshal(data []byte, v any) error {
	s := v.(*models.Whetherer)
	err := easyjson.Unmarshal(data, s)
	if err != nil {
		return err
	}
	return err
}

func GetData() ([]byte, error) {
	url := "https://demo.apistubs.io/api/v1/users"
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("GET", url, nil) // создает новый http запрос
	if err != nil {
		fmt.Println("Ошибка при создании запроса:", err)
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка при отправке запроса:", err)
		return nil, err
	}
	defer resp.Body.Close()

	dataByte, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении запроса:", err)
		return nil, err
	}
	return dataByte, err
}

func main() {
	dataByte, err := GetData()
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	var sonMarsh jsonMarsh
	err = sonMarsh.Unmarshal(dataByte, &sonMarsh)
	if err != nil {
		fmt.Println("Ошибка операции Unmarshal:", err)
		return
	}
	var newDataB = make([]byte, len(dataByte))
	newDataB, err = sonMarsh.Marshal(sonMarsh)
	if err != nil {
		fmt.Println("Ошибка операции Marshal:", err)
		return
	}
	fmt.Println(string(newDataB))

	//--------------------------jiter
	var sonIter jsonIter
	err = sonIter.Unmarshal(dataByte, &sonIter)
	if err != nil {
		fmt.Println("Ошибка операции Unmarshal:", err)
		return
	}
	//var newDataB = make([]byte, len(dataByte))
	//newDataB, err = sonIter.Marshal(sonIter)
	//if err != nil {
	//	fmt.Println("Ошибка операции Marshal:", err)
	//	return
	//}
	//fmt.Println(string(newDataB))

	//---------------------easy
	var sonEasy EasyMarshJson
	var sss models.Whetherer
	err = sonEasy.Unmarshal(dataByte, &sss)
	if err != nil {
		fmt.Println("Ошибка операции Unmarshal:", err)
		return
	}
	//var newDataB = make([]byte, len(dataByte))
	//newDataB, err = sonEasy.Marshal(sss)
	//if err != nil {
	//	fmt.Println("Ошибка операции Marshal:", err)
	//	return
	//}
	//fmt.Println(string(newDataB))
}
